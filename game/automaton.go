package game

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

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

	Cache             Cache
	Book              Book
	movesAnalyzed     int
	totalMovesAnalzed int

	MaxRecursionDepth int
}

const RECURSION int = 3
const HALFMOVE_CACHE_THRESHOLD int = 2

func NewAutomaton(e *engine.Engine, color engine.Player, recursion int) *Automaton {
	a := new(Automaton)
	a.Engine = e
	a.Color = color
	a.Halfmove = 0

	a.InputMoveChan = make(chan engine.Move)
	a.OutputMovechan = make(chan engine.Move)

	a.TakeMoveChan = make(chan int)

	a.QuitChan = make(chan int)
	a.Cache = *NewCache(HALFMOVE_CACHE_THRESHOLD, true)
	a.Book = *NewBook("book.json")
	if recursion == 0 {
		a.MaxRecursionDepth = RECURSION
	} else {
		a.MaxRecursionDepth = recursion
	}

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

func (a *Automaton) FlushEvalMoves() {
	flush_count := a.Cache.Flush(a.Halfmove)

	fmt.Printf("Flushed %v elements at halfmove %v. Cache count: %v\n", flush_count, a.Halfmove, a.Cache.Len())
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
			m := a.GetNextMove()
			fmt.Printf("taking move %v", m)
			a.Engine.TakeMove(m)
			a.OutputMovechan <- m
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
	fmt.Printf("Getting next move. Max analysis: %v\n", a.MaxRecursionDepth)
	a.movesAnalyzed = 0

	MAX := false
	if a.Color == engine.WHITE {
		MAX = true
	}

	if a.Halfmove < 5 {
		moveStr, err := a.Book.GetNextMove(engine.ExportToFENNoMoves(a.Engine.CurrentGamestate()))
		if err == nil {
			validMoves := a.Engine.GetValidMoves()
			strToMove := a.Engine.GetStringToMoveMap(validMoves)
			fmt.Println("Using book move")
			return strToMove[moveStr]
		}
	}

	_, nextMove := a.GetNextLevelRecursive(a.MaxRecursionDepth, math.MinInt, math.MaxInt, MAX)
	a.totalMovesAnalzed += a.movesAnalyzed

	fmt.Printf("Moves analyzed: %v. Total moves analyzed: %v\n", a.movesAnalyzed, a.totalMovesAnalzed)
	fmt.Printf("Taking move: %v\n", a.Engine.GetMoveString(nextMove, a.Engine.GetValidMoves()))
	return nextMove
}

// level is depth level. When 0, find best of the available moves
// if no moves available, return 1000000
func (a *Automaton) GetNextLevelRecursive(level, alpha, beta int, MAX bool) (int, engine.Move) {

	if level == RECURSION {
		fmt.Println("Starting recursion")
	}
	//base level of recursion, eval and report back
	if level == 0 {
		a.movesAnalyzed += 1
		return a.GetBoardScore(level), engine.Move{}
	}

	//no moves left, eval and report back
	moves := a.Engine.GetValidMoves()
	if len(moves) == 0 {
		a.movesAnalyzed += 1
		return a.GetBoardScore(level), engine.Move{}
	}

	fen := engine.ExportToFENNoMoves(a.Engine.CurrentGamestate())
	pos, ok := a.Cache.Lookup(fen)

	if ok && pos.DepthAnalyzed > level {
		return pos.Score, pos.BestMove
	}

	//randomize moves
	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	//sort on capture
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].Capture
	})

	var bestScore int
	var bestMove engine.Move

	if MAX {
		bestScore = math.MinInt
		for i, m := range moves {
			if level == RECURSION {
				fmt.Printf("Evaluating move %v out of %v\n", i, len(moves))
			}
			a.Engine.TakeMove(m)
			newScore, _ := a.GetNextLevelRecursive(level-1, alpha, beta, false)
			a.Engine.UndoMove()

			if newScore > bestScore {
				bestMove = m
			}

			bestScore = Max(bestScore, newScore)
			alpha = Max(alpha, bestScore)

			if bestScore > beta && beta != -1000000 {
				break
			}
		}
	} else {
		bestScore = math.MaxInt
		for i, m := range moves {
			if level == RECURSION {
				fmt.Printf("Evaluating move %v out of %v\n", i, len(moves))
			}
			a.Engine.TakeMove(m)
			newScore, _ := a.GetNextLevelRecursive(level-1, alpha, beta, true)
			a.Engine.UndoMove()

			if newScore < bestScore {
				bestMove = m
			}

			bestScore = Min(bestScore, newScore)
			beta = Min(beta, bestScore)

			if bestScore <= alpha && alpha != 10000000 {
				break
			}
		}
		//fmt.Println(bestScore, bestMove)
	}

	//if level == RECURSION-1 {
	//	if bestScore == 1000000 {
	//		fmt.Println(fen, a.Engine.GetMoveString(bestMove, moves))
	//	}
	//}

	if !ok {
		pos = PositionEval{
			Score:         bestScore,
			BestMove:      bestMove,
			LastHalfmove:  a.Halfmove,
			TimesAccessed: 0,
			DepthAnalyzed: level - 1,
		}
		a.Cache.Update(fen, pos)
	} else {
		if bestScore > pos.Score {
			pos.Score = bestScore
			pos.BestMove = bestMove
			pos.LastHalfmove = a.Halfmove
			pos.TimesAccessed = pos.TimesAccessed + 1
			pos.DepthAnalyzed = level - 1
			a.Cache.Update(fen, pos)
		}
	}

	return bestScore, bestMove
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
func (a *Automaton) GetBoardScore(currDepth int) int {

	gs := a.Engine.CurrentGamestate()
	fen := engine.ExportToFENNoMoves(gs)

	eval, ok := a.Cache.Lookup(fen)
	if ok && eval.DepthAnalyzed > currDepth {
		eval.TimesAccessed += 1
		eval.LastHalfmove = a.Halfmove
		a.Cache.Update(fen, eval)

		return eval.Score
	}

	newScore := 0

	board := gs.Board
	if a.Color == engine.BLACK {
		newScore = -EvaluateBoard(*board)
	} else {
		newScore = EvaluateBoard(*board)
	}

	if a.Engine.PlayerInCheckmate() {
		//if color is player in checkmate, bad bad bad
		if a.Color == gs.Player {
			newScore = -1000000
		} else {
			newScore = 1000000
		}
	}

	posEval := PositionEval{
		Score:         newScore,
		TimesAccessed: 1,
		LastHalfmove:  a.Halfmove,
		DepthAnalyzed: currDepth,
	}
	a.Cache.Update(fen, posEval)

	//if newScore > 10000 {
	//	a.Engine.CurrentGamestate().PrintBoard()
	//	fmt.Println(currDepth, newScore)
	//}
	return newScore
}
