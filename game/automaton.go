package game

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/connoraubry/chessbot-go/engine"
)

type Automaton struct {
	Engine *engine.Engine

	Halfmove int
	Color    engine.Player

	InputMoveChan  chan engine.Move
	OutputMovechan chan engine.Move

	//flag channels
	TakeMoveChan chan int
	QuitChan     chan int

	EvaluatedMoves map[string]PositionEval

	movesAnalyzed     int
	totalMovesAnalzed int
}

const RECURSION int = 4
const HALFMOVE_CACHE_THRESHOLD int = 2

func NewAutomaton(e *engine.Engine, color engine.Player) *Automaton {
	a := new(Automaton)
	a.Engine = e
	a.Color = color
	a.Halfmove = 0

	a.InputMoveChan = make(chan engine.Move)
	a.OutputMovechan = make(chan engine.Move)

	a.TakeMoveChan = make(chan int)

	a.QuitChan = make(chan int)
	a.EvaluatedMoves = make(map[string]PositionEval)
	return a
}
func (a *Automaton) Quit() {
	a.QuitChan <- 1
}

func (a *Automaton) Update(move engine.Move) {
	a.InputMoveChan <- move
}

func (a *Automaton) GetMove() engine.Move {
	a.TakeMoveChan <- 1
	return <-a.OutputMovechan
}

func (a *Automaton) Dump() {
	more_than_one := 0
	total := 0
	max := 0
	var max_fen string
	for fen, values := range a.EvaluatedMoves {
		if values.timesAccessed > max {
			max = values.timesAccessed
			max_fen = fen
		}
		if values.timesAccessed > 1 {
			more_than_one += 1
		}
		total += 1
	}
	fmt.Printf("Positions with more than one visit: %v/%v\n", more_than_one, total)
	fmt.Printf("Max fen: %v. %v", max_fen, max)
}

func (a *Automaton) FlushEvalMoves() {
	threshold := a.Halfmove - HALFMOVE_CACHE_THRESHOLD
	flush_count := 0

	for fen, values := range a.EvaluatedMoves {
		if values.lastHalfmove < threshold {
			delete(a.EvaluatedMoves, fen)
			flush_count += 1
		}
	}

	fmt.Printf("Flushed %v elements at halfmove %v. Cache count: %v\n", flush_count, a.Halfmove, len(a.EvaluatedMoves))

}

/*
Main loop for automaton. Stays in channel loop.

a.TakeMoveChan Gets the next move and sends to output move channel
a.InputMoveChan gets a move from the gamestate and updates the current gamestate
*/
func (a *Automaton) Run() {
	var m engine.Move
	currHalfMove := a.Halfmove
	for {
		select {
		case <-a.TakeMoveChan:
			a.OutputMovechan <- a.GetNextMove()
			a.Halfmove += 1
		case m = <-a.InputMoveChan:

			a.Engine.TakeMove(m)
			a.Halfmove += 1
		case <-a.QuitChan:
			fmt.Println("quit")
			return
		}
		if currHalfMove != a.Halfmove {
			a.FlushEvalMoves()
			currHalfMove = a.Halfmove
		}
	}
}
func (a *Automaton) GetNextMove() engine.Move {
	a.movesAnalyzed = 0
	nextMove := a.GetNextLevel()
	a.totalMovesAnalzed += a.movesAnalyzed

	fmt.Printf("Moves analyzed: %v. Total moves analyzed: %v\n", a.movesAnalyzed, a.totalMovesAnalzed)
	fmt.Printf("Taking move: %v\n", a.Engine.GetMoveString(nextMove, a.Engine.GetValidMoves()))
	return nextMove
}

func (a *Automaton) GetNextLevel() engine.Move {
	moves := a.Engine.GetValidMoves()
	rand.Seed(time.Now().UnixNano())

	//shuffle the moves so we don't pick the first one every time.
	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	MAX := true
	bestScore := initScore(MAX)
	var bestMove engine.Move
	for _, m := range moves {
		a.Engine.TakeMove(m)

		score := a.GetNextLevelRecursive(RECURSION, math.MinInt, math.MaxInt, !MAX)
		// fmt.Printf("%v %v\n", a.Engine.GetMoveString(m, moves), score)
		if score > bestScore {
			bestMove = m
			bestScore = score
		}

		a.Engine.UndoMove()
	}
	//a.Engine.Print(2, a.LastMove)

	a.Engine.TakeMove(bestMove)
	return bestMove
}

// level is depth level. When 1, find best of the available moves
// if no moves available, return 1000000
func (a *Automaton) GetNextLevelRecursive(level, alpha, beta int, MAX bool) int {
	if level == 0 {
		a.movesAnalyzed += 1
		return a.GetBoardScore()
	}
	moves := a.Engine.GetValidMoves()
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	var bestScore int

	if MAX {
		bestScore = math.MinInt
		for _, m := range moves {
			a.Engine.TakeMove(m)
			newScore := a.GetNextLevelRecursive(level-1, alpha, beta, !MAX)
			a.Engine.UndoMove()

			bestScore = Max(bestScore, newScore)
			alpha = Max(alpha, bestScore)
			if bestScore >= beta {
				break
			}
		}
	} else {
		bestScore := math.MaxInt
		for _, m := range moves {
			a.Engine.TakeMove(m)
			newScore := a.GetNextLevelRecursive(level-1, alpha, beta, !MAX)
			a.Engine.UndoMove()

			bestScore = Min(bestScore, newScore)
			beta = Min(beta, bestScore)
			if bestScore <= alpha {
				break
			}
		}
	}

	return bestScore
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func initScore(MAX bool) int {
	if MAX {
		return math.MinInt
	}
	return math.MaxInt
}

// always want current player's score to be positive
func (a *Automaton) GetBoardScore() int {

	gs := a.Engine.CurrentGamestate()

	fen := engine.ExportToFENNoMoves(gs)
	eval, ok := a.EvaluatedMoves[fen]
	if ok {
		eval.timesAccessed += 1
		eval.lastHalfmove = a.Halfmove
		a.EvaluatedMoves[fen] = eval
		return eval.score
	}

	newScore := 0

	board := gs.Board
	if a.Color == engine.BLACK {
		newScore = -EvaluateBoard(*board)
	} else {
		newScore = EvaluateBoard(*board)
	}
	a.EvaluatedMoves[fen] = PositionEval{score: newScore, timesAccessed: 1}
	return newScore
}
