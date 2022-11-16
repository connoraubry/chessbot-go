package engine

import (
	"io/ioutil"
)

type Castle struct {
	whiteKing  bool
	whiteQueen bool
	blackKing  bool
	blackQueen bool
}

func (cs *Castle) Copy() Castle {
	newCS := Castle{
		whiteKing:  cs.whiteKing,
		blackKing:  cs.blackKing,
		blackQueen: cs.blackQueen,
		whiteQueen: cs.whiteQueen,
	}
	return newCS
}

type Gamestate struct {
	Board      *Board
	player     Player
	castle     Castle
	en_passant int
	halfmove   int
	fullmove   int
}

func NewGamestateFile(filepath string) *Gamestate {
	fen := readFENFile(filepath)
	return NewGamestateFEN(fen)
}

func NewGamestateFEN(fen string) *Gamestate {
	gs, err := FenLoader(fen)
	if err != nil {
		panic(err)
	}
	return gs
}

func readFENFile(filepath string) string {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (gs *Gamestate) EnPassantBitboard() Bitboard {
	if gs.en_passant > -1 {
		return Bitboard(1 << gs.en_passant)
	}
	return Bitboard(0)
}

func (gs *Gamestate) PrintBoard() {
	gs.Board.PrintBoard()
}

func (cs *Castle) ToString() string {
	var v []byte
	if cs.whiteKing {
		v = append(v, 'K')
	}
	if cs.whiteQueen {
		v = append(v, 'Q')
	}
	if cs.blackKing {
		v = append(v, 'k')
	}
	if cs.blackQueen {
		v = append(v, 'q')
	}
	return string(v)
}
