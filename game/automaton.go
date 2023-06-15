package game

import (
	"chessbot-go/engine"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Automaton struct {
	Engine *engine.Engine

	Color engine.Player

	InputMoveChan  chan engine.Move
	OutputMovechan chan engine.Move

	//flag channels
	TakeMoveChan chan int
	QuitChan     chan int

	EvaluatedMoves map[string]PositionEval

	movesAnalyzed     int
	movesAnalyzedLock sync.Mutex
}

type PositionEval struct {
	score          int
	times_accessed int
}

const RECURSION int = 3

func NewAutomaton(e *engine.Engine, color engine.Player) *Automaton {
	a := new(Automaton)
	a.Engine = e
	a.Color = color

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
		fmt.Printf("%v %v\n", fen, values)
		if values.times_accessed > max {
			max = values.times_accessed
			max_fen = fen
		}
		if values.times_accessed > 1 {
			more_than_one += 1
		}
		total += 1
	}
	fmt.Printf("Positions with more than one visit: %v/%v\n", more_than_one, total)
	fmt.Printf("Max fen: %v. %v", max_fen, max)
}

/*
Main loop for automaton. Stays in channel loop.

a.TakeMoveChan Gets the next move and sends to output move channel
a.InputMoveChan gets a move from the gamestate and updates the current gamestate
*/
func (a *Automaton) Run() {
	var m engine.Move
	for {
		select {
		case <-a.TakeMoveChan:
			a.OutputMovechan <- a.GetNextMove()
		case m = <-a.InputMoveChan:
			a.Engine.TakeMove(m)
		case <-a.QuitChan:
			fmt.Println("quit")
			return
		}
	}
}
func (a *Automaton) GetNextMove() engine.Move {
	a.resetMovesAnalyzed()
	nextMove := a.GetNextLevel()
	// fmt.Printf("Moves analyzed: %v\n", a.movesAnalyzed)
	return nextMove
}

func (a *Automaton) GetNextMoveRandom() engine.Move {
	moves := a.Engine.GetValidMoves()
	randomIndex := rand.Intn(len(moves))

	move := moves[randomIndex]

	a.Engine.TakeMove(move)
	time.Sleep(1 * time.Second)
	a.Engine.Print()
	return move
}

func (a *Automaton) GetNextEvaluation() engine.Move {

	moves := a.Engine.GetValidMoves()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	var bestMove engine.Move
	var bestScore int
	bestScoreSet := false

	for _, m := range moves {
		a.Engine.TakeMove(m)
		newScore := a.GetBoardScore()
		if a.isScoreBetter(newScore, bestScore, bestScoreSet) {
			bestScore = newScore
			bestMove = m
			bestScoreSet = true
		}
		a.Engine.UndoMove()
	}
	a.Engine.Print()

	a.Engine.TakeMove(bestMove)
	time.Sleep(1 * time.Second)
	return bestMove
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

		score := a.GetNextLevelRecursive(RECURSION, !MAX)
		// fmt.Printf("%v %v\n", a.Engine.GetMoveString(m, moves), score)
		if score > bestScore {
			bestMove = m
			bestScore = score
		}

		a.Engine.UndoMove()
	}
	a.Engine.Print()

	a.Engine.TakeMove(bestMove)
	return bestMove
}

// level is depth level. When 1, find best of the available moves
// if no moves available, return 1000000
func (a *Automaton) GetNextLevelRecursive(level int, MAX bool) int {
	if level == 0 {
		a.incrementMovesAnalyzed()
		return a.GetBoardScore()
	}
	moves := a.Engine.GetValidMoves()
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	bestScore := initScore(MAX)

	for _, m := range moves {
		a.Engine.TakeMove(m)
		newScore := a.GetNextLevelRecursive(level-1, !MAX)
		if MAX {
			if newScore > bestScore {
				bestScore = newScore
			}
		} else {
			//minimizing
			if newScore < bestScore {
				bestScore = newScore
			}
		}
		a.Engine.UndoMove()
	}

	return bestScore
}

func initScore(MAX bool) int {
	if MAX {
		return math.MinInt
	}
	return math.MaxInt

}
func (a *Automaton) isScoreBetter(new, old int, scoreSet bool) bool {
	if !scoreSet {
		return true
	}
	return new > old
}

// always want current player's score to be positive
func (a *Automaton) GetBoardScore() int {

	gs := a.Engine.CurrentGamestate()

	fen := engine.ExportToFENNoMoves(gs)
	eval, ok := a.EvaluatedMoves[fen]
	if ok {
		eval.times_accessed += 1
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
	a.EvaluatedMoves[fen] = PositionEval{score: newScore, times_accessed: 1}
	return newScore
}

func (a *Automaton) incrementMovesAnalyzed() {
	a.movesAnalyzedLock.Lock()
	a.movesAnalyzed += 1
	a.movesAnalyzedLock.Unlock()
}

func (a *Automaton) resetMovesAnalyzed() {
	a.movesAnalyzedLock.Lock()
	a.movesAnalyzed += 0
	a.movesAnalyzedLock.Unlock()
}
