package main

import (
	"chessbot-go/game"
)

var (
// fen = flag.String("fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "FEN string")
)

func main() {

	g := game.NewGame()
	g.Run()
	// input = *fen
	// input := "k7/6p1/8/7P/8/8/8/K7 w - - 0 1"

	// e := engine.NewEngine(engine.OptFenString(input))
	// e.Print()
	// aut := game.NewAutomaton(e, engine.WHITE)

	// fmt.Println(aut.GetBoardScore())
	// m := aut.GetNextMove()
	// fmt.Println(m.String())
}
