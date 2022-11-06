package main

import (
	"chessbot-go/src/GameEngine/board"
	"fmt"
)

func isNumberEven(x int64) bool {
	return x%2 == 0
}

func main() {
	fmt.Println("Hello")
	fmt.Println(isNumberEven(32))

	b := board.NewBoard(board.Starting_board_fen_string)
	fmt.Println(b)
}
