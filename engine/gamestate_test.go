package engine

import "testing"

func TestCastleStruct(t *testing.T) {
	cs := Castle{}
	cs.whiteKing = true
	cs.blackKing = true
	cs.whiteQueen = true
	cs.blackQueen = true
}

func TestNewGamestate(t *testing.T) {

}
