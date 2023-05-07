package engine

import (
	"os"
)

type Gamestate struct {
	Board      *Board
	Player     Player
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
	bytes, err := os.ReadFile(filepath)
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
