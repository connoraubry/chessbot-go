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

	// e := *engine.NewEngine(
	// 	engine.OptBitboardDirectory(*filepath),
	// 	engine.OptFenFile(*fen),
	// )

	// engine.PrintBitboard(e.GameState.Board.Kings)

	// engine.PrintBitboard(engine.KNIGHT_ATTACKS[1])
	// fmt.Println(e.SlidingBitrow)

	// kmoves := e.GameState.GetAllRooklMoves()
	// for _, move:= range kmoves {
	// 	fmt.Prin tln(move.String())
	// }

	gs := engine.NewGamestateFEN("rnbqkbn1/1pppppp1/8/7r/8/P7/1PPPPPP1/R3KBNR b KQkq - 0 1")
	engine.PrintBitboard(gs.Board.Rooks)
	moves := gs.GetAllRookMoves()

	for _, move := range moves {
		fmt.Println(move.String())
	}

	ur_diag := engine.MAIN_DIAGONAL & ^engine.FILE_A_BB & ^engine.FILE_C_BB
	expected := engine.RANK_1_BB & ^engine.FILE_A_BB & ^engine.FILE_C_BB

	output := engine.URDiagonalToRank(ur_diag)

	engine.PrintBitboard(expected)
	engine.PrintBitboard(output)

	engine.PrintBitboard(engine.RankToURDiagonal(expected))

	// test_bb := engine.FILE_A_BB & (engine.RANK_1_BB | engine.RANK_2_BB | engine.RANK_3_BB | engine.RANK_5_BB | engine.RANK_7_BB)
	// engine.PrintBitboard(test_bb)
	// fmt.Println("rank")
	// rank_bb := test_bb.FileToRank()
	// engine.PrintBitboard(rank_bb)
	// fmt.Println("Back to file")
	// file_bb := rank_bb.RankToFile()
	// engine.PrintBitboard(file_bb)

	gs = engine.NewGamestateFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	moves = gs.GetAllKnightMoves()
	for _, move := range moves {
		fmt.Println(move.String())
	}

	gs = engine.NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")
	board := gs.Board.AllPieces()
	// all_expected := 1441174017036779520
	expected_urd_moves := engine.Bitboard(288248105776709632)
	bishop := gs.Board.Bishops & gs.Board.PlayerPieces[engine.WHITE]

	engine.PrintBitboard(bishop)
	engine.PrintBitboard(board)
	fmt.Println("STart")
	engine.PrintBitboard(expected_urd_moves)
	idx := bishop.Index()

	toURD := engine.ConvertToDRDiagonal(board, idx)
	engine.PrintBitboard(toURD)
	toRank := engine.DRDiagonalToRank(toURD)
	engine.PrintBitboard(toRank)

	lsb_row := engine.Bitrow(bishop >> (bishop.Rank() * 8))

	moves_bb := engine.Bitboard(engine.SlidingBitrow[toRank][lsb_row])
	engine.PrintBitboard(moves_bb)

	backToURD := engine.RankToDRDiagonal(moves_bb)
	engine.PrintBitboard(backToURD)
	backToStart := engine.ReverseConvertToDRDiagonal(backToURD, idx)
	engine.PrintBitboard(backToStart)

}
