package automaton

import "chessbot-go/engine"

type Player struct {
	Engine *engine.Engine

	InputMoveChan  chan engine.Move
	OutputMovechan chan engine.Move

	//flag channels
	TakeMoveChan chan int
	QuitChan     chan int
}
