package engine

import (
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
