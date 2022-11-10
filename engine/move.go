package engine

type MoveGenerator struct {
	// gs *Gamestate
	// moves []Move
}

type Move struct {
	// 	start int
	// 	end   int

	// 	en_passant_revealed   int
	// 	en_passant_piece_spot int

	// 	check bool
}

// func (mg *MoveGenerator) GetAllMoves(gs *Gamestate) {

// }

func rank_and_file(index int) (int, int) {
	return index / 8, index % 8
}

// func (mg *MoveGenerator) GetPawnMoves(index int) []Move {
// 	var moves []Move
// 	piece := mg.gs.board.getPiece(index)

// 	rank, file := rank_and_file(index)

// 	if rank == 0 || rank == 7 {
// 		return moves
// 	}
// 	offset := 8
// 	if piece.Player == BLACK {
// 		offset = -8
// 	}

// 	one_spot := index + offset
// 	one_spot_piece := mg.gs.board.getPiece(one_spot)
// 	if one_spot_piece.Name == EMPTY {
// 		//empty
// 	}

// 	return moves
// }
