package game

import (
	"github.com/connoraubry/chessbot-go/engine"
	"testing"
)

func TestCountBits(t *testing.T) {
	bb := engine.Bitboard(0b10010100101)
	expected := 5

	if CountBits(bb) != expected {
		t.Fatalf(`CountBits(%v) == %v. Expected %v`, bb, CountBits(bb), expected)
	}

}
