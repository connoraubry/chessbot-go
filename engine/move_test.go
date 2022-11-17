package engine

import (
	"testing"
)

//for benchmarks
var result []Move

func TestGetAllKnightMoves(t *testing.T) {

	gs := NewGamestateFEN(starting_fen)

	moves := gs.GetAllKnightMoves()
	expected_length_moves := 4
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestGetPawnOneMoves(t *testing.T) {
	gs := NewGamestateFEN(starting_fen)

	moves := gs.GetAllPawnMoves()
	expected_length_moves := 16
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestPawnAttack(t *testing.T) {
	gs := NewGamestateFEN("8/8/6p1/5P2/8/8/8/8 w - - 0 1")
	moves := gs.GetAllPawnMoves()
	expected_length_moves := 2
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}

}

func TestPawnAttackAndPromotion(t *testing.T) {
	gs := NewGamestateFEN("5pp1/5P2/8/8/8/8/8/8 w - - 0 1")
	moves := gs.GetAllPawnMoves()
	expected_length_moves := 4
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestPawnAttackSid(t *testing.T) {
	gs := NewGamestateFEN("8/p7/6p1/7P/8/8/8/8 w - - 0 1")
	moves := gs.GetAllPawnAttackMoves()
	expected_length_moves := 1
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestGetRookMoves(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppppppp/8/8/8/P7/PPPPPPPP/R3KBNR w KQkq - 0 1")
	moves := gs.GetAllRookMoves()
	expected_length_moves := 3
	if len(moves) != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, len(moves), expected_length_moves)
	}
}

func TestGetVerticalMoves(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppppppp/8/P7/8/8/R2PPPPP/4KBNR w KQkq - 0 1")
	gs.Board.Rooks.PopLSB()
	lsb := gs.Board.Rooks.LSB()
	moves := gs.GetAllVerticalMovesBitboard(lsb)

	moves_length := 0
	for moves > 0 {
		moves.PopLSB()
		moves_length += 1
	}

	expected_length_moves := 3
	if moves_length != expected_length_moves {
		t.Fatalf(`len(moves) == %v. Expected %v`, moves_length, expected_length_moves)
	}
}

func TestGetAllURDiagonalMovesBB(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_urd_moves := Bitboard(1152925911260069888)

	whiteBishosp := gs.Board.Bishops & gs.Board.PlayerPieces(WHITE)

	moves_bb := gs.GetAllURDiagonalMovesBitboard(whiteBishosp)
	if moves_bb != expected_urd_moves {
		t.Fatalf(`Moves != expected. %v != %v`, moves_bb, expected_urd_moves)
	}
}

func TestGetAllDRDiagonalMovesBB(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_urd_moves := Bitboard(288248105776709632)

	whiteBishosp := gs.Board.Bishops & gs.Board.PlayerPieces(WHITE)

	moves_bb := gs.GetAllDRDiagonalMovesBitboard(whiteBishosp)
	if moves_bb != expected_urd_moves {
		t.Fatalf(`Moves != expected. %v != %v`, moves_bb, expected_urd_moves)
	}
}

func TestGetBishopMoves(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_length := 7
	moves_bb := gs.GetAllBishopMoves()
	if len(moves_bb) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves_bb), expected_length)
	}
}
func TestGetQueenMoves(t *testing.T) {
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	// all_expected := 1441174017036779520
	expected_length := 14
	moves_bb := gs.GetAllQueenMoves()
	if len(moves_bb) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves_bb), expected_length)
	}
}

func TestGetAllCastleMoves(t *testing.T) {
	gs := NewGamestateFEN("r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R1B1K2R w KQ - 1 10")

	expected_length := 1
	moves := gs.GetAllCastleMoves()
	if len(moves) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves), expected_length)

	}
}
func TestGetAllCastleMoves2(t *testing.T) {
	gs := NewGamestateFEN("r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R3K2R w KQ - 1 10")

	expected_length := 2
	moves := gs.GetAllCastleMoves()
	if len(moves) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves), expected_length)

	}
}
func TestGetAllCastleMoves3(t *testing.T) {
	gs := NewGamestateFEN("r3kb1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R3K2R b KQkq - 1 10")

	expected_length := 1
	moves := gs.GetAllCastleMoves()
	if len(moves) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves), expected_length)

	}
}
func TestGetAllCastleMoves4(t *testing.T) {
	gs := NewGamestateFEN("r1b1k2r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R3K2R b KQkq - 1 10")

	expected_length := 1
	moves := gs.GetAllCastleMoves()
	if len(moves) != expected_length {
		t.Fatalf(`Moves != expected. %v != %v`, len(moves), expected_length)

	}
}

