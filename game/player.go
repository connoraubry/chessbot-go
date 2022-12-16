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

	//flag channels
	TakeMoveChan chan int
	QuitChan     chan int
}

func NewPlayer(t PlayerType, e *engine.Engine) *Player {
	p := new(Player)
	p.Type = t
	p.Engine = e

	p.InputMoveChan = make(chan engine.Move)
	p.OutputMovechan = make(chan engine.Move)

	p.TakeMoveChan = make(chan int)

	p.QuitChan = make(chan int)

	return p
}

func (p *Player) Run() {
	var m engine.Move
	for {
		select {
		case <-p.TakeMoveChan:
			p.OutputMovechan <- p.GetMove()
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
			time.Sleep(1 * time.Second)
			p.Engine.Print()
			return m
		}
	}

	p.Engine.Print()
	fmt.Println(moves[randomIndex])
	fmt.Println(p.Engine.ExportToFEN())
	return defaultMove

}

func EvaluateBoard(board engine.Board, player engine.Player) int {
	score := 0
	playerPieces := board.PlayerPieces(player)
	score += 100 * CountBits(board.Pawns&playerPieces)
	score += 300 * CountBits(board.Bishops&playerPieces)
	score += 300 * CountBits(board.Knights&playerPieces)
	score += 500 * CountBits(board.Rooks&playerPieces)
	score += 800 * CountBits(board.Queens&playerPieces)

	return score
}

func CountBits(bb engine.Bitboard) int {
	res := 0

	for bb > 0 {
		bb.PopLSB()
		res += 1
	}
	return res
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
