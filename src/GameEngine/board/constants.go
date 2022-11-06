package board

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

type Piece struct {
	name   PieceName
	player Player
}

// const starting_fen string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const Starting_board_fen_string string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
