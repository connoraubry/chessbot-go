package game

import "chessbot-go/engine"

var pawn_table = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, -20, -20, 10, 10, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, 5, 10, 25, 25, 10, 5, 5,
	10, 10, 20, 30, 30, 20, 10, 10,
	50, 50, 50, 50, 50, 50, 50, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var knight_table = [64]int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-30, 5, 15, 20, 20, 15, 0, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

func EvaluateBoard(board engine.Board) int {
	score := 0
	score += EvaluateBoardOneColor(board, board.WhitePieces)
	score -= EvaluateBoardOneColor(board, board.BlackPieces)
	return score
}

func EvaluateBoardOneColor(board engine.Board, color_mask engine.Bitboard) int {
	score := 0
	score += 100 * CountBits(board.Pawns&color_mask)
	score += 300 * CountBits(board.Bishops&color_mask)
	score += 300 * CountBits(board.Knights&color_mask)
	score += 500 * CountBits(board.Rooks&color_mask)
	score += 800 * CountBits(board.Queens&color_mask)

	return score
}

func CountBits(bb engine.Bitboard) int {
	res := 0

	for bb > 0 {
		bb.PopLSB()
		res += 1
	}
	return res
}
