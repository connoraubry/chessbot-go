package engine

import "unicode"

type PieceName int64
type Player int64

const (
	EMPTY PieceName = iota
	PAWN
	KNIGHT
	BISHOP
	ROOK
	QUEEN
	KING
)
const (
	WHITE Player = iota
	BLACK
)

func letterToPlayer(letter rune) Player {

	player := WHITE

	// black
	if letter > 96 {
		player = BLACK
	}

	return player
}

func letterToPieceName(letter rune) PieceName {
	name := EMPTY

	switch unicode.ToLower(letter) {
	case 'p':
		name = PAWN
	case 'n':
		name = KNIGHT
	case 'b':
		name = BISHOP
	case 'r':
		name = ROOK
	case 'q':
		name = QUEEN
	case 'k':
		name = KING
	}
	return name
}
