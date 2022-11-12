package engine

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
)

type Engine struct {
	//'constants'
	SlidingBitrow map[Bitrow]map[Bitrow]Bitrow
	Constants     Constants
	opts          Options
	GameState     Gamestate
}

func NewEngine(opts ...interface{}) *Engine {

	e := &Engine{}
	var err error
	e.opts, err = ParseOptions(opts...)
	if err != nil {
		panic(err)
	}
	e.Constants = *loadConstants()
	e.LoadBitboards()

	e.GameState = *NewGamestateFile(e.opts.FenFilePath)

	return e
}

// func (e *Engine) GetAllMoves() []Moves{

// }

// func (e *Engine) TakeMove(Move) {

// }

// func (e *Engine) UndoMove() {

// }

func (e *Engine) LoadBitboards() {
	e.SlidingBitrow = make(map[Bitrow]map[Bitrow]Bitrow)

	sliding_file := filepath.Join(e.opts.BitboardDirectory, "SlidingBitrows256.csv")

	e.LoadSlidingBitrows(sliding_file)
}

func (e *Engine) LoadSlidingBitrows(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, entry := range records {
		matchRow, err := strconv.ParseInt(entry[0], 2, 64)
		if err != nil {
			panic(err)
		}

		matchRowBits := Bitrow(matchRow)
		e.SlidingBitrow[matchRowBits] = make(map[Bitrow]Bitrow)

		for index := 0; index < 8; index++ {
			matchResult, err := strconv.ParseInt(entry[index+1], 2, 64)
			// fmt.Println(matchResult, index, 1<<index)
			if err != nil {
				panic(err)
			}
			e.SlidingBitrow[matchRowBits][Bitrow(1<<index)] = Bitrow(matchResult)
		}
	}
}
