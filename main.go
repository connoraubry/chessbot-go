package main

import "github.com/connoraubry/chessbot-go/game"

var (
// fen = flag.String("fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "FEN string")
)

//func main() {
//
//	//g := game.NewGame()
//	//g.Run()
//
//	e := engine.NewEngine(engine.OptFenString("r1bqkbnr/1ppp1ppp/p1B5/4p3/4P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 0 4"))
//	m := e.GetStringToMoveMap(e.GetValidMoves())
//	fmt.Println(m)
//	e.Print(0, "e4")
//
//}

func main() {
	g := game.NewGame()
	g.RunWithOutput()
}
