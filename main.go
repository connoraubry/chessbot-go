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
}
