package game

import (
	"chessbot-go/engine"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
)

var (
// fen = flag.String("fen", "./FEN_configs/start.fen", "Starting FEN filepath")
)

const FEN_Start string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Game struct {
	PlayerWhite Player
	PlayerBlack Player

	WhiteType PlayerType
	BlackType PlayerType

	engine   engine.Engine
	moveList []string
}

func NewGame() *Game {
	var err error

	g := new(Game)

	g.engine = *engine.NewEngine(engine.OptFenString(FEN_Start))

	whiteEngine := engine.NewEngine(engine.OptFenString(FEN_Start))
	blackEngine := engine.NewEngine(engine.OptFenString(FEN_Start))

	g.PlayerWhite, err = NewPlayer(HUMAN, engine.WHITE, whiteEngine)
	g.WhiteType = HUMAN
	if err != nil {
		log.Fatal(err)
	}

	g.PlayerBlack, err = NewPlayer(AUTOMATON, engine.BLACK, blackEngine)
	g.BlackType = AUTOMATON
	if err != nil {
		log.Fatal(err)
	}

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

	g.loop()
}

func (g *Game) Quit() {
	fmt.Println("")
	g.PlayerWhite.Quit()
	g.PlayerBlack.Quit()
	fmt.Println(g.moveList)
	g.ExportToPGN("pgnfile.txt")
	g.PlayerBlack.Dump()
	os.Exit(0)
}

func (g *Game) loop() {
	var m engine.Move
	for {
		if len(g.engine.CurrentGamestate().GetAllMoves()) == 0 {
			g.Quit()
		}
		if g.engine.CurrentGamestate().Player == engine.WHITE {
			m = g.PlayerWhite.GetMove()

			stringMove := g.engine.GetMoveString(m, g.engine.GetValidMoves())
			g.moveList = append(g.moveList, stringMove)

			g.engine.TakeMove(m)
			g.PlayerBlack.Update(m)
		} else {
			m = g.PlayerBlack.GetMove()

			stringMove := g.engine.GetMoveString(m, g.engine.GetValidMoves())
			g.moveList = append(g.moveList, stringMove)

			g.engine.TakeMove(m)
			g.PlayerWhite.Update(m)
		}
	}
}
