package main

import (
	"fmt"

	"github.com/connoraubry/chessbot-go/engine"
	"github.com/connoraubry/chessbot-go/game"
)

func main() {
	fmt.Println("vim-go")

	fen_string := "r1b3kr/ppp1Bp1p/1b6/n2P4/2p3q1/2Q2N2/P4PPP/RN2R1K1 w - - 1 0"

	e := engine.NewEngine(engine.OptFenString(fen_string))
	autoE := engine.NewEngine(engine.OptFenString(fen_string))
	a := game.NewAutomaton(autoE, engine.WHITE, 0)

	m := a.GetNextMove()

	fmt.Println(m)

	allMoves := e.GetValidMoves()
	for _, move := range allMoves {

		moveStr := e.GetMoveString(move, allMoves)

		e.TakeMove(move)

		fen := engine.ExportToFENNoMoves(e.CurrentGamestate())
		score, _ := a.Cache.GetScore(fen)
		bestMove, _ := a.Cache.GetBestMove(fen)
		bestMoveStr := e.GetMoveString(bestMove, e.GetValidMoves())

		e.TakeMove(bestMove)

		fen = engine.ExportToFENNoMoves(e.CurrentGamestate())
		scoreSecond, _ := a.Cache.GetScore(fen)
		bestMoveSecond, _ := a.Cache.GetBestMove(fen)
		bestMoveSecondStr := e.GetMoveString(bestMoveSecond, e.GetValidMoves())

		fmt.Printf("Move: %5v %5v %5v Score: %10v %6v\n",
			moveStr, bestMoveStr, bestMoveSecondStr, score, scoreSecond)

		e.UndoMove()
		e.UndoMove()
	}

	fmt.Printf("Taking move: %v\n", e.GetMoveString(m, allMoves))

}
