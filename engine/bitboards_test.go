package engine

import (
	"math/rand"
	"testing"
)

// var resultBitboard Bitboard

func TestLSB(t *testing.T) {

	bboard := Bitboard(3241)
	lsb := bboard.LSB()
	if lsb != Bitboard(1) {
		t.Fatalf(`Bitboard(%v).LSB() == %v. Expected %v`, bboard, lsb, 1)
	}

	bboard = Bitboard(5949342780000)
	lsb = bboard.LSB()
	if lsb != Bitboard(32) {
		t.Fatalf(`Bitboard(%v).LSB() == %v. Expected %v`, bboard, lsb, 32)
	}
}

func TestPopLSB(t *testing.T) {
	bboard_num := 3241
	expected := 3240
	bboard := Bitboard(bboard_num)

	_ = bboard.PopLSB()

	if bboard != Bitboard(expected) {
		t.Fatalf(`Bitboard(%v).PopLSB() == %v. Expected %v`, bboard_num, bboard, expected)
	}

}
func TestURDiagonalToRank(t *testing.T) {
	ur_diag := MAIN_DIAGONAL
	expected := RANK_1_BB

	output := URDiagonalToRank(ur_diag)
	if output != expected {
		t.Fatalf(`URDiagonal(%v) = %v. Expected %v`, ur_diag, output, expected)
	}

	ur_diag = MAIN_DIAGONAL & ^FILE_A_BB & ^FILE_C_BB
	expected = RANK_1_BB & ^FILE_A_BB & ^FILE_C_BB

	output = URDiagonalToRank(ur_diag)
	if output != expected {
		t.Fatalf(`URDiagonal(%v) = %v. Expected %v`, ur_diag, output, expected)
	}
}

func TestDRDiagonalToRank(t *testing.T) {
	ur_diag := DR_DIAGONAL
	expected := RANK_1_BB

	output := DRDiagonalToRank(ur_diag)
	if output != expected {
		t.Fatalf(`URDiagonal(%v) = %v. Expected %v`, ur_diag, output, expected)
	}

	ur_diag = DR_DIAGONAL & ^FILE_A_BB & ^FILE_C_BB
	expected = RANK_1_BB & ^FILE_A_BB & ^FILE_C_BB

	output = DRDiagonalToRank(ur_diag)
	if output != expected {
		t.Fatalf(`URDiagonal(%v) = %v. Expected %v`, ur_diag, output, expected)
	}
}

func TestRankToURDiagonal(t *testing.T) {
	start := RANK_1_BB
	expected := MAIN_DIAGONAL

	output := RankToURDiagonal(start)
	if output != expected {
		t.Fatalf(`URDiagonal(%v) = %v. Expected %v`, start, output, expected)
	}

	expected = MAIN_DIAGONAL & ^FILE_A_BB & ^FILE_C_BB
	start = RANK_1_BB & ^FILE_A_BB & ^FILE_C_BB

	output = RankToURDiagonal(start)
	if output != expected {
		t.Fatalf(`URDiagonal(%v) = %v. Expected %v`, start, output, expected)
	}
}

func TestURDiagonalAndBack(t *testing.T) {

	for j := 0; j < 20; j++ {
		start := Bitboard(rand.Uint64())

		for idx := 0; idx < 16; idx++ {
			toURD := ConvertToURDiagonal(start, idx)
			toRank := URDiagonalToRank(toURD)
			backToURD := RankToURDiagonal(toRank)
			// backToStart := ReverseConvertToURDiagonal(backToURD, idx)

			if toURD != backToURD {
				t.Fatalf(`URD and reverse do not equal. %v != %v`, toURD, backToURD)
			}
		}
	}
}

func TestDRDiagonalAndBack(t *testing.T) {

	for j := 0; j < 20; j++ {
		start := Bitboard(rand.Uint64())

		for idx := 0; idx < 16; idx++ {
			toURD := ConvertToDRDiagonal(start, idx)
			toRank := DRDiagonalToRank(toURD)
			backToURD := RankToDRDiagonal(toRank)
			// backToStart := ReverseConvertToURDiagonal(backToURD, idx)

			if toURD != backToURD {
				t.Fatalf(`URD and reverse do not equal. %v != %v`, toURD, backToURD)
			}
		}
	}
}

func TestURDiagonalAndBack1(t *testing.T) {
	start := Bitboard(537919492)

	idx := 2

	toURD := ConvertToURDiagonal(start, idx)
	toRank := URDiagonalToRank(toURD)
	backToURD := RankToURDiagonal(toRank)
	backToStart := ReverseConvertToURDiagonal(backToURD, idx)

	if toURD != backToURD {
		t.Fatalf(`URD and reverse do not equal. %v != %v`, toURD, backToURD)
	}
	if start != backToStart {
		t.Fatalf(`start and backToStart not equal. %v = %v`, start, backToStart)
	}
}

