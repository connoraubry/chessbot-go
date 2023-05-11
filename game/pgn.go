package game

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func (g *Game) ExportToPGN(filepath string) error {
	var err error

	output := g.createPGNStringList()
	file, err := os.Create(filepath)

	if err != nil {
		return err
	}
	defer file.Close()

	datawriter := bufio.NewWriter(file)

	for _, data := range output {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()

	return err
}

func (g *Game) createPGNStringList() []string {
	var output []string
	output = append(output, "[Event \"?\" ]")
	output = append(output, "[Site \"chessbot\" ]")
	output = append(output, fmt.Sprintf("[Date \"%v\" ]", getPGNDate()))
	output = append(output, "[Round \"?\" ]")

	output = append(output, fmt.Sprintf("[White \"%v\" ]", playerToString(g.WhiteType)))
	output = append(output, fmt.Sprintf("[Black \"%v\" ]", playerToString(g.BlackType)))
	output = append(output, "[Result \"*\" ]")
	output = append(output, "")

	var first string
	var second string

	for i := 0; i < len(g.moveList); i += 2 {
		idx := (i / 2) + 1
		first = g.moveList[i]

		if i+1 < len(g.moveList) {
			second = g.moveList[i+1]
		} else {
			second = ""
		}

		output = append(output, fmt.Sprintf("%v. %v %v", idx, first, second))
	}

	return output
}

func playerToString(player PlayerType) string {
	var humanString string
	switch player {
	case HUMAN:
		humanString = "Human"
	case AUTOMATON:
		humanString = "Bot"
	}
	return humanString
}

func getPGNDate() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d.%d.%d", currentTime.Year(), currentTime.Month(), currentTime.Day())
}
