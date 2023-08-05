package main

import (
	"github.com/connoraubry/chessbot-go/engine"
	"flag"
	"fmt"
	"time"
)

var (
	// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
	n   = flag.Int("n", 6, "Perft depth")
	opt = flag.Int("opt", 1, "Perft board")
)

func main() {
	flag.Parse()

	startingFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	//https://www.chessprogramming.org/Perft_Results
	switch *opt {
	case 1:
		startingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	case 2:
		//good through 5
		startingFen = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	case 3:
		//good through 5
		startingFen = "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1"
	case 4:
		//good through 5
		startingFen = "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1"
	case 5:
		//good through 5
		startingFen = "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"
	case 6:
		//good through 5
		startingFen = "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10"
	case 7:
		startingFen = "rnb2k1r/pp1PbBpp/2p5/q7/8/8/PPP1NnPP/RNBQK2R w KQ - 1 9"
	case 8:
		startingFen = "r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1"
	}

	e := engine.NewEngine(
		engine.OptFenString(startingFen),
	)
	fmt.Printf("Running perft board %v\n", *opt)
	for i := 1; i <= *n; i++ {
		start := time.Now()
		res := Perft(e, i)
		duration := time.Since(start)
		fmt.Printf("Depth: %v. Legal moves: %v\n", i, res)
		fmt.Println(duration)
	}

}

func Perft(e *engine.Engine, depth int) int {
	count := 0
	if depth == 0 {
		return 1
	}

	for _, move := range e.GetAllMoves() {
		res := e.TakeMove(move)
		if res {
			count += Perft(e, depth-1)
			e.UndoMove()
		}
	}

	return count
}
