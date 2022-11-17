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

var Enemy = map[Player]Player{
	WHITE: BLACK,
	BLACK: WHITE,
}

func getLetter(name PieceName) rune {
	var letter rune

	switch name {
	case KNIGHT:
		letter = 'N'
	case BISHOP:
		letter = 'B'
	case ROOK:
		letter = 'R'
	case QUEEN:
		letter = 'Q'
	case KING:
		letter = 'K'
	}

	return letter
}

func getUnicode(name PieceName, player Player) rune {
	var uni rune
	if player == WHITE {
		switch name {
		case KING:
			uni = '\u2654'
		case QUEEN:
			uni = '\u2655'
		case ROOK:
			uni = '\u2656'
		case BISHOP:
			uni = '\u2657'
		case KNIGHT:
			uni = '\u2658'
		case PAWN:
			uni = '\u2659'
		}
	} else if player == BLACK {
		switch name {
		case KING:
			uni = '\u265A'
		case QUEEN:
			uni = '\u265B'
		case ROOK:
			uni = '\u265C'
		case BISHOP:
			uni = '\u265D'
		case KNIGHT:
			uni = '\u265E'
		case PAWN:
			uni = '\u265F'
		}
	} else {
		uni = '\u00B7'
	}

	return uni
}

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

func pieceAndPlayertoLetter(piece PieceName, player Player) rune {
	var uni rune
	if player == WHITE {
		switch piece {
		case KING:
			uni = 'K'
		case QUEEN:
			uni = 'Q'
		case ROOK:
			uni = 'R'
		case BISHOP:
			uni = 'B'
		case KNIGHT:
			uni = 'N'
		case PAWN:
			uni = 'P'
		}
	} else if player == BLACK {
		switch piece {
		case KING:
			uni = 'k'
		case QUEEN:
			uni = 'q'
		case ROOK:
			uni = 'r'
		case BISHOP:
			uni = 'b'
		case KNIGHT:
			uni = 'n'
		case PAWN:
			uni = 'p'
		}
	}

	return uni
}
