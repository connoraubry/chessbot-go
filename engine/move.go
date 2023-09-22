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
	Start int `json:"start"`
	End   int `json:"end"`

	PieceName PieceName `json:"piece_name"`
	Player    Player    `json:"player"`

	Capture bool `json:"capture"`

	Promotion PieceName `json:"promotion"`

	Castle CastleOpt `json:"castle"`

	En_passant_revealed int  `json:"en_passant_revealed"`
	En_passant          bool `json:"en_passant"`
	// 	en_passant_piece_spot int

	// 	check bool
}

func GetAlgebraicString(m Move) string {
	return fmt.Sprintf("%v%v", indexToString(m.Start), indexToString(m.End))
}

func indexToString(index int) string {
	rank := index >> 3
	file := index & 7

	return fmt.Sprintf("%v%v", string(fileToLetter[file]), string(rankToLetter[rank]))
}

func (b *Board) SpotUnderAttack(spot_bitboard Bitboard, player Player) bool {
	return b.GetAttackingPieces(spot_bitboard, player) > 0
}

// TODO: Add en passant possible attacking piece.
func (b *Board) GetAttackingPieces(spot_bitboard Bitboard, player Player) Bitboard {

	straight := GetStraightMovesBitboard(b, spot_bitboard)
	diag := GetDiagonalMovesBitboard(b, spot_bitboard)

	knight_attack := NewKnightMoves[spot_bitboard.Index()]
	king_attack := NewKingMoves[spot_bitboard.Index()]

	var pawnmoves Bitboard
	switch player {
	case WHITE:
		pawnmoves = (spot_bitboard ^ RANK_8_BB) << 8
	case BLACK:
		pawnmoves = spot_bitboard >> 8
	}

	// pawnmoves := Bitboard(1 << PawnMoveOffsets[gs.Player])
	pawn_attack_bb := ((pawnmoves & ^FILE_A_BB) >> 1) | ((pawnmoves & ^FILE_H_BB) << 1)

	opponent_bb := b.PlayerPieces(Enemy[player])

	rook_queen := (b.Rooks | b.Queens) & opponent_bb
	bishop_queen := (b.Bishops | b.Queens) & opponent_bb

	res := (straight & rook_queen)
	res |= (diag & bishop_queen)
	res |= (knight_attack & b.Knights & opponent_bb)
	res |= (pawn_attack_bb & opponent_bb & b.Pawns)
	res |= (king_attack & opponent_bb & b.Kings)
	return res
}

func (b *Board) AttackdByPawns(spot_bitboard Bitboard, player Player) bool {
	var pawnmoves Bitboard
	switch player {
	case WHITE:
		pawnmoves = (spot_bitboard ^ RANK_8_BB) << 8
	case BLACK:
		pawnmoves = spot_bitboard >> 8
	}
	pawn_attack_bb := ((pawnmoves & ^FILE_A_BB) >> 1) | ((pawnmoves & ^FILE_H_BB) << 1)
	opponent_bb := b.PlayerPieces(Enemy[player])
	return pawn_attack_bb&opponent_bb&b.Pawns > 0
}

