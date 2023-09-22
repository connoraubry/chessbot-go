package engine

func (e *Engine) GetStringToMoveMap(moves []Move) map[string]Move {
	stm := make(map[string]Move)

	for _, m := range moves {
		stm[e.GetMoveString(m, moves)] = m
	}

	return stm
}

func (e *Engine) GetMoveString(m Move, moves []Move) string {
	var suffix string = e.GetMoveSuffix(m)

	switch m.Castle {
	case WHITEOO, BLACKOO:
		return "O-O" + suffix
	case WHITEOOO, BLACKOOO:
		return "O-O-O" + suffix
	}
	stringBytes := specifyWithOtherPieces(m, moves)

	if m.Capture {

		if m.PieceName == PAWN {
			stringBytes = append(stringBytes, byte(fileToLetter[m.Start&7]))
		}

		stringBytes = append(stringBytes, 'x')
	}

	stringBytes = append(stringBytes, []byte(indexToString(m.End))...)

	if m.Promotion != EMPTY {
		stringBytes = append(stringBytes, '=', byte(getLetter(m.Promotion)))

	}

	return string(stringBytes) + suffix
}

func (e *Engine) GetMoveSuffix(m Move) string {

	suffix := ""

	res := e.TakeMove(m)
	if res {
		if e.PlayerInCheckmate() {
			suffix = "#"
		} else if e.PlayerInCheck() {
			suffix = "+"
		}
		e.UndoMove()
	}
	return suffix
}

func specifyWithOtherPieces(m Move, moves []Move) []byte {
	var stringBytes []byte
	letter := getLetter(m.PieceName)

	if letter != 0 {
		stringBytes = append(stringBytes, byte(letter))
	}
	needToSpecify := false
	canSpecifyFile := true
	canSpecifyRank := true

	if m.PieceName == PAWN {
		return stringBytes
	}

	for _, otherM := range moves {
		//if we hit the same move, skip
		if otherM.Start == m.Start {
			continue
		}
		//if the piece isn't moving to the same spot, we don't care
		if otherM.End != m.End {
			continue
		}
		// if it's a different piece type, skip
		if otherM.PieceName != m.PieceName {
			continue
		}
		needToSpecify = true

		source_r, source_f := IndexToRankFile(m.Start)
		other_r, other_f := IndexToRankFile(otherM.Start)

		if source_f == other_f {
			canSpecifyFile = false
		} else if source_r == other_r {
			canSpecifyRank = false
		}

	}

	s := indexToString(m.Start)
	if needToSpecify {
		if canSpecifyFile {
			stringBytes = append(stringBytes, s[0])
		} else if canSpecifyRank {
			stringBytes = append(stringBytes, s[1])
		} else {
			stringBytes = append(stringBytes, s[0], s[1])
		}
	}

	return stringBytes
}
