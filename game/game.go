package game

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

const FEN_Start string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Game struct {
	PlayerWhite Player
	PlayerBlack Player

	engine   engine.Engine
	moveList []engine.Move
}

func NewGame() *Game {
	g := new(Game)

	g.PlayerWhite = *NewPlayer(
		AUTOMATON,
		engine.NewEngine(engine.OptFenString(FEN_Start)),
	)
	g.PlayerBlack = *NewPlayer(
		AUTOMATON,
		engine.NewEngine(engine.OptFenString(FEN_Start)),
	)

	g.engine = *engine.NewEngine(engine.OptFenString(FEN_Start))

	return g
}

func (g *Game) Run() {
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		g.Quit()
	}()

	go g.PlayerWhite.Run()
	go g.PlayerBlack.Run()

	g.PlayerWhite.TakeMoveChan <- 1

	g.loop()
	time.Sleep(1 * time.Second)
}

func (g *Game) Quit() {
	g.PlayerWhite.QuitChan <- 1
	g.PlayerBlack.QuitChan <- 1
	fmt.Println(g.moveList)
	os.Exit(1)
}

func (g *Game) loop() {
	m := engine.Move{}
	for {
		select {
		case m = <-g.PlayerWhite.OutputMovechan:
			g.moveList = append(g.moveList, m)
			g.engine.TakeMove(m)

			g.PlayerBlack.InputMoveChan <- m
		case m = <-g.PlayerBlack.OutputMovechan:
			g.moveList = append(g.moveList, m)
			g.engine.TakeMove(m)
			g.PlayerWhite.InputMoveChan <- m
		}
	}
}
