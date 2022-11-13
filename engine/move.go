package engine

import (
	"fmt"
)

type CastleOpt int

const (
	NO_CASTLE CastleOpt = iota
	WHITEOO
	BLACKOO
	WHITEOOO
	BLACKOOO
)

type Move struct {
	start int
	end   int

	pieceName PieceName
	player    Player

	capture bool

	promotion PieceName

	castle CastleOpt

	// 	en_passant_revealed   int
	// 	en_passant_piece_spot int

	// 	check bool
}

func indexToString(index int) string {
	rank := index >> 3
	file := index & 7

	return fmt.Sprintf("%v%v", string(fileToLetter[file]), string(rankToLetter[rank]))
}

//TODO: add support for multiple pieces attacking the same spot
func (m *Move) String() string {

	switch m.castle {
	case WHITEOO, BLACKOO:
		return "O-O"
	case WHITEOOO, BLACKOOO:
		return "O-O-O"
	}

	var string_Bytes []byte

	letter := getLetter(m.pieceName)

	if letter != 0 {
		string_Bytes = append(string_Bytes, byte(letter))
	}

	if m.capture {

		if m.pieceName == PAWN {
			string_Bytes = append(string_Bytes, byte(fileToLetter[m.start&7]))
		}

		string_Bytes = append(string_Bytes, 'x')
	}

	string_Bytes = append(string_Bytes, []byte(indexToString(m.end))...)

	if m.promotion != EMPTY {
		string_Bytes = append(string_Bytes, '=', byte(getLetter(m.promotion)))

	}

	return string(string_Bytes)
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
	moves = append(moves, gs.GetAllKingMoves()...)
	return moves
}

func (gs *Gamestate) GetAllPawnMoves() []Move {
	var moves []Move

	moves = append(moves, gs.GetAllPawnOneMoves()...)
	moves = append(moves, gs.GetAllPawnDoubleMoves()...)
	moves = append(moves, gs.GetAllPawnAttackMoves()...)

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
func (gs *Gamestate) GetAllKingMoves() []Move {
	var moves []Move

	king_bb := gs.Board.Kings & gs.Board.PlayerPieces[gs.player]
	attack_spots := KING_ATTACKS[king_bb] & ^gs.Board.PlayerPieces[gs.player]
	moves = gs.GetMovesFromMoveBitboard(attack_spots, king_bb, KING)

	moves = append(moves, gs.GetAllCastleMoves()...)
	return moves
}

func (gs *Gamestate) GetAllCastleMoves() []Move {
	var moves []Move
	var m Move
	//WOO
	emptyBoard := gs.Board.EmptySpots()

	if gs.player == WHITE {

		if gs.castle.whiteKing {
			if emptyBoard&WOO_EMPTY_BOARD == WOO_EMPTY_BOARD {
				m = Move{
					start:     4,
					end:       6,
					pieceName: KING,
					player:    gs.player,
					castle:    WHITEOO,
				}
				moves = append(moves, m)
			}
		}
		if gs.castle.whiteQueen {
			if emptyBoard&WOOO_EMPTY_BOARD == WOOO_EMPTY_BOARD {
				m = Move{
					start:     4,
					end:       2,
					pieceName: KING,
					player:    gs.player,
					castle:    WHITEOOO,
				}
				moves = append(moves, m)

			}
		}
	}

	if gs.castle.whiteKing {
		if gs.castle.blackKing {
			if emptyBoard&BOO_EMPTY_BOARD == BOO_EMPTY_BOARD {
				m = Move{
					start:     60,
					end:       62,
					pieceName: KING,
					player:    gs.player,
					castle:    BLACKOO,
				}
				moves = append(moves, m)
			}
		}
		if gs.castle.blackQueen {
			if emptyBoard&BOOO_EMPTY_BOARD == BOOO_EMPTY_BOARD {
				m = Move{
					start:     60,
					end:       58,
					pieceName: KING,
					player:    gs.player,
					castle:    BLACKOOO,
				}
				moves = append(moves, m)

			}
		}
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

func (gs *Gamestate) GetPawnPromotions(start, end int, capture bool) []Move {
	var moves []Move

	for _, promotion := range Promotions {

		m := Move{
			start:     start,
			end:       end,
			promotion: promotion,
			pieceName: PAWN,
			player:    gs.player,
			capture:   capture,
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
		start := idx - PawnMoveOffsets[gs.player]
		end := idx
		if lsb.Rank() == BackRank[gs.player] {
			moves = append(moves, gs.GetPawnPromotions(start, end, false)...)
		} else {
			m := Move{
				start:     start,
				end:       end,
				pieceName: PAWN,
				player:    gs.player,
			}
			moves = append(moves, m)
		}
	}

	return moves
}

func (gs *Gamestate) GetAllPawnAttackMoves() []Move {
	var moves []Move

	pawn_bb := gs.Board.Pawns & gs.Board.PlayerPieces[gs.player]

	var offset int
	switch gs.player {
	case WHITE:
		offset = NORTH
	case BLACK:
		offset = SOUTH
	}

	for pawn_bb > 0 {
		lsb := pawn_bb.PopLSB()
		idx := lsb.Index()

		if (lsb & FILE_A_BB) == 0 {
			attack_spot_one := idx + offset - 1
			attack_bb := Bitboard(1 << attack_spot_one)

			if (attack_bb & gs.Board.PlayerPieces[Enemy[gs.player]]) > 0 {
				if attack_bb.Rank() == BackRank[gs.player] {
					moves = append(moves, gs.GetPawnPromotions(idx, attack_spot_one, true)...)
				} else {
					m := Move{
						start:     idx,
						end:       attack_spot_one,
						pieceName: PAWN,
						player:    gs.player,
						capture:   true,
					}
					moves = append(moves, m)
				}

			}
		}

		if (lsb & FILE_H_BB) == 0 {
			attack_spot_two := idx + offset + 1
			attack_bb := Bitboard(1 << attack_spot_two)

			if (attack_bb & gs.Board.PlayerPieces[Enemy[gs.player]]) > 0 {
				if attack_bb.Rank() == BackRank[gs.player] {
					moves = append(moves, gs.GetPawnPromotions(idx, attack_spot_two, true)...)
				} else {
					m := Move{
						start:     idx,
						end:       attack_spot_two,
						pieceName: PAWN,
						player:    gs.player,
						capture:   true,
					}
					moves = append(moves, m)
				}

			}

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
