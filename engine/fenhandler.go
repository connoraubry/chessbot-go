package engine

import (
	"fmt"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func FenLoader(FEN string) (*Gamestate, error) {
	s := strings.Split(FEN, " ")
	board := NewBoard(s[0])
	move, move_err := getMoveFromString(s[1])
	handle(move_err)
	castle, c_err := getCastleFromString(s[2])
	handle(c_err)
	ep, ep_err := getEnPassantFromString(s[3])
	handle(ep_err)

	halfmove := 0
	fullmove := 0

	if len(s) > 4 {
		var h_err, f_err error
		halfmove, h_err = strconv.Atoi(s[4])
		handle(h_err)
		fullmove, f_err = strconv.Atoi(s[5])
		handle(f_err)
	}

	gs := &Gamestate{
		Board:      board,
		Player:     move,
		castle:     castle,
		en_passant: ep,
		halfmove:   halfmove,
		fullmove:   fullmove,
	}

	return gs, nil
}

func getMoveFromString(move_string string) (Player, error) {
	switch move_string {
	case ("w"):
		return WHITE, nil
	case ("b"):
		return BLACK, nil
	default:
		return WHITE, fmt.Errorf("invalid move string: %v", move_string)
	}
}

func getStringFromPlayer(p Player) string {
	switch p {
	case WHITE:
		return "w"
	case BLACK:
		return "b"
	}
	return "-"
}

func getCastleFromString(castle_string string) (Castle, error) {
	cs := Castle{}
	for _, letter := range castle_string {
		switch letter {
		case 'K':
			cs.whiteKing = true
		case 'Q':
			cs.whiteQueen = true
		case 'k':
			cs.blackKing = true
		case 'q':
			cs.blackQueen = true
		default:
			continue
		}
	}
	return cs, nil
}

func getEnPassantFromString(ep_string string) (int, error) {
	var ep int
	var err error
	switch ep_string {
	case "-":
		ep = -1
	default:
		ep, err = algebraicToInteger(ep_string)
		if err != nil {
			return -1, err
		}
	}

	return ep, nil
}

func getStringFromEnPassant(ep_index int) string {
	if ep_index == -1 {
		return "-"
	} else {
		return integerToAlgebraic(ep_index)
	}
}

func algebraicToInteger(algebraic string) (int, error) {
	if len(algebraic) != 2 {
		return -1, fmt.Errorf("invalid algebraic string: %v", algebraic)
	}
	file := algebraic[0] - 'a'
	rank := algebraic[1] - '1'

	if file > 7 || rank > 7 {
		fmt.Println(file, rank)
		return -1, fmt.Errorf("invalid algebraic string: %v", algebraic)
	}

	return int(file + rank*8), nil
}

func integerToAlgebraic(index int) string {

	file := index & 7
	rank := index >> 3

	fileB := byte(file) + 'a'
	rankB := byte(rank) + '1'

	return fmt.Sprintf("%c%c", fileB, rankB)

}

func ExportToFEN(gs *Gamestate) string {

	fen := fmt.Sprintf("%v %v %v %v %v %v",
		gs.Board.ExportToFEN(),
		getStringFromPlayer(gs.Player),
		gs.castle.ToString(),
		getStringFromEnPassant(gs.en_passant),
		gs.halfmove,
		gs.fullmove)

	return fen
}

func ExportToFENNoMoves(gs *Gamestate) string {
	return fmt.Sprintf("%v %v %v %v",
		gs.Board.ExportToFEN(),
		getStringFromPlayer(gs.Player),
		gs.castle.ToString(),
		getStringFromEnPassant(gs.en_passant))
}
