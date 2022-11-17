package engine

import (
	"testing"
)

func TestGetEnPassantFromString(t *testing.T) {
	ep_string := "-"
	value, _ := getEnPassantFromString(ep_string)
	expected := -1
	if value != expected {
		t.Fatalf(`getEnPassantFromString(%v) = %v. Expected %v`, ep_string, value, expected)
	}
	ep_string = "c1"
	value, _ = getEnPassantFromString(ep_string)
	expected = 2
	if value != expected {
		t.Fatalf(`getEnPassantFromString(%v) = %v. Expected %v`, ep_string, value, expected)
	}
}

func TestAlgebraicToInteger(t *testing.T) {
	algebraic := "a1"
	expected := 0
	value, err := algebraicToInteger(algebraic)
	if err != nil {
		t.Fatalf(`Unexpected error %v`, err)
	}
	if value != expected {
		t.Fatalf(`algebraicToInteger(%v) = %v. Expected %v`, algebraic, value, expected)
	}

	algebraic = "a2"
	expected = 8
	value, err = algebraicToInteger(algebraic)
	if err != nil {
		t.Fatalf(`Unexpected error %v`, err)
	}
	if value != expected {
		t.Fatalf(`algebraicToInteger(%v) = %v. Expected %v`, algebraic, value, expected)
	}
	algebraic = "b2"
	expected = 9
	value, err = algebraicToInteger(algebraic)
	if err != nil {
		t.Fatalf(`Unexpected error %v`, err)
	}
	if value != expected {
		t.Fatalf(`algebraicToInteger(%v) = %v. Expected %v`, algebraic, value, expected)
	}

	algebraic = "j2"
	expected = -1
	value, err = algebraicToInteger(algebraic)
	if err == nil {
		t.Fatalf(`Expected error out of bounds`)
	}
	if value != expected {
		t.Fatalf(`algebraicToInteger(%v) = %v. Expected %v`, algebraic, value, expected)
	}

	algebraic = "A4"
	expected = -1
	value, err = algebraicToInteger(algebraic)
	if err == nil {
		t.Fatalf(`Expected error out of bounds`)
	}
	if value != expected {
		t.Fatalf(`algebraicToInteger(%v) = %v. Expected %v`, algebraic, value, expected)
	}

	algebraic = "abc"
	expected = -1
	value, err = algebraicToInteger(algebraic)
	if err == nil {
		t.Fatalf(`Expected error out of bounds`)
	}
	if value != expected {
		t.Fatalf(`algebraicToInteger(%v) = %v. Expected %v`, algebraic, value, expected)
	}

}

func TestGetCastleFromString(t *testing.T) {
	input_string := "KQkq"
	expected := Castle{true, true, true, true}
	result, err := getCastleFromString(input_string)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if result != expected {
		t.Fatalf(`Castle = %v. Expected %v`, result, expected)
	}

	input_string = "KQ"
	expected = Castle{true, true, false, false}
	result, err = getCastleFromString(input_string)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if result != expected {
		t.Fatalf(`Castle = %v. Expected %v`, result, expected)
	}

	input_string = "Kk"
	expected = Castle{true, false, true, false}
	result, err = getCastleFromString(input_string)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if result != expected {
		t.Fatalf(`Castle = %v. Expected %v`, result, expected)
	}

	input_string = "-"
	expected = Castle{false, false, false, false}
	result, err = getCastleFromString(input_string)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if result != expected {
		t.Fatalf(`Castle = %v. Expected %v`, result, expected)
	}

}
func TestFenLoaderStart(t *testing.T) {
	gamestate, err := FenLoader(starting_fen)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if gamestate.en_passant != -1 {
		t.Fatalf(`En passant = %v. Expected %v`, gamestate.en_passant, -1)
	}
	if gamestate.Board.AllPieces().LSB() != 1 {
		t.Fatalf(`Board[0] = %v. Expected %v`, gamestate.Board.AllPieces().LSB(), 1)
	}
	if gamestate.player != WHITE {
		t.Fatalf(`gamestate.move == %v. Expected %v`, gamestate.player, WHITE)
	}
	expected_cs := Castle{true, true, true, true}
	if gamestate.castle != expected_cs {
		t.Fatalf(`gamestate.castle.blackKing == %v. Expected %v`, gamestate.castle.blackKing, true)
	}

	if gamestate.halfmove != 0 {
		t.Fatalf(`gamestate.halfmove = %v. Expected %v`, gamestate.halfmove, 0)
	}
	if gamestate.fullmove != 1 {
		t.Fatalf(`gamestate.fullmove = %v. Expected %v`, gamestate.halfmove, 1)
	}
}

