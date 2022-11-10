package engine

import (
	"fmt"
	"io/ioutil"
	"time"
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
	start := time.Now()
	fen := readFENFile(filepath)
	gs, err := FenLoader(fen)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
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
