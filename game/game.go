package main

import (
	"bufio"
	"chessbot-go/engine"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

func main() {
	flag.Parse()

	e := engine.NewEngine(
		// engine.OptFenFile(*fen),
		engine.OptFenString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
	)

	for i := 0; i < 1000; i++ {
		moves := e.GetAllMoves()
		stringToMove := make(map[string]engine.Move)
		stMoves := make([]string, len(moves))
		for idx, move := range moves {
			stringToMove[move.String()] = move
			stMoves[idx] = move.String()
		}

		e.Print()
		fmt.Println(stMoves)
		fmt.Println(e.ExportToFEN())

		for j := 0; j < 3; j++ {
			playerInput := GetPlayerInput(stringToMove)
			success := e.TakeMove(playerInput)
			if success {
				break
			}
		}

	}

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

func Perft(e *engine.Engine, depth int) int {
	count := 0
	if depth == 1 {
		for _, m := range e.GetAllMoves() {
			res := e.TakeMove(m)
			if res {
				count += 1
				e.UndoMove()
			}
		}
		return count
	}
	for _, move := range e.GetAllMoves() {
		res := e.TakeMove(move)
		if res {
			count += Perft(e, depth-1)
			e.UndoMove()
		}

	}
	return count

}
