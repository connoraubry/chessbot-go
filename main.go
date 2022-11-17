package main

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
)

var (
	fen = flag.String("fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "FEN string")
)

func main() {
	flag.Parse()

	e := engine.NewEngine(
		engine.OptFenString(*fen),
	)

	// e.Print()

	moves := e.GetAllMoves()
	moves_st := make([]string, len(moves))
	for idx, m := range moves {
		moves_st[idx] = m.String()
	}
	count := 0
	for _, m := range moves {
		success := e.TakeMove(m)
		if success {
			count += 1
			e.UndoMove()
		}

	}
	// fmt.Println(moves_st)
	fmt.Println(count)
}