// works
func (gs *Gamestate) GetMovesFromMoveBitboard(move_bb Bitboard, lsb Bitboard, piece PieceName) []Move {
	var moves []Move
	for move_bb > 0 {
		lsb_move := move_bb.PopLSB()

		is_capture := (lsb_move & gs.Board.PlayerPieces(Enemy[gs.Player])) > 0

		m := Move{
			Start:               lsb.Index(),
			End:                 lsb_move.Index(),
			PieceName:           piece,
			Player:              gs.Player,
			Capture:             is_capture,
			En_passant_revealed: -1,
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

func (gs *Gamestate) GetAllQueenMoves() []Move {
	var moves []Move

	queen_bb := gs.Board.Queens & gs.Board.PlayerPieces(gs.Player)

	iterate_bb := queen_bb

	for iterate_bb > 0 {
		lsb := iterate_bb.PopLSB()
		attack_Bb := GetStraightMovesBitboard(gs.Board, lsb) | GetDiagonalMovesBitboard(gs.Board, lsb)
		attack_Bb = attack_Bb &^ gs.Board.PlayerPieces(gs.Player)
		moves = append(moves, gs.GetMovesFromMoveBitboard(attack_Bb, lsb, QUEEN)...)
	}

	return moves
}

func (gs *Gamestate) GetAllKnightMoves() []Move {
	var moves []Move
	knight_bb := gs.Board.Knights & gs.Board.PlayerPieces(gs.Player)

	for knight_bb > 0 {
		current_knight := knight_bb.PopLSB()
		attack_spots := GetAllKnightMovesBitboard(gs.Board, current_knight, gs.Player)
		moves = append(moves, gs.GetMovesFromMoveBitboard(attack_spots, current_knight, KNIGHT)...)

	}
	return moves
}
func GetAllKnightMovesBitboard(board *Board, target Bitboard, player Player) Bitboard {
	idx := target.Index()
	return NewKnightMoves[idx] & ^board.PlayerPieces(player)
}

func (gs *Gamestate) GetAllKingMoves() []Move {
	var moves []Move

	king_bb := gs.Board.Kings & gs.Board.PlayerPieces(gs.Player)
	king_idx := king_bb.Index()
	attack_spots := NewKingMoves[king_idx] & ^gs.Board.PlayerPieces(gs.Player)
	moves = gs.GetMovesFromMoveBitboard(attack_spots, king_bb, KING)

	moves = append(moves, gs.GetAllCastleMoves()...)
	return moves
}

func (gs *Gamestate) GetAllCastleMoves() []Move {
	var moves []Move
	var m Move
	//WOO
	emptyBoard := gs.Board.EmptySpots()

	if gs.Player == WHITE {

		if gs.castle.whiteKing {
			if emptyBoard&WOO_EMPTY_BOARD == WOO_EMPTY_BOARD {
				m = Move{
					Start:               4,
					End:                 6,
					PieceName:           KING,
					Player:              gs.Player,
					Castle:              WHITEOO,
					En_passant_revealed: -1,
				}
				moves = append(moves, m)
			}
		}
		if gs.castle.whiteQueen {
			if emptyBoard&WOOO_EMPTY_BOARD == WOOO_EMPTY_BOARD {
				m = Move{
					Start:               4,
					End:                 2,
					PieceName:           KING,
					Player:              gs.Player,
					Castle:              WHITEOOO,
					En_passant_revealed: -1,
				}
				moves = append(moves, m)

			}
		}
	}

	if gs.Player == BLACK {
		if gs.castle.blackKing {
			if emptyBoard&BOO_EMPTY_BOARD == BOO_EMPTY_BOARD {
				m = Move{
					Start:               60,
					End:                 62,
					PieceName:           KING,
					Player:              gs.Player,
					Castle:              BLACKOO,
					En_passant_revealed: -1,
				}
				moves = append(moves, m)
			}
		}
		if gs.castle.blackQueen {
			if emptyBoard&BOOO_EMPTY_BOARD == BOOO_EMPTY_BOARD {
				m = Move{
					Start:               60,
					End:                 58,
					PieceName:           KING,
					Player:              gs.Player,
					Castle:              BLACKOOO,
					En_passant_revealed: -1,
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
	pawn_bb := gs.Board.Pawns & gs.Board.PlayerPieces(gs.Player)
	pawn_bb &= StartingPawnRank[gs.Player]

	pawn_bb.ShiftPawns(gs.Player)
	pawn_bb &= gs.Board.EmptySpots()

	pawn_bb.ShiftPawns(gs.Player)
	pawn_bb &= gs.Board.EmptySpots()

	for pawn_bb > 0 {
		lsb := pawn_bb.PopLSB()
		idx := lsb.Index()
		m := Move{
			Start:               idx - (PawnMoveOffsets[gs.Player] << 1),
			End:                 idx,
			PieceName:           PAWN,
			Player:              gs.Player,
			En_passant_revealed: idx - PawnMoveOffsets[gs.Player],
		}
		moves = append(moves, m)
	}

	return moves
}

func (gs *Gamestate) GetPawnPromotions(start, end int, capture bool) []Move {
	var moves []Move

	for _, promotion := range Promotions {

		m := Move{
			Start:               start,
			End:                 end,
			Promotion:           promotion,
			PieceName:           PAWN,
			Player:              gs.Player,
			Capture:             capture,
			En_passant_revealed: -1,
		}
		moves = append(moves, m)
	}
	return moves
}

func (gs *Gamestate) GetAllPawnOneMoves() []Move {

	var moves []Move

	pawn_bb := gs.Board.Pawns & gs.Board.PlayerPieces(gs.Player)
	pawn_bb.ShiftPawns(gs.Player)
	one_move_bb := pawn_bb & gs.Board.EmptySpots()

	for one_move_bb > 0 {
		lsb := one_move_bb.PopLSB()
		idx := lsb.Index()
		start := idx - PawnMoveOffsets[gs.Player]
		end := idx
		if lsb.Rank() == BackRank[gs.Player] {
			moves = append(moves, gs.GetPawnPromotions(start, end, false)...)
		} else {
			m := Move{
				Start:               start,
				End:                 end,
				PieceName:           PAWN,
				Player:              gs.Player,
				En_passant_revealed: -1,
			}
			moves = append(moves, m)
		}
	}

	return moves
}

func (gs *Gamestate) GetAllPawnAttackMoves() []Move {
	var moves []Move

	pawn_bb := gs.Board.Pawns & gs.Board.PlayerPieces(gs.Player)

	valid_attack_bb := gs.Board.PlayerPieces(Enemy[gs.Player])
	valid_ep_bb := gs.EnPassantBitboard()

	var offset int
	switch gs.Player {
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

			if (attack_bb & valid_attack_bb) > 0 {
				if attack_bb.Rank() == BackRank[gs.Player] {
					moves = append(moves, gs.GetPawnPromotions(idx, attack_spot_one, true)...)
				} else {
					m := Move{
						Start:               idx,
						End:                 attack_spot_one,
						PieceName:           PAWN,
						Player:              gs.Player,
						Capture:             true,
						En_passant_revealed: -1,
					}
					moves = append(moves, m)
				}
			}
			if (attack_bb & valid_ep_bb) > 0 {
				m := Move{
					Start:               idx,
					End:                 attack_spot_one,
					PieceName:           PAWN,
					Player:              gs.Player,
					Capture:             true,
					En_passant_revealed: -1,
					En_passant:          true,
				}
				moves = append(moves, m)
			}
		}

		if (lsb & FILE_H_BB) == 0 {
			attack_spot_two := idx + offset + 1
			attack_bb := Bitboard(1 << attack_spot_two)

			if (attack_bb & valid_attack_bb) > 0 {
				if attack_bb.Rank() == BackRank[gs.Player] {
					moves = append(moves, gs.GetPawnPromotions(idx, attack_spot_two, true)...)
				} else {
					m := Move{
						Start:               idx,
						End:                 attack_spot_two,
						PieceName:           PAWN,
						Player:              gs.Player,
						Capture:             true,
						En_passant_revealed: -1,
					}
					moves = append(moves, m)
				}

			}
			if (attack_bb & valid_ep_bb) > 0 {
				m := Move{
					Start:               idx,
					End:                 attack_spot_two,
					PieceName:           PAWN,
					Player:              gs.Player,
					Capture:             true,
					En_passant_revealed: -1,
					En_passant:          true,
				}
				moves = append(moves, m)
			}

		}

	}

	return moves
}

func GetLineAttacksHorizontal(occupied, slider, mask Bitboard) Bitboard {
	return mask & ((occupied - (slider << 1)) ^ (occupied.Reverse() - (slider.Reverse() << 1)).Reverse())
}

func GetLineAttacks(occupied, slider, mask Bitboard) Bitboard {
	return mask & ((occupied - (slider << 1)) ^ (occupied.VReverse() - (slider.VReverse() << 1)).VReverse())
}

func (gs *Gamestate) GetAllRookMoves() []Move {
	var moves []Move

	rook_bb := gs.Board.Rooks & gs.Board.PlayerPieces(gs.Player)

	iterate_bb := rook_bb

	for iterate_bb > 0 {
		moves = append(moves, gs.GetRookMoves(iterate_bb.PopLSB())...)
	}

	return moves
}

func (gs *Gamestate) GetRookMoves(slider Bitboard) []Move {
	bb := GetStraightMovesBitboard(gs.Board, slider) & ^gs.Board.PlayerPieces(gs.Player)
	return gs.GetMovesFromMoveBitboard(bb, slider, ROOK)
}

func GetStraightMovesBitboard(b *Board, slider Bitboard) Bitboard {

	occupied := b.AllPieces()
	rank := slider.Rank()
	file := slider.File()

	rank_mask := RANK_MASKS[rank]
	file_mask := FILE_MASKS[file]

	bb := GetLineAttacksHorizontal(occupied&rank_mask, slider, rank_mask) | GetLineAttacks(occupied&file_mask, slider, file_mask)

	return bb
}

func (gs *Gamestate) GetAllBishopMoves() []Move {
	var moves []Move

	bishop_bb := gs.Board.Bishops & gs.Board.PlayerPieces(gs.Player)

	iterate_bb := bishop_bb

	for iterate_bb > 0 {
		moves = append(moves, gs.GetBishopMoves(iterate_bb.PopLSB())...)
	}

	return moves
}

func (gs *Gamestate) GetBishopMoves(slider Bitboard) []Move {
	bb := GetDiagonalMovesBitboard(gs.Board, slider) & ^gs.Board.PlayerPieces(gs.Player)
	return gs.GetMovesFromMoveBitboard(bb, slider, BISHOP)
}

func GetDiagonalMovesBitboard(b *Board, slider Bitboard) Bitboard {

	occupied := b.AllPieces()
	rank := slider.Rank()
	file := slider.File()

	diag_mask := GetBishopDiagonal(rank, file)
	anti_diag_mask := GetBishopAntiDiagonal(rank, file)

	bb := GetLineAttacks(occupied&diag_mask, slider, diag_mask) | GetLineAttacks(occupied&anti_diag_mask, slider, anti_diag_mask)
	return bb
}
