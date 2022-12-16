package engine

import "testing"

func TestCastleStruct(t *testing.T) {
	cs := Castle{}
	cs.whiteKing = true
	cs.blackKing = true
	cs.whiteQueen = true
	cs.blackQueen = true
}

func TestNewGamestate(t *testing.T) {
	gs := NewGamestateFEN(starting_fen)

	expected_castle := Castle{true, true, true, true}
	if gs.castle != expected_castle {
		t.Fatalf(`gs.castle == %v. Expected %v`, gs.castle, expected_castle)
	}

	expected_player := WHITE
	if gs.Player != expected_player {
		t.Fatalf(`gs.Player == %v. Expected %v`, gs.Player, expected_player)
	}

}

func TestEnPassantBoard(t *testing.T) {
	gs := NewGamestateFEN(starting_fen)
	expected_board := Bitboard(0)

	if gs.EnPassantBitboard() != expected_board {
		t.Fatalf(`gs.EnPassantBitboard() == %v. Expected %v.`, gs.EnPassantBitboard(), expected_board)
	}
}
func TestEnPassantBoard2(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")
	expected_board := Bitboard(1048576)

	if gs.EnPassantBitboard() != expected_board {
		t.Fatalf(`gs.EnPassantBitboard() == %v. Expected %v.`, gs.EnPassantBitboard(), expected_board)
	}
}

func TestEnPassantBoard3(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2")
	expected_board := Bitboard(17592186044416)

	if gs.EnPassantBitboard() != expected_board {
		t.Fatalf(`gs.EnPassantBitboard() == %v. Expected %v.`, gs.EnPassantBitboard(), expected_board)
	}
}
