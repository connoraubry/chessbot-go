package engine

import (
	"testing"
)

func TestBoardInit(t *testing.T) {
	b := Board{}
	if b.WhiteKing() != 0 {
		t.Fatalf(`b.whiteKing == %v. Expected 0`, b.WhiteKing())

	}
	if b.BlackKing() != 0 {
		t.Fatalf(`b.blackKing == %v. Expected 0`, b.BlackKing())

	}
	if b.AllPieces() != 0 {
		t.Fatalf(`b.AllPieces() == %v. Expected %v`, b.AllPieces(), 0)
	}

}

func TestNewBoard(t *testing.T) {
	b := NewBoard(Starting_board_fen_string)

	whiteKing := b.PlayerPieces[WHITE] & b.Kings
	whiteKingIndex := whiteKing.Index()
	if whiteKingIndex != 4 {
		t.Fatalf(`White king index == %v. Expected %v`, whiteKingIndex, 4)
	}

	blackKing := b.PlayerPieces[BLACK] & b.Kings
	blackKingIndex := blackKing.Index()
	if blackKingIndex != 60 {
		t.Fatalf(`Black king index == %v. Expected %v`, blackKingIndex, 60)
	}

}

func TestNewBoardWhiteKing(t *testing.T) {

	b := NewBoard(Starting_board_fen_string)
	if b.WhiteKing().Index() != 4 {
		t.Fatalf(`b.WhiteKing().Index() == %v. Expected 4`, b.WhiteKing().Index())

	}
}
func TestNewBoardBlackKing(t *testing.T) {
	b := NewBoard(Starting_board_fen_string)
	if b.BlackKing().Index() != 60 {
		t.Fatalf(`b.BlackKing().Index() == %v. Expected 60`, b.BlackKing().Index())
	}
}

func TestNewBoardRooks(t *testing.T) {
	b := NewBoard(Starting_board_fen_string)

	//white rooks
	expected := Bitboard(0b10000001)
	whiteRooks := b.PlayerPieces[WHITE] & b.Rooks
	if whiteRooks != expected {
		t.Fatalf(`White Rooks == %v. Expected %v`, whiteRooks, expected)
	}

	black_expected := expected << (56)
	blackRooks := b.PlayerPieces[BLACK] & b.Rooks
	if blackRooks != black_expected {
		t.Fatalf(`Black Rooks == %v. Expected %v`, blackRooks, black_expected)
	}
}

func TestNewBoardBishops(t *testing.T) {
	b := NewBoard(Starting_board_fen_string)

	whiteBishops := b.PlayerPieces[WHITE] & b.Bishops
	blackBishops := b.PlayerPieces[BLACK] & b.Bishops

	expected := Bitboard(0b00100100)
	if whiteBishops != expected {
		t.Fatalf(`whiteBishops == %v. Expected %v`, whiteBishops, expected)
	}

	black_expected := expected << (56)
	if blackBishops != black_expected {
		t.Fatalf(`blackBishops == %v. Expected %v`, blackBishops, black_expected)
	}
}

func TestNewBoardKnights(t *testing.T) {
	b := NewBoard(Starting_board_fen_string)

	whiteKnights := b.PlayerPieces[WHITE] & b.Knights
	blackKnights := b.PlayerPieces[BLACK] & b.Knights

	expected := Bitboard(0b01000010)
	if whiteKnights != expected {
		t.Fatalf(`whiteKnights == %v. Expected %v`, whiteKnights, expected)
	}

	black_expected := expected << (56)
	if blackKnights != black_expected {
		t.Fatalf(`blackKnights == %v. Expected %v`, blackKnights, black_expected)
	}
}

func TestMiddleFEN(t *testing.T) {

	b := NewBoard("r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R1B1K2R")

	expected_rooks := Bitboard(1) | 1<<7 | 1<<56 | 1<<63
	expected_bishops := Bitboard(4) | 1<<26 | 1<<58 | 1<<61
	expected_kings := Bitboard(1<<4) | 1<<59
	expected_knights := Bitboard(1<<21) | 1<<45

	all_pieces_except_pawns := expected_bishops | expected_kings | expected_rooks | expected_knights
	expected_pawns := b.AllPieces() - all_pieces_except_pawns

	whiteKingindex := b.WhiteKing().Index()
	if whiteKingindex != 4 {
		t.Fatalf(`b.WhiteKingIndex == %v. Expected 4`, whiteKingindex)
	}
	blackKingindex := (b.PlayerPieces[BLACK] & b.Kings).Index()

	if blackKingindex != 59 {
		t.Fatalf(`b.BlackKingIndex == %v. Expected 60`, blackKingindex)
	}
	if b.Bishops != expected_bishops {
		t.Fatalf(`b.Bishops == %v. Expected %v`, b.Bishops, expected_bishops)
	}
	if b.Rooks != expected_rooks {
		t.Fatalf(`b.Rooks == %v. Expected %v`, b.Rooks, expected_rooks)
	}
	if b.Kings != expected_kings {
		t.Fatalf(`b.Kings == %v. Expected %v`, b.Kings, expected_kings)
	}
	if b.Knights != expected_knights {
		t.Fatalf(`b.Knights == %v. Expected %v`, b.Knights, expected_knights)
	}
	if b.Pawns != expected_pawns {
		t.Fatalf(`b.Pawns == %v. Expected %v`, b.Pawns, expected_pawns)
	}
}
