package engine

import (
	"fmt"
)

type Engine struct {
	//'constants'
	opts       Options
	GameStates []Gamestate
	MoveList   []Move
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

func (e *Engine) GameInStalemate() bool {
	moves := e.GetValidMoves()
	return len(moves) == 0 && !e.PlayerInCheck()

}

func (e *Engine) PlayerInCheck() bool {

	gs := e.CurrentGamestate()

	board := gs.Board
	player := gs.Player
	currKing := board.PlayerPieces(player) & board.Kings
	return board.SpotUnderAttack(currKing, player)
}

func (e *Engine) PlayerInCheckmate() bool {
	return len(e.GetValidMoves()) == 0 && e.PlayerInCheck()
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
	newCastle.blackQueen = newCastle.blackQueen && BlackQueenCastleValid(newBoard)

	currKing := newBoard.PlayerPieces(m.player) & newBoard.Kings
	if newBoard.SpotUnderAttack(currKing, m.player) {
		return false
	}
	new_halfmove := currentGs.halfmove + 1
	fullmove_increment := 0

	if m.Capture || m.pieceName == PAWN || m.Castle != NO_CASTLE || currentGs.castle != newCastle {
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
	e.MoveList = append(e.MoveList, m)
	return true
}

func WhiteKingCastleValid(b *Board) bool {
	if b.Kings&b.WhitePieces&WHITE_KING_START == 0 {
		return false
	}
	if b.Rooks&b.WhitePieces&WHITE_KING_ROOK_START == 0 {
		return false
	}
	return true
}
func WhiteQueenCastleValid(b *Board) bool {
	if b.Kings&b.WhitePieces&WHITE_KING_START == 0 {
		return false
	}
	if b.Rooks&b.WhitePieces&WHITE_QUEEN_ROOK_START == 0 {
		return false
	}
	return true
}

func BlackKingCastleValid(b *Board) bool {
	if b.Kings&b.BlackPieces&BLACK_KING_START == 0 {
		return false
	}
	if b.Rooks&b.BlackPieces&BLACK_KING_ROOK_START == 0 {
		return false
	}
	return true
}
func BlackQueenCastleValid(b *Board) bool {
	if b.Kings&b.BlackPieces&BLACK_KING_START == 0 {
		return false
	}
	if b.Rooks&b.BlackPieces&BLACK_QUEEN_ROOK_START == 0 {
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
		startKing = WHITE_KING_START
		endKing = WHITE_KING_OO_CASTLE
		startRook = WHITE_KING_ROOK_START
		endRook = WHITE_KING_ROOK_CASTLE
		player = WHITE

	case WHITEOOO:
		startKing = WHITE_KING_START
		endKing = WHITE_KING_OOO_CASTLE
		startRook = WHITE_QUEEN_ROOK_START
		endRook = WHITE_QUEEN_ROOK_CASTLE
		player = WHITE

	case BLACKOO:
		startKing = BLACK_KING_START
		endKing = BLACK_KING_OO_CASTLE
		startRook = BLACK_KING_ROOK_START
		endRook = BLACK_KING_ROOK_CASTLE
		player = BLACK
	case BLACKOOO:
		startKing = BLACK_KING_START
		endKing = BLACK_KING_OOO_CASTLE
		startRook = BLACK_QUEEN_ROOK_START
		endRook = BLACK_QUEEN_ROOK_CASTLE
		player = BLACK
	}

	if newBoard.SpotUnderAttack(startKing, player) {
		return newBoard, false
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

	return newBoard, !newBoard.SpotUnderAttack(endKing, player)

}
func (e *Engine) UndoMove() {
	if len(e.GameStates) > 0 {
		e.GameStates = e.GameStates[:len(e.GameStates)-1]
	}
}

func (e *Engine) Print(spaces int, lastMove string) {

	cgs := e.CurrentGamestate()
	boardString := cgs.Board.GetBoardVisualString()

	var player string

	switch cgs.Player {
	case WHITE:
		player = "WHITE"
	case BLACK:
		player = "BLACK"
	}

	rightSpaces := "   "

	if lastMove != "" {
		boardString[4] = fmt.Sprintf("%v%vLast Move: %v", boardString[4], rightSpaces, lastMove)
	}
	boardString[5] = fmt.Sprintf("%v%vCurr Player: %v", boardString[5], rightSpaces, player)
	boardString[6] = fmt.Sprintf("%v%vCastle: %v", boardString[6], rightSpaces, cgs.castle.ToString())
	boardString[7] = fmt.Sprintf("%v%vEn Passant: %v", boardString[7], rightSpaces, EPToString(cgs.en_passant))
	boardString[8] = fmt.Sprintf("%v%vHalfmove: %v", boardString[8], rightSpaces, cgs.halfmove)
	boardString[9] = fmt.Sprintf("%v%vFullmove: %v", boardString[9], rightSpaces, cgs.fullmove)

	for _, line := range boardString {
		fmt.Println(line)
	}

	for i := 0; i < spaces; i++ {
		fmt.Printf("\n")
	}
}

func EPToString(ep int) string {
	if ep == -1 {
		return "-"
	} else {
		return indexToString(ep)
	}
}
