package game

import (
	"bufio"
	"chessbot-go/engine"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type PlayerType int

const (
	NO_PLAYER PlayerType = iota
	HUMAN
	AUTOMATON
)

type Player struct {
	Type PlayerType

	Engine *engine.Engine

	InputMoveChan  chan engine.Move
	OutputMovechan chan engine.Move

	QuitChan chan int
}

func NewPlayer(t PlayerType, e *engine.Engine, in, out chan engine.Move, quit chan int) *Player {
	p := new(Player)
	p.Type = t
	p.Engine = e

	p.InputMoveChan = in
	p.OutputMovechan = out

	p.QuitChan = quit

	return p
}

func (p *Player) Run() {
	var m engine.Move
	for {
		select {
		case m = <-p.InputMoveChan:
			p.Engine.TakeMove(m)
			p.OutputMovechan <- p.GetMove()
		case <-p.QuitChan:
			fmt.Println("quit")
			return
		}
	}
}

func (p *Player) GetMove() engine.Move {
	var m engine.Move
	switch p.Type {
	case HUMAN:
		return p.GetHumanMove()
	case AUTOMATON:
		return p.GetAutomatonMove()
	}
	return m
}

func (p *Player) GetAutomatonMove() engine.Move {
	defaultMove := engine.Move{}
	moves := p.Engine.GetAllMoves()
	randomIndex := rand.Intn(len(moves))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	for _, m := range moves {
		success := p.Engine.TakeMove(m)
		if success {
			return m
		}
	}

	p.Engine.Print()
	fmt.Println(moves[randomIndex])
	fmt.Println(p.Engine.ExportToFEN())
	return defaultMove

}

func (p *Player) GetHumanMove() engine.Move {

	moves := p.Engine.GetAllMoves()
	stringToMove := make(map[string]engine.Move)
	stMoves := make([]string, len(moves))
	for idx, move := range moves {
		stringToMove[move.String()] = move
		stMoves[idx] = move.String()
	}

	p.Engine.Print()
	fmt.Println(stMoves)
	fmt.Println(p.Engine.ExportToFEN())

	for j := 0; j < 3; j++ {
		playerInput := GetPlayerInput(stringToMove)
		success := p.Engine.TakeMove(playerInput)
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
