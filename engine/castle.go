package engine

type Castle struct {
	whiteKing  bool
	whiteQueen bool
	blackKing  bool
	blackQueen bool
}

func (cs *Castle) Copy() Castle {
	newCS := Castle{
		whiteKing:  cs.whiteKing,
		blackKing:  cs.blackKing,
		blackQueen: cs.blackQueen,
		whiteQueen: cs.whiteQueen,
	}
	return newCS
}

func (cs *Castle) ToString() string {
	var v []byte
	if cs.whiteKing {
		v = append(v, 'K')
	}
	if cs.whiteQueen {
		v = append(v, 'Q')
	}
	if cs.blackKing {
		v = append(v, 'k')
	}
	if cs.blackQueen {
		v = append(v, 'q')
	}
	if len(v) == 0 {
		return "-"
	}
	return string(v)
}
