package main

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
)

var (
	filepath = flag.String("bitboards", "./saved_bitboards", "Directory of saved bitboards")
	fen      = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

func main() {
	flag.Parse()

	e := *engine.NewEngine(
		engine.OptBitboardDirectory(*filepath),
		engine.OptFenFile(*fen),
	)

	engine.PrintBitboard(e.GameState.Board.Kings)

	engine.PrintBitboard(engine.KNIGHT_ATTACKS[1])
	fmt.Println(e.SlidingBitrow)
}
