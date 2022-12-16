package main

import (
	"chessbot-go/engine"
	"fmt"
)

var (
// fen = flag.String("fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "FEN string")
)

func main() {

	e := engine.NewEngine(engine.OptFenString("4B1r1/P1b2PpQ/3R4/3p4/1P1B4/1p3k2/1N1K4/N7 w - - 0 1"))

	e.Print()

	moves := e.GetAllMoves()

	for _, m := range moves {
		if m.String() == "Ra1" {
			fmt.Println(m)
		}
		// fmt.Println(m.String())
	}

	allWhitePieces := e.CurrentGamestate().Board.WhitePieces

	for allWhitePieces > 0 {
		lsb := allWhitePieces.PopLSB()
		engine.PrintBitboard(lsb)

	}

}
