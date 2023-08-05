package engine

import "fmt"

type Board struct {
	Kings   Bitboard
	Queens  Bitboard
	Rooks   Bitboard
	Knights Bitboard
	Bishops Bitboard
	Pawns   Bitboard

	WhitePieces Bitboard
	BlackPieces Bitboard
}

func (b *Board) PlayerPieces(player Player) Bitboard {
	var pieces Bitboard
	switch player {
	case WHITE:
		pieces = b.WhitePieces
	case BLACK:
		pieces = b.BlackPieces
	}
	return pieces
}

func NewBoard(FEN_board string) *Board {
	b := new(Board)
	b.loadFENPositionsBitBoard(FEN_board)
	return b
}

func (b *Board) EmptySpots() Bitboard {
	return ^b.AllPieces()
}

func (b *Board) AllPieces() Bitboard {
	return b.PlayerPieces(WHITE) | b.PlayerPieces(BLACK)
}

func (b *Board) CopyBoard() *Board {
	newB := Board{
		Kings:       b.Kings,
		Queens:      b.Queens,
		Rooks:       b.Rooks,
		Knights:     b.Knights,
		Bishops:     b.Bishops,
		Pawns:       b.Pawns,
		WhitePieces: b.WhitePieces,
		BlackPieces: b.BlackPieces,
	}

	return &newB
}

func (b *Board) ClearSpot(piece Bitboard) {
	b.Pawns &= ^piece
	b.Knights &= ^piece
	b.Bishops &= ^piece
	b.Rooks &= ^piece
	b.Queens &= ^piece
	b.Kings &= ^piece
	b.WhitePieces &= ^piece
	b.BlackPieces &= ^piece
}

func (b *Board) RemovePiece(piece Bitboard, pieceName PieceName, player Player) {
	switch pieceName {
	case PAWN:
		b.Pawns &= ^piece
	case KNIGHT:
		b.Knights &= ^piece
	case BISHOP:
		b.Knights &= ^piece
	case ROOK:
		b.Rooks &= ^piece
	case QUEEN:
		b.Queens &= ^piece
	case KING:
		b.Kings &= ^piece
	}
	switch player {
	case WHITE:
		b.WhitePieces &= ^piece
	case BLACK:
		b.BlackPieces &= ^piece
	}
}

func (b *Board) AddPiece(piece Bitboard, pieceName PieceName, player Player) {
	switch pieceName {
	case PAWN:
		b.Pawns |= piece
	case KNIGHT:
		b.Knights |= piece
	case BISHOP:
		b.Bishops |= piece
	case ROOK:
		b.Rooks |= piece
	case QUEEN:
		b.Queens |= piece
	case KING:
		b.Kings |= piece
	}
	switch player {
	case WHITE:
		b.WhitePieces |= piece
	case BLACK:
		b.BlackPieces |= piece
	}
}

func (b *Board) PiecesFromName(n PieceName) Bitboard {
	switch n {
	case PAWN:
		return b.Pawns
	case KNIGHT:
		return b.Knights
	case BISHOP:
		return b.Bishops
	case ROOK:
		return b.Rooks
	case QUEEN:
		return b.Queens
	case KING:
		return b.Kings
	default:
		return Bitboard(0)
	}
}

func (b *Board) WhiteKing() Bitboard {
	return b.PlayerPieces(WHITE) & b.Kings
}
func (b *Board) BlackKing() Bitboard {
	return b.PlayerPieces(BLACK) & b.Kings
}

func (b *Board) PieceAndPlayerFromIndex(index int) (PieceName, Player) {
	var piece PieceName
	var player Player

	mask := Bitboard(1 << index)

	if mask&b.WhitePieces > 0 {
		player = WHITE
	} else if mask&b.BlackPieces > 0 {
		player = BLACK
	}

	if mask&b.Pawns > 0 {
		piece = PAWN
	} else if mask&b.Knights > 0 {
		piece = KNIGHT
	} else if mask&b.Bishops > 0 {
		piece = BISHOP
	} else if mask&b.Rooks > 0 {
		piece = ROOK
	} else if mask&b.Queens > 0 {
		piece = QUEEN
	} else if mask&b.Kings > 0 {
		piece = KING
	}

	return piece, player

}

func (b *Board) ExportToFEN() string {
	rank := 7
	file := 0
	// split := 7

	var bytelist []byte

	count := 0
	for file < 8 && rank >= 0 {

		piece, player := b.PieceAndPlayerFromIndex(file + (rank * 8))

		if piece == EMPTY {
			count += 1
		} else {
			if count > 0 {
				bytelist = append(bytelist, byte(count)+'0')
				count = 0
			}
			bytelist = append(bytelist, byte(pieceAndPlayertoLetter(piece, player)))
		}

		file += 1

		if file == 8 {
			file = 0
			rank -= 1

			if count > 0 {
				bytelist = append(bytelist, byte(count)+'0')
				count = 0
			}
			if rank >= 0 {
				bytelist = append(bytelist, '/')

			}

		}
	}

	return string(bytelist)

}

func (b *Board) loadFENPositionsBitBoard(FEN_board string) {

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

		if number >= 0 && number < 9 {
			file += int(number)
			continue
		}
		player := letterToPlayer(char)
		name := letterToPieceName(char)

		mask := Bitboard(1 << index)
		b.AddPiece(mask, name, player)

		file += 1

	}
}

func (b *Board) PrintBoard() {
	stringBoard := b.GetBoardVisualString()
	for _, line := range stringBoard {
		fmt.Println(line)
	}
}
func (b *Board) GetBoardVisualString() []string {
	visString := make([]string, 10)
	rank := 7
	file := 0
	bottom := "  a b c d e f g h"
	visString[0] = bottom
	visString[9] = bottom

	row := []rune{'8'}
	idx := 1
	for (file >= 0) && (rank >= 0) {
		index := file + (rank * 8)

		piecename, player := b.GetSpotPiece(index)

		uni := getUnicode(piecename, player)

		row = append(row, ' ', uni)

		file += 1

		if file > 7 {
			rank -= 1
			file = 0
			visString[idx] = string(row)
			idx += 1
			row = []rune{rune(rank + 49)}
		}
	}
	return visString
}

func (b *Board) GetSpotPiece(idx int) (PieceName, Player) {
	var name PieceName
	var player Player
	if (b.Kings>>idx)&1 > 0 {
		name = KING
	} else if (b.Queens>>idx)&1 > 0 {
		name = QUEEN
	} else if (b.Rooks>>idx)&1 > 0 {
		name = ROOK
	} else if (b.Bishops>>idx)&1 > 0 {
		name = BISHOP
	} else if (b.Knights>>idx)&1 > 0 {
		name = KNIGHT
	} else if (b.Pawns>>idx)&1 > 0 {
		name = PAWN
	}

	if (b.PlayerPieces(WHITE)>>idx)&1 > 0 {
		player = WHITE
	} else if (b.PlayerPieces(BLACK)>>idx)&1 > 0 {
		player = BLACK
	}

	return name, player
}
