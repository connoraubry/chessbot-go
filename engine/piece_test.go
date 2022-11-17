package engine

import "testing"

func TestLetterToPieceName(t *testing.T) {
	letterToExpected := map[rune]PieceName{
		'p': PAWN,
		'P': PAWN,
		'r': ROOK,
		'R': ROOK,
		'q': QUEEN,
		'Q': QUEEN,
		'K': KING,
		'k': KING,
		'b': BISHOP,
		'B': BISHOP,
		'n': KNIGHT,
		'N': KNIGHT,
	}

	for input, output := range letterToExpected {
		if letterToPieceName(input) != output {
			t.Fatalf(`letterToPieceName(%v) == %v. Expected %v`, input, letterToPieceName(input), output)
		}

	}
}

func TestLetterToPlayer(t *testing.T) {
	letterToExpected := map[rune]Player{
		'p': BLACK,
		'P': WHITE,
		'r': BLACK,
		'R': WHITE,
		'q': BLACK,
		'Q': WHITE,
		'K': WHITE,
		'k': BLACK,
		'b': BLACK,
		'B': WHITE,
		'n': BLACK,
		'N': WHITE,
	}

	for input, output := range letterToExpected {
		if letterToPlayer(input) != output {
			t.Fatalf(`letterToPieceName(%v) == %v. Expected %v`, input, letterToPieceName(input), output)
		}

	}
}

func TestPieceAndPlayerToLetter(t *testing.T) {

	output := pieceAndPlayertoLetter(PAWN, BLACK)

	if output != 'p' {
		t.Fatalf(`%v != %v`, 'p', output)
	}
}
