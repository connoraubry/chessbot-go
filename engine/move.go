package engine

import (
	"fmt"
)

type Move struct {
	start int
	end   int

	pieceName PieceName
	player    Player

	capture bool

	promotion PieceName

	// 	en_passant_revealed   int
	// 	en_passant_piece_spot int

	// 	check bool
}

func indexToString(index int) string {
	rank := index >> 3
	file := index & 7

	var fileToLetter = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	var rankToLetter = []rune{'1', '2', '3', '4', '5', '6', '7', '8'}

	return fmt.Sprintf("%v%v", string(fileToLetter[file]), string(rankToLetter[rank]))
}

//TODO: add support for multiple pieces attacking the same spot
func (m *Move) String() string {

	letter := getLetter(m.pieceName)
	var capture rune
	if m.capture {
		capture = 'x'
	}
	s := fmt.Sprintf("%v%v%v", string(letter), string(capture), indexToString(m.end))
	return s
}

//works
func (gs *Gamestate) GetMovesFromMoveBitboard(move_bb Bitboard, lsb Bitboard, piece PieceName) []Move {
	var moves []Move
	for move_bb > 0 {
		lsb_move := move_bb.PopLSB()

		is_capture := (lsb_move & gs.Board.PlayerPieces[Enemy[gs.player]]) > 0

		m := Move{
			start:     lsb.Index(),
			end:       lsb_move.Index(),
			pieceName: piece,
			player:    gs.player,
			capture:   is_capture,
		}
		moves = append(moves, m)
	}
	return moves
}

func (gs *Gamestate) GetAllMoves() []Move {
	var moves []Move
	moves = gs.GetAllPawnMoves()
	moves = append(moves, gs.GetAllKnightMoves()...)
	moves = append(moves, gs.GetAllBishopMoves()...)
	moves = append(moves, gs.GetAllRookMoves()...)
	moves = append(moves, gs.GetAllQueenMoves()...)
	return moves
}

func (gs *Gamestate) GetAllPawnMoves() []Move {
	var moves []Move

	moves = append(moves, gs.GetAllPawnOneMoves()...)
	moves = append(moves, gs.GetAllPawnDoubleMoves()...)

	return moves
}

func (gs *Gamestate) GetAllBishopMoves() []Move {
	var moves []Move

	bishop_bb := gs.Board.Bishops & gs.Board.PlayerPieces[gs.player]

	iterate_bb := bishop_bb

	for iterate_bb > 0 {
		lsb := iterate_bb.PopLSB()
		ur_bitboard := gs.GetAllURDiagonalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(ur_bitboard, lsb, BISHOP)...)
		dr_bitboard := gs.GetAllDRDiagonalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(dr_bitboard, lsb, BISHOP)...)
	}

	return moves
}

func (gs *Gamestate) GetAllRookMoves() []Move {
	var moves []Move

	rook_bb := gs.Board.Rooks & gs.Board.PlayerPieces[gs.player]

	iterate_bb := rook_bb

	for iterate_bb > 0 {
		lsb := iterate_bb.PopLSB()
		horizontal_bitboard := gs.GetAllHorizontalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(horizontal_bitboard, lsb, ROOK)...)
		vertical_bitboard := gs.GetAllVerticalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(vertical_bitboard, lsb, ROOK)...)
	}

	return moves
}

func (gs *Gamestate) GetAllQueenMoves() []Move {
	var moves []Move

	queen_bb := gs.Board.Queens & gs.Board.PlayerPieces[gs.player]

	iterate_bb := queen_bb

	for iterate_bb > 0 {
		lsb := iterate_bb.PopLSB()
		horizontal_bitboard := gs.GetAllHorizontalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(horizontal_bitboard, lsb, QUEEN)...)
		vertical_bitboard := gs.GetAllVerticalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(vertical_bitboard, lsb, QUEEN)...)
		ur_bitboard := gs.GetAllURDiagonalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(ur_bitboard, lsb, QUEEN)...)
		dr_bitboard := gs.GetAllDRDiagonalMovesBitboard(lsb)
		moves = append(moves, gs.GetMovesFromMoveBitboard(dr_bitboard, lsb, QUEEN)...)
	}

	return moves
}

func (gs *Gamestate) GetAllKnightMoves() []Move {
	var moves []Move

	knight_bb := gs.Board.Knights & gs.Board.PlayerPieces[gs.player]

	for knight_bb > 0 {
		current_knight := knight_bb.PopLSB()
		attack_spots := KNIGHT_ATTACKS[current_knight] & ^gs.Board.PlayerPieces[gs.player]

		moves = append(moves, gs.GetMovesFromMoveBitboard(attack_spots, current_knight, KNIGHT)...)

	}
	return moves
}

