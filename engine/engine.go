package engine

import (
	"fmt"
)

type Engine struct {
	//'constants'
	opts       Options
	GameStates []Gamestate
}

func NewEngine(opts ...interface{}) *Engine {

	e := &Engine{}
	var err error
	e.opts, err = ParseOptions(opts...)
	if err != nil {
		panic(err)
	}

	if e.opts.FenString != "" {
		e.GameStates = []Gamestate{*NewGamestateFEN(e.opts.FenString)}
	} else {
		e.GameStates = []Gamestate{*NewGamestateFile(e.opts.FenFilePath)}

	}

	return e
}

func (e *Engine) CurrentGamestate() *Gamestate {

	return &e.GameStates[len(e.GameStates)-1]
}

func (e *Engine) GetAllMoves() []Move {
	return e.CurrentGamestate().GetAllMoves()
}

func (e *Engine) GetValidMoves() []Move {
	var valid_moves []Move

	all_moves := e.GetAllMoves()

	for _, m := range all_moves {
		valid := e.TakeMove(m)
		if valid {
			valid_moves = append(valid_moves, m)
			e.UndoMove()
		}
	}
	return valid_moves
}

func (e *Engine) ExportToFEN() string {
	return ExportToFEN(e.CurrentGamestate())
}

func (e *Engine) TakeMove(m Move) bool {

	var newGamestate Gamestate
	var newBoard *Board
	var success bool

	currentGs := e.CurrentGamestate()

	newCastle := currentGs.castle.Copy()

	if m.Castle != NO_CASTLE {
		newBoard, success = e.TakeCastle(m)
		if !success {
			return false
		}
		switch m.player {
		case WHITE:
			newCastle.whiteKing = false
			newCastle.whiteQueen = false
		case BLACK:
			newCastle.blackKing = false
			newCastle.blackQueen = false
		}
	} else {
		newBoard = e.CurrentGamestate().Board.CopyBoard()

		newBoard.ClearSpot(Bitboard(1 << m.start))
		newBoard.ClearSpot(Bitboard(1 << m.end))

		if m.pieceName == PAWN {
			if m.promotion != EMPTY {
				newBoard.AddPiece(Bitboard(1<<m.end), m.promotion, m.player)
			} else if m.en_passant {
				//en passant
				newBoard.AddPiece(Bitboard(1<<m.end), m.pieceName, m.player)
				pawn_offset := PawnMoveOffsets[m.player]
				newBoard.RemovePiece(Bitboard(1<<(m.end-pawn_offset)), PAWN, Enemy[m.player])
			} else {
				newBoard.AddPiece(Bitboard(1<<m.end), m.pieceName, m.player)
			}
		} else {
			newBoard.AddPiece(Bitboard(1<<m.end), m.pieceName, m.player)

		}

	}

	//edit castling
	newCastle.whiteKing = newCastle.whiteKing && WhiteKingCastleValid(newBoard)
	newCastle.whiteQueen = newCastle.whiteQueen && WhiteQueenCastleValid(newBoard)
	newCastle.blackKing = newCastle.blackKing && BlackKingCastleValid(newBoard)
	newCastle.blackQueen = newCastle.blackKing && BlackQueenCastleValid(newBoard)

	currKing := newBoard.PlayerPieces(m.player) & newBoard.Kings
	if newBoard.SpotUnderAttack(currKing, m.player) {
		return false
	}
	new_halfmove := currentGs.halfmove + 1
	fullmove_increment := 0

	if m.capture || m.pieceName == PAWN || m.Castle != NO_CASTLE || currentGs.castle != newCastle {
		new_halfmove = 0
	}

	if m.player == BLACK {
		fullmove_increment = 1
	}

	newGamestate = Gamestate{
		Board:      newBoard,
		Player:     Enemy[currentGs.Player],
		castle:     newCastle,
		en_passant: m.en_passant_revealed,
		halfmove:   new_halfmove,
		fullmove:   currentGs.fullmove + fullmove_increment,
	}

	e.GameStates = append(e.GameStates, newGamestate)
	return true
}

func WhiteKingCastleValid(b *Board) bool {
	startKing := Bitboard(16)
	startRook := Bitboard(128)
	if b.Kings&b.WhitePieces&startKing == 0 {
		return false
	}
	if b.Rooks&b.WhitePieces&startRook == 0 {
		return false
	}
	return true
}
func WhiteQueenCastleValid(b *Board) bool {
	startKing := Bitboard(16)
	startRook := Bitboard(1)
	if b.Kings&b.WhitePieces&startKing == 0 {
		return false
	}
	if b.Rooks&b.WhitePieces&startRook == 0 {
		return false
	}
	return true
}

