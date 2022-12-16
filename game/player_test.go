package game

import (
	"chessbot-go/engine"
	"testing"
)

func TestCountBits(t *testing.T) {
	bb := engine.Bitboard(0b10010100101)
	expected := 5

	if CountBits(bb) != expected {
		t.Fatalf(`CountBits(%v) == %v. Expected %v`, bb, CountBits(bb), expected)
	}

}

func TestEvaluateBoard(t *testing.T) {

	board := engine.NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR")

	evaluation := EvaluateBoard(*board, engine.WHITE)
	expected := 3800
	if evaluation != expected {
		t.Fatalf(`Evaluation != expected. %v != %v`, evaluation, expected)
	}

}
