package game

import (
	"chessbot-go/engine"
	"testing"
)

func TestEvaluateBoard(t *testing.T) {

	board := engine.NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR")

	evaluation := EvaluateBoard(*board)
	expected := 0
	if evaluation != expected {
		t.Fatalf(`Evaluation != expected. %v != %v`, evaluation, expected)
	}

}

func TestEvaluateBoardWhite(t *testing.T) {

	board := engine.NewBoard("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR")

	evaluation := EvaluateBoardOneColor(*board, board.WhitePieces)
	expected := 3800
	if evaluation != expected {
		t.Fatalf(`Evaluation != expected. %v != %v`, evaluation, expected)
	}

}
