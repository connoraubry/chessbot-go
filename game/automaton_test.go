package game

import (
	"testing"

	"github.com/connoraubry/chessbot-go/engine"
)

func TestCountBits(t *testing.T) {
	bb := engine.Bitboard(0b10010100101)
	expected := 5

	if CountBits(bb) != expected {
		t.Fatalf(`CountBits(%v) == %v. Expected %v`, bb, CountBits(bb), expected)
	}
}

func TestEvaluationMateInTwoOne(t *testing.T) {
	fen_string := "r2qkb1r/pp2nppp/3p4/2pNN1B1/2BnP3/3P4/PPP2PPP/R2bK2R w KQkq - 1 0"
	expectedMove := "Nf6+"

	actualMove := getExpectedString(fen_string, 3)

	if actualMove != expectedMove {
		t.Fatalf(`Predicted move %v != best %v`, actualMove, expectedMove)
	}
}
func TestEvaluationMateInTwoTwo(t *testing.T) {
	fen_string := "1rb4r/pkPp3p/1b1P3n/1Q6/N3Pp2/8/P1P3PP/7K w - - 1 0"
	expectedMove := "Qd5+"

	actualMove := getExpectedString(fen_string, 3)

	if actualMove != expectedMove {
		t.Fatalf(`Predicted move %v != best %v`, actualMove, expectedMove)
	}
}

func TestEvaluationMateInTwoThree(t *testing.T) {
	fen_string := "4kb1r/p2n1ppp/4q3/4p1B1/4P3/1Q6/PPP2PPP/2KR4 w k - 1 0"
	expectedMove := "Qb8+"

	actualMove := getExpectedString(fen_string, 3)

	if actualMove != expectedMove {
		t.Fatalf(`Predicted move %v != best %v`, actualMove, expectedMove)
	}
}

func TestEvaluationMateInTwoFour(t *testing.T) {
	fen_string := "4rr2/1p5R/3p1p2/p2Bp3/P2bPkP1/1P5R/1P2K3/8 w - - 1 0"
	expectedMove := "Rg7"

	actualMove := getExpectedString(fen_string, 3)

	if actualMove != expectedMove {
		t.Fatalf(`Predicted move %v != best %v`, actualMove, expectedMove)
	}
}

func getExpectedString(fen_string string, recursion int) string {
	e := engine.NewEngine(engine.OptFenString(fen_string))
	autoE := engine.NewEngine(engine.OptFenString(fen_string))
	a := NewAutomaton(autoE, engine.WHITE, 3)

	moveString := e.GetMoveString(a.GetNextMove(), e.GetValidMoves())

	return moveString
}