func (gs *Gamestate) GetAllPawnDoubleMoves() []Move {
	var moves []Move

	//only pawns on starting rank
	pawn_bb := gs.Board.Pawns & gs.Board.PlayerPieces[gs.player]
	pawn_bb &= StartingPawnRank[gs.player]

	pawn_bb.ShiftPawns(gs.player)
	pawn_bb &= gs.Board.EmptySpots()

	pawn_bb.ShiftPawns(gs.player)
	pawn_bb &= gs.Board.EmptySpots()

	for pawn_bb > 0 {
		lsb := pawn_bb.PopLSB()
		idx := lsb.Index()
		m := Move{
			start:     idx - (PawnMoveOffsets[gs.player] << 1),
			end:       idx,
			pieceName: PAWN,
			player:    gs.player,
		}
		moves = append(moves, m)
	}

	return moves
}

func (gs *Gamestate) GetAllPawnOneMoves() []Move {

	var moves []Move

	pawn_bb := gs.Board.Pawns & gs.Board.PlayerPieces[gs.player]
	pawn_bb.ShiftPawns(gs.player)
	one_move_bb := pawn_bb & gs.Board.EmptySpots()

	for one_move_bb > 0 {
		lsb := one_move_bb.PopLSB()
		idx := lsb.Index()
		if lsb.Rank() == BackRank[gs.player] {

			for _, promotion := range Promotions {

				m := Move{
					start:     idx - PawnMoveOffsets[gs.player],
					end:       idx,
					promotion: promotion,
					pieceName: PAWN,
					player:    gs.player,
				}
				moves = append(moves, m)
			}
		} else {
			m := Move{
				start:     idx - PawnMoveOffsets[gs.player],
				end:       idx,
				pieceName: PAWN,
				player:    gs.player,
			}
			moves = append(moves, m)
		}
	}

	return moves
}

//b only has one bit set -- the rook
func (gs *Gamestate) GetAllHorizontalMovesBitboard(lsb Bitboard) Bitboard {
	rank := lsb.Rank()

	bitrow := Bitrow(gs.Board.AllPieces() >> (8 * rank))
	lsb_bitrow := Bitrow(lsb >> (8 * rank))

	valid_moves := SlidingBitrow[bitrow][lsb_bitrow]
	valid_moves_bb := Bitboard(valid_moves) << (8 * rank)
	valid_moves_bb &= ^gs.Board.PlayerPieces[gs.player]
	return valid_moves_bb

}

func (gs *Gamestate) GetAllVerticalMovesBitboard(lsb Bitboard) Bitboard {

	file := lsb.File()

	a_file_board := (gs.Board.AllPieces() >> (file)) & FILE_A_BB
	a_file_lsb := (lsb >> (file)) & FILE_A_BB

	bitrow_flipped := Bitrow(AFileToRank(a_file_board))
	lsb_flipped := Bitrow(AFileToRank(a_file_lsb))

	valid_moves := Bitboard(SlidingBitrow[bitrow_flipped][lsb_flipped])

	valid_moves_a_file := RankToAFile(valid_moves)

	valid_moves_vertical := valid_moves_a_file << file
	valid_moves_vertical &= ^gs.Board.PlayerPieces[gs.player]

	return valid_moves_vertical

}

func (gs *Gamestate) GetAllURDiagonalMovesBitboard(lsb Bitboard) Bitboard {
	idx := lsb.Index()
	board := gs.Board.AllPieces()

	urd := ConvertToURDiagonal(board, idx)
	rankBB := Bitrow(URDiagonalToRank(urd))

	lsb_row := Bitrow(lsb >> (lsb.Rank() * RANK_SHIFT_1))
	moves_bb := Bitboard(SlidingBitrow[rankBB][lsb_row])

	movesUrd := RankToURDiagonal(moves_bb)
	moves := ReverseConvertToURDiagonal(movesUrd, idx)

	return moves & ^gs.Board.PlayerPieces[gs.player]
}

func (gs *Gamestate) GetAllDRDiagonalMovesBitboard(lsb Bitboard) Bitboard {
	idx := lsb.Index()
	board := gs.Board.AllPieces()

	drd := ConvertToDRDiagonal(board, idx)
	rankBB := Bitrow(DRDiagonalToRank(drd))
	lsb_row := Bitrow(lsb >> (lsb.Rank() * RANK_SHIFT_1))
	moves_bb := Bitboard(SlidingBitrow[rankBB][lsb_row])
	movesDrd := RankToDRDiagonal(moves_bb)
	moves := ReverseConvertToDRDiagonal(movesDrd, idx)

	return moves & ^gs.Board.PlayerPieces[gs.player]
}
