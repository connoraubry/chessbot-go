package main

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
)

var (
// filepath = flag.String("bitboards", "./saved_bitboards", "Directory of saved bitboards")
// fen      = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

func main() {
	flag.Parse()

	gs := engine.NewGamestateFEN("5pp1/5P2/8/8/8/8/8/8 w - - 0 1")
	moves := gs.GetAllPawnMoves()

	for _, m := range moves {
		fmt.Println(m.String())
	}

}
