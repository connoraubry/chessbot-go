package board

import (
	"unicode"
)

type Board struct {
	board     [64]Piece
	whiteKing int
	blackKing int
}

func NewBoard(FEN_board string) *Board {
	b := new(Board)
	b.whiteKing = -1
	b.blackKing = -1
	b.loadFENPositions(FEN_board)
	return b
}

func (b *Board) loadFENPositions(FEN_board string) {

	rank := 7
	file := 0

	split := '/'

	for _, char := range FEN_board {

		index := file + (rank * 8)

		if char == split {
			file = 0
			rank -= 1
			continue
		}
		number := char - '0'

		if number >= 0 && number < 7 {
			// for i := 0; i < int(number); i ++ {
			// 	b.board[8*rank + file] = Piece{name: EMPTY}
			// }
			file += int(number)
		} else {
			b.board[index] = letterToPiece(char)
			switch char {
			case 'K':
				b.whiteKing = index
			case 'k':
				b.blackKing = index
			}
			file += 1
		}
	}
}

func letterToPiece(letter rune) Piece {

	name := EMPTY
	player := WHITE

	// black
	if letter > 96 {
		player = BLACK
	}
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

	return Piece{name: name, player: player}
}
