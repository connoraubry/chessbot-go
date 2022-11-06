package board

import "testing"

func TestBoardInit(t *testing.T) {
	b := Board{}
	if b.whiteKing != 0 {
		t.Fatalf(`b.whiteKing == %v. Expected 0`, b.whiteKing)

	}
	if b.blackKing != 0 {
		t.Fatalf(`b.blackKing == %v. Expected 0`, b.whiteKing)

	}
	emptyPiece := Piece{name: EMPTY}
	for idx, piece := range b.board {
		if piece != emptyPiece {
			t.Fatalf(`b.board[%v] == %v. Expected %v`, idx, piece, EMPTY)
		}
	}
}

func TestNewBoard(t *testing.T) {

	var idxToPiece [64]Piece
	for i := 0; i < 64; i++ {
		idxToPiece[i] = Piece{name: EMPTY}
	}
	for i := 8; i < 16; i++ {
		idxToPiece[i] = Piece{name: PAWN, player: WHITE}
	}
	for i := 48; i < 56; i++ {
		idxToPiece[i] = Piece{name: PAWN, player: BLACK}
	}
	idxToPiece[0] = Piece{name: ROOK, player: WHITE}
	idxToPiece[1] = Piece{name: KNIGHT, player: WHITE}
	idxToPiece[2] = Piece{name: BISHOP, player: WHITE}
	idxToPiece[3] = Piece{name: QUEEN, player: WHITE}
	idxToPiece[4] = Piece{name: KING, player: WHITE}
	idxToPiece[5] = Piece{name: BISHOP, player: WHITE}
	idxToPiece[6] = Piece{name: KNIGHT, player: WHITE}
	idxToPiece[7] = Piece{name: ROOK, player: WHITE}

	idxToPiece[56] = Piece{name: ROOK, player: BLACK}
	idxToPiece[57] = Piece{name: KNIGHT, player: BLACK}
	idxToPiece[58] = Piece{name: BISHOP, player: BLACK}
	idxToPiece[59] = Piece{name: QUEEN, player: BLACK}
	idxToPiece[60] = Piece{name: KING, player: BLACK}
	idxToPiece[61] = Piece{name: BISHOP, player: BLACK}
	idxToPiece[62] = Piece{name: KNIGHT, player: BLACK}
	idxToPiece[63] = Piece{name: ROOK, player: BLACK}

	b := NewBoard(Starting_board_fen_string)
	if b.whiteKing != 4 {
		t.Fatalf(`b.whiteKing == %v. Expected 4`, b.whiteKing)

	}
	if b.blackKing != 60 {
		t.Fatalf(`b.blackKing == %v. Expected 60`, b.whiteKing)

	}
	for idx, piece := range b.board {
		if piece != idxToPiece[idx] {
			t.Fatalf(`b.board[%v] == %v. Expected %v`, idx, piece, idxToPiece[idx])
		}
	}
}

func TestMiddleFEN(t *testing.T) {
	var idxToPiece [64]Piece
	for i := 0; i < 64; i++ {
		idxToPiece[i] = Piece{name: EMPTY}
	}

	idxToPiece[0] = Piece{name: ROOK, player: WHITE}
	idxToPiece[2] = Piece{name: BISHOP, player: WHITE}
	idxToPiece[4] = Piece{name: KING, player: WHITE}
	idxToPiece[7] = Piece{name: ROOK, player: WHITE}
	idxToPiece[8] = Piece{name: PAWN, player: WHITE}
	idxToPiece[9] = Piece{name: PAWN, player: WHITE}
	idxToPiece[13] = Piece{name: PAWN, player: WHITE}
	idxToPiece[14] = Piece{name: PAWN, player: WHITE}
	idxToPiece[15] = Piece{name: PAWN, player: WHITE}
	idxToPiece[20] = Piece{name: PAWN, player: WHITE}
	idxToPiece[21] = Piece{name: KNIGHT, player: WHITE}
	idxToPiece[26] = Piece{name: BISHOP, player: WHITE}
	idxToPiece[32] = Piece{name: PAWN, player: BLACK}
	idxToPiece[34] = Piece{name: PAWN, player: WHITE}
	idxToPiece[36] = Piece{name: PAWN, player: BLACK}
	idxToPiece[45] = Piece{name: KNIGHT, player: BLACK}
	idxToPiece[49] = Piece{name: PAWN, player: BLACK}
	idxToPiece[53] = Piece{name: PAWN, player: BLACK}
	idxToPiece[54] = Piece{name: PAWN, player: BLACK}
	idxToPiece[55] = Piece{name: PAWN, player: BLACK}
	idxToPiece[56] = Piece{name: ROOK, player: BLACK}
	idxToPiece[58] = Piece{name: BISHOP, player: BLACK}
	idxToPiece[59] = Piece{name: KING, player: BLACK}
	idxToPiece[61] = Piece{name: BISHOP, player: BLACK}
	idxToPiece[63] = Piece{name: ROOK, player: BLACK}

	b := NewBoard("r1bk1b1r/1p3ppp/5n2/p1P1p3/2B5/4PN2/PP3PPP/R1B1K2R")
	if b.whiteKing != 4 {
		t.Fatalf(`b.whiteKing == %v. Expected 4`, b.whiteKing)

	}
	if b.blackKing != 59 {
		t.Fatalf(`b.blackKing == %v. Expected 60`, b.whiteKing)

	}
	for idx, piece := range b.board {
		if piece != idxToPiece[idx] {
			t.Fatalf(`b.board[%v] == %v. Expected %v`, idx, piece, idxToPiece[idx])
		}
	}

}
