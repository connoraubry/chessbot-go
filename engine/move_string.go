package engine

func (e *Engine) GetStringToMoveMap(moves []Move) map[string]Move {
	stm := make(map[string]Move)

	for _, m := range moves {
		stm[e.GetMoveString(m, moves)] = m
	}

	return stm
}

func (e *Engine) GetMoveString(m Move, moves []Move) string {
	//TODO: find suffix
	var suffix string = ""

	switch m.Castle {
	case WHITEOO, BLACKOO:
		return "O-O" + suffix
	case WHITEOOO, BLACKOOO:
		return "O-O-O" + suffix
	}
	stringBytes := specifyWithOtherPieces(m, moves)

	if m.capture {

		if m.pieceName == PAWN {
			stringBytes = append(stringBytes, byte(fileToLetter[m.start&7]))
		}

		stringBytes = append(stringBytes, 'x')
	}

	stringBytes = append(stringBytes, []byte(indexToString(m.end))...)

	if m.promotion != EMPTY {
		stringBytes = append(stringBytes, '=', byte(getLetter(m.promotion)))

	}

	return string(stringBytes) + e.GetMoveSuffix(m)
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
	letter := getLetter(m.pieceName)

	if letter != 0 {
		stringBytes = append(stringBytes, byte(letter))
	}
	needToSpecify := false
	canSpecifyFile := true
	canSpecifyRank := true

	if m.pieceName == PAWN {
		return stringBytes
	}

	for _, otherM := range moves {
		//if we hit the same move, skip
		if otherM.start == m.start {
			continue
		}
		//if the piece isn't moving to the same spot, we don't care
		if otherM.end != m.end {
			continue
		}
		// if it's a different piece type, skip
		if otherM.pieceName != m.pieceName {
			continue
		}
		needToSpecify = true

		source_r, source_f := IndexToRankFile(m.start)
		other_r, other_f := IndexToRankFile(otherM.start)

		if source_f == other_f {
			canSpecifyFile = false
		} else if source_r == other_r {
			canSpecifyRank = false
		}

	}

	s := indexToString(m.start)
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
