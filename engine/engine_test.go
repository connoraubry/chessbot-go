package engine

import "testing"

func TestTakeCastleWOO(t *testing.T) {

	e := NewEngine(
		OptFenString("r1bqkbnr/ppp2ppp/2np4/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 0 4"),
	)

	newBoard, success := e.TakeCastle(Move{Castle: WHITEOO})

	if newBoard.Kings&newBoard.Queens&newBoard.Rooks&newBoard.Knights&newBoard.Bishops&newBoard.Pawns > 0 {
		t.Fatalf(`newBoard has more than one piece per spot`)
	}

	if newBoard.WhitePieces&Bitboard(0b10010000) > 1 {
		PrintBitboard(newBoard.Kings | newBoard.Rooks)
		t.Fatalf(`newBoard has more than one piece in king/rook spot`)

	}

	if !success {
		t.Fatalf(`Error`)
	}

}

func TestTakeCastleBOO(t *testing.T) {

	e := NewEngine(
		OptFenString("r2qk2r/pppbbppp/2np1n2/4p3/2B1P3/2NP1N2/PPPBQPPP/R3K2R b KQkq - 6 7"),
	)

	newBoard, success := e.TakeCastle(Move{Castle: BLACKOO})

	if newBoard.Kings&newBoard.Queens&newBoard.Rooks&newBoard.Knights&newBoard.Bishops&newBoard.Pawns > 0 {
		t.Fatalf(`newBoard has more than one piece per spot`)
	}
	if !success {
		t.Fatalf(`Error`)
	}
}

func TestTakeCastleWOOO(t *testing.T) {

	e := NewEngine(
		OptFenString("r2qk2r/ppp1bppp/2npbn2/4p3/2B1P3/2NP1N2/PPPBQPPP/R3K2R w KQkq - 7 8"),
	)

	newBoard, success := e.TakeCastle(Move{Castle: WHITEOOO})

	if newBoard.Kings&newBoard.Queens&newBoard.Rooks&newBoard.Knights&newBoard.Bishops&newBoard.Pawns > 0 {
		t.Fatalf(`newBoard has more than one piece per spot`)
	}
	if newBoard.WhitePieces&Bitboard(0b10001) > 1 {
		t.Fatalf(`newBoard has more than one piece in king/rook spot`)

	}
	if !success {
		t.Fatalf(`Error`)
	}
}

func TestTakeCastleBOOO(t *testing.T) {

	e := NewEngine(
		OptFenString("r3k2r/pppqbppp/2npbn2/4p3/2B1P3/1PNP1NP1/P1PBQP1P/R3K2R b KQkq - 0 9"),
	)
	newBoard, success := e.TakeCastle(Move{Castle: WHITEOOO})

	if newBoard.Kings&newBoard.Queens&newBoard.Rooks&newBoard.Knights&newBoard.Bishops&newBoard.Pawns > 0 {
		t.Fatalf(`newBoard has more than one piece per spot`)
	}
	if !success {
		t.Fatalf(`Error`)
	}
}

// func TestGetAllStrings(t *testing.T) {
// 	fen := "3b1q1R/2kp4/3p1q1q/4pp1p/4pPp1/8/8/5K2 b - f3 0 1"
// 	e := NewEngine(
// 		OptFenString(fen),
// 	)
// 	allMoves := e.GetAllMoves()

// 	expected_moves = []string{
// 		"Kb8", "Kc8", "Kb7", "Kb6", "Kc6",
// 		"d5", "e3", "exf3", "g3", "gxf3",
// 		"h4", "exf4", "Be7",
// 		"Qe8", "Qg8", "Q8xh8", "Q8f7",
// 		"Q8g7", "Q8e7", "Qfg5", "Qfg6",
// 		"Qh4", "Qe6", "Q6e7", "Q6f7",
// 		"Qf6g7", "Qf6xh8", "Qhg5", "Qh7", "Qhg7",
// 		"Qhg6", "Qhxh8",

// 		//checks
// 		"Qxf4",
// 	}

// }

// func TestScenarioOneValid(t *testing.T) {
// 	e := NewEngine(
// 		OptFenString("4B1r1/P1b2PpQ/3R4/3p4/1P1B4/1p3k2/1N1K4/N7 w - - 0 1"),
// 	)

// 	var moves = []string{
// 		"a8=Q", "a8=B", "a8=N", "a8=R", "Bd7",
// 		"Bc6", "Bb5", "Ba4", "f8=B", "f8=N",
// 		"fxg8=Q", "fxg8=B", "fxg8=N", "fxg8=R",
// 		"Qxg8", "Qh8", "Qxg7", "Qh6", "Qh4", "Qh2",
// 		"Qg6", "Qc2", "Qb1", "Rd8", "Rd7", "Rxd5",
// 		"Ra6", "Rb6", "Rc6", "Re6", "Rg6", "Rh6",
// 		"Bb6", "Bc5", "Be5", "Bf6", "Bxg7", "Bc3",
// 		"Be3", "Bf2", "Bg1", "b5", "Nxb3", "Nc2",
// 		"Na4", "Nc4", "Nd3", "Nd1", "Kc3", "Kd3",
// 		"Kc1", "Ke1", "Kd1",

// 		//Next ones are checks. Should have a plus for legal moves
// 		"f8=Q", "Rf6", "Qf5", "Qe4", "Qd3",
// 		"Qh5", "Qh1", "Qh3", "f8=R",

// 		"Ke2", "Ke3", "Kc2", // these are invalid moves, psuedo-legal
// 	}
// 	var movesFound = make(map[string]bool)
// 	for _, str := range moves {
// 		movesFound[str] = false
// 	}

// 	allMoves := e.GetValidMoves()

// 	for _, move := range allMoves {
// 		_, ok := movesFound[move.String()]
// 		if ok {
// 			movesFound[move.String()] = true
// 		} else {
// 			t.Fatalf(`Move %v not in supplied list.`, move.String())
// 		}
// 	}

// 	for move, result := range movesFound {
// 		if !result {
// 			t.Fatalf(`Did not find move %v.`, move)
// 		}
// 	}
// }
