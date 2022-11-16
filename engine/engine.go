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

func (e *Engine) TakeMove(m Move) {

	var newGamestate Gamestate
	var newBoard Board

	currentGs := e.CurrentGamestate()

	newCastle := currentGs.castle.Copy()

	// if m.Castle != NO_CASTLE {
	// 	newBoard = *e.TakeCastle(m)
	// } else {
	// 	newBoard
	// }

	newBoard = *e.CurrentGamestate().Board.CopyBoard()

	newBoard.ClearSpot(Bitboard(1 << m.start))
	newBoard.ClearSpot(Bitboard(1 << m.end))

	newBoard.AddPiece(Bitboard(1<<m.end), m.pieceName, m.player)

	// currentGs := e.CurrentGamestate()

	new_halfmove := currentGs.halfmove + 1
	fullmove_increment := 0

	if m.capture || m.pieceName == PAWN || m.Castle != NO_CASTLE {
		new_halfmove = 0
	}

	if m.player == BLACK {
		fullmove_increment = 1
	}

	newGamestate = Gamestate{
		Board:      &newBoard,
		player:     Enemy[currentGs.player],
		castle:     newCastle,
		en_passant: m.en_passant_revealed,
		halfmove:   new_halfmove,
		fullmove:   currentGs.fullmove + fullmove_increment,
	}

	e.GameStates = append(e.GameStates, newGamestate)
}

//TODO: Check if king in check for spots 1 and 2
// func (e *Engine) TakeCastle(m Move) *Board {
// 	newBoard := e.CurrentGamestate().Board.CopyBoard()
// 	switch m.Castle {
// 	case WHITEOO:
// 		newBoard.Kings &= ^Bitboard(16)
// 		newBoard.Rooks &= ^Bitboard(128)
// 		newBoard.PlayerPieces(WHITE) &= Bitboard(18446744073709551375)

// 		newBoard.Rooks |= Bitboard(32)
// 		newBoard.Kings |= Bitboard(64)
// 		newBoard.PlayerPieces(WHITE) |= Bitboard(96)
// 	case WHITEOOO:
// 		newBoard.Kings &= ^Bitboard(16)
// 		newBoard.Rooks &= ^Bitboard(1)
// 		newBoard.PlayerPieces(WHITE) &= Bitboard(18446744073709551584)

// 		newBoard.Kings |= Bitboard(4)
// 		newBoard.Rooks |= Bitboard(8)
// 		newBoard.PlayerPieces(WHITE) |= Bitboard(12)
// 	case BLACKOO:
// 		newBoard.Kings &= ^Bitboard(1152921504606846976)
// 		newBoard.Rooks &= ^Bitboard(9223372036854775808)
// 		newBoard.PlayerPieces[BLACK] &= Bitboard(1152921504606846975)

// 		newBoard.Kings |= Bitboard(4611686018427387904)
// 		newBoard.Rooks |= Bitboard(2305843009213693952)
// 		newBoard.PlayerPieces[BLACK] |= Bitboard(6917529027641081856)
// 	case BLACKOOO:
// 		newBoard.Kings &= ^Bitboard(1152921504606846976)
// 		newBoard.Rooks &= ^Bitboard(72057594037927936)
// 		newBoard.PlayerPieces[BLACK] &= Bitboard(16212958658533785599)

// 		newBoard.Kings |= Bitboard(288230376151711744)
// 		newBoard.Rooks |= Bitboard(576460752303423488)
// 		newBoard.PlayerPieces[BLACK] |= Bitboard(864691128455135232)
// 	}
// 	return newBoard
// }

func (e *Engine) UndoMove() {
	if len(e.GameStates) > 0 {
		e.GameStates = e.GameStates[:len(e.GameStates)-1]
	}
}

func (e *Engine) Print() {

	cgs := e.CurrentGamestate()
	cgs.PrintBoard()

	fmt.Printf("Castle: %v\nEn Passant: %v\nHalfmove: %v\nFullmove: %v\n",
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