func BlackKingCastleValid(b *Board) bool {
	startKing := Bitboard(1152921504606846976)
	startRook := Bitboard(9223372036854775808)
	if b.Kings&b.BlackPieces&startKing == 0 {
		return false
	}
	if b.Rooks&b.BlackPieces&startRook == 0 {
		return false
	}
	return true
}
func BlackQueenCastleValid(b *Board) bool {
	startKing := Bitboard(1152921504606846976)
	startRook := Bitboard(72057594037927936)
	if b.Kings&b.BlackPieces&startKing == 0 {
		return false
	}
	if b.Rooks&b.BlackPieces&startRook == 0 {
		return false
	}
	return true
}

func (e *Engine) TakeCastle(m Move) (*Board, bool) {

	newBoard := e.CurrentGamestate().Board.CopyBoard()

	var startKing Bitboard
	var endKing Bitboard
	var startRook Bitboard
	var endRook Bitboard
	var player Player

	switch m.Castle {
	case WHITEOO:
		startKing = Bitboard(16)
		endKing = Bitboard(64)
		startRook = Bitboard(128)
		endRook = Bitboard(32)
		player = WHITE

	case WHITEOOO:
		startKing = Bitboard(16)
		endKing = Bitboard(4)
		startRook = Bitboard(1)
		endRook = Bitboard(8)
		player = WHITE

	case BLACKOO:
		startKing = Bitboard(1152921504606846976)
		endKing = Bitboard(4611686018427387904)
		startRook = Bitboard(9223372036854775808)
		endRook = Bitboard(2305843009213693952)
		player = BLACK
	case BLACKOOO:
		startKing = Bitboard(1152921504606846976)
		endKing = Bitboard(288230376151711744)
		startRook = Bitboard(72057594037927936)
		endRook = Bitboard(576460752303423488)
		player = BLACK
	}
	newBoard.RemovePiece(startKing, KING, player)
	newBoard.RemovePiece(startRook, ROOK, player)

	newBoard.AddPiece(endRook, KING, player)
	if newBoard.SpotUnderAttack(endRook, player) {
		return newBoard, false
	}
	newBoard.RemovePiece(endRook, KING, player)

	newBoard.AddPiece(endKing, KING, player)
	newBoard.AddPiece(endRook, ROOK, player)

	if newBoard.SpotUnderAttack(endKing, player) {
		return newBoard, false
	}

	return newBoard, !newBoard.SpotUnderAttack(endKing, player)

}
func (e *Engine) UndoMove() {
	if len(e.GameStates) > 0 {
		e.GameStates = e.GameStates[:len(e.GameStates)-1]
	}
}

func (e *Engine) Print() {

	cgs := e.CurrentGamestate()
	cgs.PrintBoard()

	var player string

	switch cgs.Player {
	case WHITE:
		player = "WHITE"
	case BLACK:
		player = "BLACK"
	}

	fmt.Printf("Move: %v\nCastle: %v\nEn Passant: %v\nHalfmove: %v\nFullmove: %v\n",
		player,
		cgs.castle.ToString(),
		EPToString(cgs.en_passant),
		cgs.halfmove,
		cgs.fullmove)
}

func EPToString(ep int) string {
	if ep == -1 {
		return "-"
	} else {
		return indexToString(ep)
	}
}

func (e *Engine) AllMovesToStrings(moves []Move) map[Move]string {
	res := make(map[Move]string)

	//check if in check
	// suffix := ""

	attack_spot_moves := make(map[int][]Move)

	for _, m := range moves {
		attack_spot_moves[m.end] = append(attack_spot_moves[m.end], m)
	}

	for _, m := range moves {

		other_moves := attack_spot_moves[m.end]
		if len(other_moves) == 1 {
			res[m] = m.String()
		} else {
			//more than one piece attacking this spot
			piece_occurance := make(map[PieceName]int)
			for _, subm := range other_moves {
				piece_occurance[subm.pieceName] += 1
			}

			if piece_occurance[m.pieceName] == 1 {
				res[m] = m.String()
			} else {
				//prioritize file, rank, both
				file_map := make(map[int]int)
				rank_map := make(map[int]int)
				for _, subm := range other_moves {
					rank, file := IndexToRankFile(subm.start)
					file_map[file] += 1
					rank_map[rank] += 1
				}

				rank, file := IndexToRankFile(m.start)

				if file_map[file] == 1 {
					fmt.Println("file")
				} else if rank_map[rank] == 1 {
					fmt.Println("rank")
				} else {
					fmt.Println("Both")
				}

			}

		}

	}

	fmt.Println(attack_spot_moves)

	return res
}
