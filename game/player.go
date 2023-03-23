package game

import (
	"chessbot-go/engine"
	"fmt"
)

type PlayerType int

const (
	NO_PLAYER PlayerType = iota
	HUMAN
	AUTOMATON
)

// what methods do all players have?
type Player interface {
	Run()
	GetMove() engine.Move
	Quit()

	Update(engine.Move)
}

func NewPlayer(t PlayerType, color engine.Player, e *engine.Engine) (Player, error) {
	switch t {
	case HUMAN:
		return NewHuman(e), nil
	case AUTOMATON:
		return NewAutomaton(e, color), nil
	default:
		return nil, fmt.Errorf("new player must be human or automaton")
	}
}
