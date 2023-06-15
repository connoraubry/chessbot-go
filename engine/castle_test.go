package engine

import "testing"

func TestCastleCopy(t *testing.T) {
	cs := Castle{true, true, false, false}
	newCs := cs.Copy()

	if cs != newCs {
		t.Fatalf(`Castle.Copy() does not return copy of castle`)
	}

	newCs.whiteQueen = false

	if newCs.whiteQueen == cs.whiteQueen {
		t.Fatalf(`Castle.Copy() deos not deep copy`)
	}
}

func TestCastleToString(t *testing.T) {
	// castles := make(map[Castle]string)
	castles := map[Castle]string{
		{true, true, true, true}:     "KQkq",
		{true, true, true, false}:    "KQk",
		{true, true, false, true}:    "KQq",
		{true, true, false, false}:   "KQ",
		{true, false, true, true}:    "Kkq",
		{true, false, true, false}:   "Kk",
		{true, false, false, true}:   "Kq",
		{true, false, false, false}:  "K",
		{false, true, true, true}:    "Qkq",
		{false, true, true, false}:   "Qk",
		{false, true, false, true}:   "Qq",
		{false, true, false, false}:  "Q",
		{false, false, true, true}:   "kq",
		{false, false, true, false}:  "k",
		{false, false, false, true}:  "q",
		{false, false, false, false}: "-",
	}

	for castle, str := range castles {
		if castle.ToString() != str {
			t.Fatalf(`castle.ToString == %v. Expected %v`, castle.ToString(), str)
		}
	}
}