func TestSpotUnderAttack(t *testing.T) {
	gs := NewGamestateFEN("rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
	expected := true
	actual := gs.Board.SpotUnderAttack(68719476736, BLACK)

	if expected != actual {
		t.Fatalf(`Spot under attack = %v. Expected %v`, actual, expected)
	}
	result = gs.GetAllMoves()
}

func TestMultipleSpotsUnderAttack(t *testing.T) {
	gs := NewGamestateFEN("rnbqkb1r/p4ppp/2p1pn2/1p1p1Q2/4P3/2N2N2/PPPP1PPP/R1B1KB1R w KQkq b6 0 6")

	black_pieces := gs.Board.BlackPieces

	expected_under_attack := 5
	actual_under_attack := 0
	for black_pieces > 0 {
		lsb := black_pieces.PopLSB()
		if gs.Board.SpotUnderAttack(lsb, BLACK) {
			actual_under_attack += 1
		}
	}
	if expected_under_attack != actual_under_attack {
		t.Fatalf(`Expected %v pieces under attack. Got %v`, expected_under_attack, actual_under_attack)
	}
}
func TestMultipleSpotsUnderAttackTwo(t *testing.T) {
	gs := NewGamestateFEN("r3kb1r/p2n1ppp/b1p1pn2/qB1pNQ2/1P2P3/2N5/P1PP1PPP/R1B1K2R w KQkq - 3 9")

	black_pieces := gs.Board.BlackPieces

	expected_under_attack := 8
	actual_under_attack := 0
	for black_pieces > 0 {
		lsb := black_pieces.PopLSB()
		if gs.Board.SpotUnderAttack(lsb, BLACK) {
			actual_under_attack += 1
		}
	}
	if expected_under_attack != actual_under_attack {
		t.Fatalf(`Expected %v pieces under attack. Got %v`, expected_under_attack, actual_under_attack)
	}

	white_pieces := gs.Board.PlayerPieces(WHITE)
	expected_under_attack = 6
	actual_under_attack = 0
	for white_pieces > 0 {
		lsb := white_pieces.PopLSB()
		if gs.Board.SpotUnderAttack(lsb, WHITE) {
			actual_under_attack += 1
		}
	}
	if expected_under_attack != actual_under_attack {
		t.Fatalf(`Expected %v pieces under attack. Got %v`, expected_under_attack, actual_under_attack)
	}
}

func TestMultipleSpotsUnderAttackTwo_Boards(t *testing.T) {
	gs := NewGamestateFEN("r3kb1r/p2n1ppp/b1p1pn2/qB1pNQ2/1P2P3/2N5/P1PP1PPP/R1B1K2R w KQkq - 3 9")

	white_pieces := gs.Board.PlayerPieces(WHITE)
	expected_under_attack_bb := Bitboard(215050354944)
	actual_under_attack_bb := Bitboard(0)
	for white_pieces > 0 {
		lsb := white_pieces.PopLSB()
		if gs.Board.SpotUnderAttack(lsb, WHITE) {
			actual_under_attack_bb |= lsb
		}
	}
	if expected_under_attack_bb != actual_under_attack_bb {
		t.Fatalf(`Expected %v pieces under attack. Got %v`, expected_under_attack_bb, actual_under_attack_bb)
	}
}

func TestUnderAttackPawn(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/8/5p2/4K3/8/8 w - - 0 1")

	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}

}
func TestNotUnderAttack(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/8/5r2/4K3/8/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}

func TestUnderAttackRook(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/8/8/4K2r/8/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}
func TestUnderAttackRook2(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/4r3/8/8/4K3/8/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}

func TestUnderAttackRook3(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/8/4r3/4K3/8/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}

func TestUnderAttackQueen(t *testing.T) {
	gs := NewGamestateFEN("4k3/q7/8/8/8/4K3/8/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}
func TestUnderAttackQueen2(t *testing.T) {
	gs := NewGamestateFEN("4k3/q7/8/8/8/4K3/8/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}
func TestUnderAttackQueen3(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/8/8/4K3/3q4/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}

func TestNotUnderAttackQueen(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/8/8/4K3/2q5/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}

func TestNotUnderAttackKnight(t *testing.T) {
	gs := NewGamestateFEN("4k3/8/8/5n2/8/4K3/2q5/8 w - - 0 1")
	white_pieces := gs.Board.PlayerPieces(WHITE)
	white_lsb := white_pieces.LSB()

	if !gs.Board.SpotUnderAttack(white_lsb, WHITE) {
		t.Fatalf(`Spot not under attack`)
	}
}

func BenchmarkGetPawnMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN(starting_fen)

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllPawnMoves()
	}
	result = mvs
}
func BenchmarkGetAllKnightMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN(starting_fen)

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllKnightMoves()
	}
	result = mvs
}

func BenchmarkGetRookMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("rnbqkbn1/1pppppp1/8/7r/8/P7/1PPPPPP1/R3KBNR b KQkq - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllRookMoves()
	}
	result = mvs
}
func BenchmarkGetBishopMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllBishopMoves()
	}
	result = mvs
}

func BenchmarkGetQueenMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("b1rR4/3B2n1/1p6/2pk4/p3p1P1/4P2p/5K1P/5Q2 w - - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllQueenMoves()
	}
	result = mvs
}

func BenchmarkGetAllMoves(b *testing.B) {
	var mvs []Move
	gs := NewGamestateFEN("r3qb1k/1b4p1/p2pr2p/3n4/Pnp1N1N1/6RP/1B3PP1/1B1QR1K1 w - - 0 1")

	for i := 0; i < b.N; i++ {
		mvs = gs.GetAllMoves()
	}
	result = mvs
}
