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
	NO_PLAYER Player = iota
	WHITE
	BLACK
)

func letterToPlayer(letter rune) Player {

	var player Player
	// black
	if letter > 64 && letter < 91 {
		player = WHITE
	} else if letter > 96 && letter < 122 {
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
