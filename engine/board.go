package engine

type Board struct {
	Kings        Bitboard
	Queens       Bitboard
	Rooks        Bitboard
	Knights      Bitboard
	Bishops      Bitboard
	Pawns        Bitboard
	PlayerPieces map[Player]Bitboard
}

func NewBoard(FEN_board string) *Board {
	b := new(Board)
	b.PlayerPieces = make(map[Player]Bitboard)
	b.loadFENPositionsBitBoard(FEN_board)
	return b
}

func (b *Board) AllPieces() Bitboard {
	return b.PlayerPieces[WHITE] | b.PlayerPieces[BLACK]
}

func (b *Board) EmptySpots() Bitboard {
	return ^b.AllPieces()
}

func (b *Board) WhitePieces() Bitboard {
	return b.PlayerPieces[WHITE]
}
func (b *Board) BlackPieces() Bitboard {
	return b.PlayerPieces[BLACK]
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
	return b.PlayerPieces[WHITE] & b.Kings
}
func (b *Board) BlackKing() Bitboard {
	return b.PlayerPieces[BLACK] & b.Kings
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
		} else {
			player := letterToPlayer(char)
			name := letterToPieceName(char)

			mask := Bitboard(1 << index)

			switch name {
			case PAWN:
				b.Pawns |= mask
			case KNIGHT:
				b.Knights |= mask
			case BISHOP:
				b.Bishops |= mask
			case ROOK:
				b.Rooks |= mask
			case QUEEN:
				b.Queens |= mask
			case KING:
				b.Kings |= mask
			}

			b.PlayerPieces[player] |= mask
			file += 1
		}
	}
}
