package engine

import "testing"

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