func TestGetMoveFromString(t *testing.T) {
	input := "w"
	expected := WHITE
	value, err := getMoveFromString(input)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if value != expected {
		t.Fatalf(`getMoveFromString(%v) = %v. Expected %v`, input, value, expected)
	}

	input = "b"
	expected = BLACK
	value, err = getMoveFromString(input)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if value != expected {
		t.Fatalf(`getMoveFromString(%v) = %v. Expected %v`, input, value, expected)
	}

	input = "W"
	expected = WHITE
	value, err = getMoveFromString(input)
	if err == nil {
		t.Fatalf(`Expected error: getMoveFromString(%v)`, input)
	}
	if value != expected {
		t.Fatalf(`getMoveFromString(%v) = %v. Expected %v`, input, value, expected)
	}
}

func TestFenLoader_e4(t *testing.T) {

	e4_fen := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"

	gamestate, err := FenLoader(e4_fen)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if gamestate.en_passant != 20 {
		t.Fatalf(`En passant = %v. Expected %v`, gamestate.en_passant, 20)
	}

	if gamestate.Board.Pawns&Bitboard(1<<28) == 0 {
		t.Fatalf(`Board[0] = %v. Expected %v`, gamestate.Board.Pawns&Bitboard(1<<28), 1<<28)
	}

	if gamestate.player != BLACK {
		t.Fatalf(`gamestate.move == %v. Expected %v`, gamestate.player, BLACK)
	}
	expected_cs := Castle{true, true, true, true}
	if gamestate.castle != expected_cs {
		t.Fatalf(`gamestate.castle.blackKing == %v. Expected %v`, gamestate.castle.blackKing, true)
	}

	if gamestate.halfmove != 0 {
		t.Fatalf(`gamestate.halfmove = %v. Expected %v`, gamestate.halfmove, 0)
	}
	if gamestate.fullmove != 1 {
		t.Fatalf(`gamestate.fullmove = %v. Expected %v`, gamestate.halfmove, 1)
	}
}

func TestFenLoader_middlegame(t *testing.T) {

	fen := "r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R1B1K2R b KQ - 1 10"

	gamestate, err := FenLoader(fen)
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	if gamestate.en_passant != -1 {
		t.Fatalf(`En passant = %v. Expected %v`, gamestate.en_passant, -1)
	}

	if gamestate.player != BLACK {
		t.Fatalf(`gamestate.move == %v. Expected %v`, gamestate.player, BLACK)
	}
	expected_cs := Castle{true, true, false, false}
	if gamestate.castle != expected_cs {
		t.Fatalf(`gamestate.castle.blackKing == %v. Expected %v`, gamestate.castle.blackKing, true)
	}

	if gamestate.halfmove != 1 {
		t.Fatalf(`gamestate.halfmove = %v. Expected %v`, gamestate.halfmove, 1)
	}
	if gamestate.fullmove != 10 {
		t.Fatalf(`gamestate.fullmove = %v. Expected %v`, gamestate.halfmove, 10)
	}
}

func TestIntegerToAlgebraic(t *testing.T) {

	var mappings = map[int]string{
		0:  "a1",
		1:  "b1",
		7:  "h1",
		8:  "a2",
		56: "a8",
		63: "h8",
	}

	for idx, expected := range mappings {
		if integerToAlgebraic(idx) != expected {
			t.Fatalf(`iTA(%v) == %v.  Expected %v`, idx, integerToAlgebraic(idx), expected)
		}
	}
}

func TestGetStringFromEnPassant(t *testing.T) {
	var mappings = map[int]string{
		0:  "a1",
		1:  "b1",
		7:  "h1",
		8:  "a2",
		56: "a8",
		63: "h8",
		-1: "-",
	}

	for idx, expected := range mappings {
		if getStringFromEnPassant(idx) != expected {
			t.Fatalf(`iTA(%v) == %v.  Expected %v`, idx, getStringFromEnPassant(idx), expected)
		}
	}
}
func TestGetStringFromMove(t *testing.T) {
	var mappings = map[Player]string{
		WHITE:     "w",
		BLACK:     "b",
		NO_PLAYER: "-",
	}

	for idx, expected := range mappings {
		if getStringFromPlayer(idx) != expected {
			t.Fatalf(`iTA(%v) == %v.  Expected %v`, idx, getStringFromPlayer(idx), expected)
		}
	}
}

func TestExportToFENGamestate(t *testing.T) {

	fen := "r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R1B1K2R b KQ - 1 10"

	gamestate, err := FenLoader(fen)
	if err != nil {
		t.Fatalf("%v", err)
	}
	export := ExportToFEN(gamestate)

	if export != fen {
		t.Fatalf(`%v != %v`, export, fen)
	}
}
