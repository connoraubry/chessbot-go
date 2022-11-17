package engine

import "testing"

var moves []Move

var (
	scenario_one_fen   = "4B1r1/P1b2PpQ/3R4/3p4/1P1B4/1p3k2/1N1K4/N7 w - - 0 1"
	scenario_seven_fen = "r2qkb1r/p1pp2pp/bpB2p1n/4p3/4P2N/5P2/PPPP2PP/RNBQK2R w KQkq - 1 7"
)

func TestScenarioSeven(t *testing.T) {
	gs := NewGamestateFEN(scenario_seven_fen)
	var moves = []string{
		"Bxa8", "Bb7", "Bd5", "Bxd7", //technically check
		"Bb5", "Ba4", "a3", "a4", "b3",
		"b4", "c3", "c4", "d3", "d4",
		"f4", "g3", "g4", "h3", "Ng6",
		"Nf5", "Nc3", "Na3", "Qe2", "Kf2",
		"Rg1", "Rf1",
		"Ke2", "Kf1", "O-O", // these are invalid moves, psuedo-legal
	}
	var movesFound = make(map[string]bool)
	for _, str := range moves {
		movesFound[str] = false
	}

	allMoves := gs.GetAllMoves()

	for _, move := range allMoves {
		_, ok := movesFound[move.String()]
		if ok {
			movesFound[move.String()] = true
		} else {
			t.Fatalf(`Move %v not in supplied list.`, move.String())
		}
	}

	for move, result := range movesFound {
		if !result {
			t.Fatalf(`Did not find move %v.`, move)
		}
	}
}

func BenchmarkScenarioOne(b *testing.B) {
	var result []Move
	for i := 0; i < b.N; i++ {
		gs := NewGamestateFEN(scenario_one_fen)
		result = gs.GetAllMoves()
	}
	moves = result
}

func BenchmarkScenarioSeven(b *testing.B) {
	var result []Move
	for i := 0; i < b.N; i++ {
		gs := NewGamestateFEN(scenario_seven_fen)
		result = gs.GetAllMoves()
	}
	moves = result
}

func TestScenarioOne(t *testing.T) {
	gs := NewGamestateFEN("4B1r1/P1b2PpQ/3R4/3p4/1P1B4/1p3k2/1N1K4/N7 w - - 0 1")
	var moves = []string{
		"a8=Q", "a8=B", "a8=N", "a8=R", "Bd7",
		"Bc6", "Bb5", "Ba4", "f8=B", "f8=N",
		"fxg8=Q", "fxg8=B", "fxg8=N", "fxg8=R",
		"Qxg8", "Qh8", "Qxg7", "Qh6", "Qh4", "Qh2",
		"Qg6", "Qc2", "Qb1", "Rd8", "Rd7", "Rxd5",
		"Ra6", "Rb6", "Rc6", "Re6", "Rg6", "Rh6",
		"Bb6", "Bc5", "Be5", "Bf6", "Bxg7", "Bc3",
		"Be3", "Bf2", "Bg1", "b5", "Nxb3", "Nc2",
		"Na4", "Nc4", "Nd3", "Nd1", "Kc3", "Kd3",
		"Kc1", "Ke1", "Kd1",

		//Next ones are checks. Should have a plus for legal moves
		"f8=Q", "Rf6", "Qf5", "Qe4", "Qd3",
		"Qh5", "Qh1", "Qh3", "f8=R",

		"Ke2", "Ke3", "Kc2", // these are invalid moves, psuedo-legal
	}
	var movesFound = make(map[string]bool)
	for _, str := range moves {
		movesFound[str] = false
	}

	allMoves := gs.GetAllMoves()

	for _, move := range allMoves {
		_, ok := movesFound[move.String()]
		if ok {
			movesFound[move.String()] = true
		} else {
			t.Fatalf(`Move %v not in supplied list.`, move.String())
		}
	}

	for move, result := range movesFound {
		if !result {
			t.Fatalf(`Did not find move %v.`, move)
		}
	}
}

func TestScenarioTwo(t *testing.T) {
	gs := NewGamestateFEN("1R1b1k2/r6B/NPnp4/p2r4/8/2Pp2Pp/8/7K b - - 0 1")
	var moves = []string{
		"Rxa6", "Ra8", "Ke8", "Nd4", "Ke7",
		"Rg5", "Kg7", "Rc7", "Rh5", "Re7",
		"h2", "Rxh7", "Rg7", "a4", "Ne7",
		"Re5", "Nxb8", "Rf7", "Nb4", "Rb7",
		"Rb5", "Rf5", "Kf7", "Rd4", "Rc5",
		"Rd7", "d2", "Ne5",

		//pseudo legal
		"Kg8", "Bc7", "Bxb6", "Be7", "Bf6", "Bg5", "Bh4",
	}
	var movesFound = make(map[string]bool)
	for _, str := range moves {
		movesFound[str] = false
	}

	allMoves := gs.GetAllMoves()

	for _, move := range allMoves {
		_, ok := movesFound[move.String()]
		if ok {
			movesFound[move.String()] = true
		} else {
			t.Fatalf(`Move %v not in supplied list.`, move.String())
		}
	}

	for move, result := range movesFound {
		if !result {
			t.Fatalf(`Did not find move %v.`, move)
		}
	}
}

func TestScenarioThree(t *testing.T) {
	e := NewEngine(OptFenString("1N5R/2p2PK1/5n1P/5N2/1PQ1P3/8/q4k1p/3r3R w - - 0 1"))

	var moves = []string{
		"Na6", "Nc6", "Nd7", "Rh7",
		"Rg8", "Rf8", "Re8", "Rd8",
		"Rc8", "f8=Q", "f8=R", "f8=N",
		"f8=B", "Kf8", "Kxf6", "Kg6",
		"h7", "Ne7", "Nd6", "Nd4", "Ne3",
		"Ng3", "Nh4", "e5", "b5", "Qxc7",
		"Qc6", "Qc3", "Qc1", "Qb3", "Qd3",
		"Qb5", "Qe6", "Qd5", "Qb5",
		"Qa6", "Rg1", "Re1", "Rxd1",

		//checks
		"Rf1", "Rxh2",
		"Qe2", "Qf1", "Qxa2",
		"Qc2", "Qc5", "Qd4",
	}

	var movesFound = make(map[string]bool)
	for _, str := range moves {
		movesFound[str] = false
	}

	allMoves := e.GetAllMoves()

	lenValidMoves := 0

	for _, move := range allMoves {
		success := e.TakeMove(move)
		if success {
			_, ok := movesFound[move.String()]
			if ok {
				movesFound[move.String()] = true
			} else {
				t.Fatalf(`Move %v not in supplied list.`, move.String())
			}
			lenValidMoves += 1
			e.UndoMove()
		}

	}

	for move, result := range movesFound {
		if !result {
			t.Fatalf(`Did not find move %v.`, move)
		}
	}
}
