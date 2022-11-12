package engine

import (
	"testing"
)

//for benchmarks
var result []Move

func TestGetAllKnightMoves(t *testing.T) {

	gs := NewGamestateFEN(starting_fen)

	moves := gs.GetAllKnightMoves()
	expected_length_moves := 4
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestGetPawnOneMoves(t *testing.T) {
	gs := NewGamestateFEN(starting_fen)

	moves := gs.GetAllPawnMoves()
	expected_length_moves := 16
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestGetRookMoves(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppppppp/8/8/8/P7/PPPPPPPP/R3KBNR w KQkq - 0 1")
	moves := gs.GetAllRookMoves()
	expected_length_moves := 3
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestGetVerticalMoves(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppppppp/8/P7/8/8/R2PPPPP/4KBNR w KQkq - 0 1")
	gs.Board.Rooks.PopLSB()
	lsb := gs.Board.Rooks.LSB()
	moves := gs.GetAllVerticalMovesBitboard(lsb)

	moves_length := 0
	for moves > 0 {
		moves.PopLSB()
		moves_length += 1
	}

	expected_length_moves := 3
	if moves_length != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, moves_length, expected_length_moves)
	}
}

func TestGetAllURDiagonalMovesBB(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_urd_moves := Bitboard(1152925911260069888)

	whiteBishosp := gs.Board.Bishops & gs.Board.PlayerPieces[WHITE]

	moves_bb := gs.GetAllURDiagonalMovesBitboard(whiteBishosp)
	if moves_bb != expected_urd_moves {
		t.Fatalf(`Moves != expected. %v != %v`, moves_bb, expected_urd_moves)
	}
}

func TestGetAllDRDiagonalMovesBB(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_urd_moves := Bitboard(288248105776709632)

	whiteBishosp := gs.Board.Bishops & gs.Board.PlayerPieces[WHITE]

	moves_bb := gs.GetAllDRDiagonalMovesBitboard(whiteBishosp)
	if moves_bb != expected_urd_moves {
		t.Fatalf(`Moves != expected. %v != %v`, moves_bb, expected_urd_moves)
	}
}

func TestGetBishopMoves(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_length := 7
	moves_bb := gs.GetAllBishopMoves()
	if len(moves_bb) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves_bb), expected_length)
	}
}
func TestGetQueenMoves(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_length := 14
	moves_bb := gs.GetAllQueenMoves()
	if len(moves_bb) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves_bb), expected_length)
	}
}

func BenchmarkGetPawnMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN(starting_fen)

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllPawnMoves()
	}
	result = mvs
}
func BenchmarkGetAllKnightMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN(starting_fen)

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllKnightMoves()
	}
	result = mvs
}

func BenchmarkGetRookMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("rnbqkbn1/1pppppp1/8/7r/8/P7/1PPPPPP1/R3KBNR b KQkq - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllRookMoves()
	}
	result = mvs
}
func BenchmarkGetBishopMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllBishopMoves()
	}
	result = mvs
}

func BenchmarkGetQueenMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllQueenMoves()
	}
	result = mvs
}

func BenchmarkGetAllMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("r3qb1k/1b4p1/p2pr2p/3n4/Pnp1N1N1/6RP/1B3PP1/1B1QR1K1 w - - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllMoves()
	}
	result = mvs
}