func TestDRDiagonalAndBack1(t *testing.T) {
	start := Bitboard(67112992)

	idx := 5

	toURD := ConvertToDRDiagonal(start, idx)
	toRank := DRDiagonalToRank(toURD)
	backToURD := RankToDRDiagonal(toRank)
	backToStart := ReverseConvertToDRDiagonal(backToURD, idx)

	if toURD != backToURD {
		t.Fatalf(`URD and reverse do not equal. %v != %v`, toURD, backToURD)
	}
	if start != backToStart {
		t.Fatalf(`start and backToStart not equal. %v = %v`, start, backToStart)
	}
}

func TestURDiagonalAndBack2(t *testing.T) {
	start := Bitboard(1152921513196781568)

	idx := 24

	toURD := ConvertToURDiagonal(start, idx)
	toRank := URDiagonalToRank(toURD)
	backToURD := RankToURDiagonal(toRank)
	backToStart := ReverseConvertToURDiagonal(backToURD, idx)

	if toURD != backToURD {
		t.Fatalf(`URD and reverse do not equal. %v != %v`, toURD, backToURD)
	}
	if start != backToStart {
		t.Fatalf(`start and backToStart not equal. %v = %v`, start, backToStart)
	}
}

func TestDRDiagonalAndBack2(t *testing.T) {
	start := Bitboard(1161929253617401856)

	idx := 39

	toURD := ConvertToDRDiagonal(start, idx)
	toRank := DRDiagonalToRank(toURD)
	backToURD := RankToDRDiagonal(toRank)
	backToStart := ReverseConvertToDRDiagonal(backToURD, idx)

	if toURD != backToURD {
		t.Fatalf(`DRD and reverse do not equal. %v != %v`, toURD, backToURD)
	}
	if start != backToStart {
		t.Fatalf(`start and backToStart not equal. %v = %v`, start, backToStart)
	}
}

func TestURDiagonalAndBack3(t *testing.T) {
	start := Bitboard(35184372351488)

	idx := 0

	toURD := ConvertToURDiagonal(start, idx)
	toRank := URDiagonalToRank(toURD)
	backToURD := RankToURDiagonal(toRank)
	backToStart := ReverseConvertToURDiagonal(backToURD, idx)

	if toURD != backToURD {
		t.Fatalf(`URD and reverse do not equal. %v != %v`, toURD, backToURD)
	}
	if start != backToStart {
		t.Fatalf(`start and backToStart not equal. %v = %v`, start, backToStart)
	}
}

func TestDRDiagonalAndBack3(t *testing.T) {
	start := Bitboard(72062026446274688)

	idx := 56

	toDRD := ConvertToDRDiagonal(start, idx)
	toRank := DRDiagonalToRank(toDRD)
	backToURD := RankToDRDiagonal(toRank)
	backToStart := ReverseConvertToDRDiagonal(backToURD, idx)

	if toDRD != backToURD {
		t.Fatalf(`DRD and reverse do not equal. %v != %v`, toDRD, backToURD)
	}
	if start != backToStart {
		t.Fatalf(`start and backToStart not equal. %v != %v`, start, backToStart)
	}
}

// func BenchmarkRankToURDiagonal(b *testing.B) {
// 	var output Bitboard
// 	for i := 0; i < b.N; i++ {
// 		output = RankToURDiagonal(RANK_1_BB)
// 	}
// 	resultBitboard = output
// }

func TestReverse(t *testing.T) {
	bb := Bitboard(538445828)
	expected := Bitboard(68685856)

	new := bb.Reverse()
	if new != expected {
		t.Fatalf(`Reversed bitboard != expected. %v != %v`, new, expected)
	}

	bb = Bitboard(16711680)
	expected = Bitboard(16711680)
	new = bb.Reverse()
	if new != expected {
		t.Fatalf(`Reversed bitboard != expected. %v != %v`, new, expected)
	}

	bb = Bitboard(9241386718983028737)
	expected = Bitboard(72620827744338048)
	new = bb.Reverse()
	if new != expected {
		t.Fatalf(`Reversed bitboard != expected. %v != %v`, new, expected)
	}
}
func TestVReverse(t *testing.T) {
	bb := Bitboard(2305913412452483074)
	expected := Bitboard(144126217690284064)
	new := bb.VReverse()
	if new != expected {
		t.Fatalf(`Reversed bitboard != expected. %v != %v`, new, expected)
	}
}
