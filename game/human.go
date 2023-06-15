package game

import (
	"bufio"
	"chessbot-go/engine"
	"fmt"
	"os"
	"strings"
)

type Human struct {
	Engine *engine.Engine

	InputMoveChan  chan engine.Move
	OutputMovechan chan engine.Move

	//flag channels
	TakeMoveChan chan int
	QuitChan     chan int
}

func NewHuman(e *engine.Engine) *Human {
	h := new(Human)
	h.Engine = e
	return h
}

func (h *Human) Update(move engine.Move) {
	h.Engine.TakeMove(move)
}

func (h *Human) Quit() {
}

func (h *Human) Run() {
}

func (h *Human) Dump() {}

func (h *Human) GetMove() engine.Move {
	moves := h.Engine.GetAllMoves()

	stringToMove := h.Engine.GetStringToMoveMap(moves)
	movesStringList := make([]string, len(stringToMove))
	var i = 0
	for k := range stringToMove {
		movesStringList[i] = k
		i++
	}

	h.Engine.Print()
	fmt.Println(movesStringList)
	fmt.Println(h.Engine.ExportToFEN())

	for j := 0; j < 3; j++ {
		playerInput := GetPlayerInput(stringToMove)

		success := h.Engine.TakeMove(playerInput)
		if success {
			return playerInput
		}
	}
	return engine.Move{}
}

func GetPlayerInput(stringToMove map[string]engine.Move) engine.Move {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		move, ok := stringToMove[text]
		if ok {
			return move
		}
	}
	return engine.Move{}
}
