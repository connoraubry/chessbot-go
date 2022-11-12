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
