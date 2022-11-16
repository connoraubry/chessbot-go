package main

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
	"time"
)

var (
// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

func main() {
	flag.Parse()

	e := engine.NewEngine(
		engine.OptFenString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
	)

	for i := 1; i < 7; i++ {
		start := time.Now()
		res := Perft(e, i)
		duration := time.Since(start)
		fmt.Printf("Depth: %v. Legal moves: %v\n", i, res)
		fmt.Println(duration)
	}

}

func Perft(e *engine.Engine, depth int) int {
	count := 0
	if depth == 1 {
		return len(e.GetAllMoves())
	}
	for _, move := range e.GetAllMoves() {
		e.TakeMove(move)
		count += Perft(e, depth-1)
		e.UndoMove()
	}
	return count

}
