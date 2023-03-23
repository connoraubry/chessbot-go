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

func (h *Human) GetMove() engine.Move {
	moves := h.Engine.GetAllMoves()

	stringToMove := make(map[string]engine.Move)
	stMoves := make([]string, len(moves))
	for idx, move := range moves {
		stringToMove[move.String()] = move
		stMoves[idx] = move.String()
	}

	h.Engine.Print()
	fmt.Println(stMoves)
	fmt.Println(h.Engine.ExportToFEN())

	stMovesNew := engine.GetStringToMoveMap(moves)
	fmt.Println("New st moves")

	all_moves := make([]string, len(stMovesNew))
	var i = 0
	for k := range stMovesNew {
		all_moves[i] = k
		i++
	}

	fmt.Println(all_moves)

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
