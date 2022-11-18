package game

import (
	"chessbot-go/engine"
	"flag"
	"time"
)

var (
// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

func Run() {
	flag.Parse()

	Player1 := NewPlayer(
		HUMAN,
		engine.NewEngine(
			engine.OptFenString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
		),
		make(chan engine.Move),
		make(chan engine.Move),
		make(chan int),
	)
	Player2 := NewPlayer(
		AUTOMATON,
		engine.NewEngine(
			engine.OptFenString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
		),
		make(chan engine.Move),
		make(chan engine.Move),
		make(chan int),
	)

	m := Player1.GetMove()

	go Player1.Run()
	go Player2.Run()

	Player2.InputMoveChan <- m

	run(Player1, Player2)
	time.Sleep(1 * time.Second)

}

func run(Player1, Player2 *Player) {
	quitMove := engine.Move{}
	m := engine.Move{}
	for {
		select {
		case m = <-Player1.OutputMovechan:
			if m == quitMove {
				Player1.QuitChan <- 1
				Player2.QuitChan <- 1
				return
			}
			Player2.InputMoveChan <- m

		case m = <-Player2.OutputMovechan:
			if m == quitMove {
				Player1.QuitChan <- 1
				Player2.QuitChan <- 1
				return
			}
			Player1.InputMoveChan <- m
		}
	}
}

// func Perft(e *engine.Engine, depth int) int {
// 	count := 0
// 	if depth == 1 {
// 		for _, m := range e.GetAllMoves() {
// 			res := e.TakeMove(m)
// 			if res {
// 				count += 1
// 				e.UndoMove()
// 			}
// 		}
// 		return count
// 	}
// 	for _, move := range e.GetAllMoves() {
// 		res := e.TakeMove(move)
// 		if res {
// 			count += Perft(e, depth-1)
// 			e.UndoMove()
// 		}

// 	}
// 	return count

// }
