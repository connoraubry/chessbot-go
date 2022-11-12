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
	if gs.player != expected_player {
		t.Fatalf(`gs.player == %v. Expected %v`, gs.player, expected_player)
	}

}
