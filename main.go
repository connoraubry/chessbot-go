package main

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
)

var (
// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

func main() {
	flag.Parse()

	// e := engine.NewEngine(
	// 	engine.OptFenString("rnbqkb1r/p4ppp/2p1pn2/1p1p1Q2/4P3/2N2N2/PPPP1PPP/R1B1KB1R w KQkq b6 0 6"),
	// )

	b := engine.NewBoard("r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R1B1K2R")
	expected_b := engine.NewBoard("r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/2B1K2R")
	b.RemovePiece(engine.Bitboard(1), engine.ROOK, engine.WHITE)

	b.PrintBoard()
	expected_b.PrintBoard()

	fmt.Println(*b == *expected_b)
	fmt.Println(b.Pawns == expected_b.Pawns)
	fmt.Println(b.Rooks == expected_b.Rooks)
	fmt.Println(b.Queens == expected_b.Queens)
	fmt.Println(b.Kings == expected_b.Kings)
	fmt.Println(b.Bishops == expected_b.Bishops)
	fmt.Println(b.Knights == expected_b.Knights)
	fmt.Println(b.WhitePieces == expected_b.WhitePieces)

}
