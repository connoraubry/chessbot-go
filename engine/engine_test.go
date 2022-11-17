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
		PrintBitboard(newBoard.Kings | newBoard.Rooks)
		PrintBitboard(newBoard.WhitePieces)
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
