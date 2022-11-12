package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type Bitboard uint64
type Bitrow uint8

/*
PrintBitboard(b Bitboard)

Prints a bitboard to the terminal.
0's are represented as '.'

Mainly used for debugging
*/
func PrintBitboard(b Bitboard) {
	bits := strconv.FormatUint(uint64(b), 2)
	slice := strings.Split(bits, "")
	if len(slice) < 64 {
		prefix := make([]string, 64-len(slice))
		for i := range prefix {
			prefix[i] = "0"
		}
		slice = append(prefix, slice...)
	}

	for i := 0; i < 8; i++ {
		s := slice[i*8 : (i*8)+8]
		printSlice := make([]string, 8)

		for idx := 7; idx >= 0; idx-- {

			val := s[idx]
			if val == "0" {
				val = "."
			}
			printSlice[7-idx] = val
		}
		fmt.Println(strings.Join(printSlice, " "))
	}
	fmt.Println()
}

/*
PrintBitrow(b Bitrow)

Prints a bitrow to the terminal.
0's are represented as '.'

Mainly used for debugging.
*/
func PrintBitrow(b Bitrow) {
	bits := strconv.FormatUint(uint64(b), 2)
	slice := strings.Split(bits, "")
	if len(slice) < 8 {
		prefix := make([]string, 8-len(slice))
		for i := range prefix {
			prefix[i] = "0"
		}
		slice = append(prefix, slice...)
	}

	printSlice := make([]string, 8)

	for idx := 7; idx >= 0; idx-- {

		val := slice[idx]
		if val == "0" {
			val = "."
		}
		printSlice[7-idx] = val

	}
	fmt.Println(strings.Join(printSlice, " "))

}

/*
Bitboard.LSB() -> Bitboard

Gets a bitboard of the least significant bit of B
*/
func (b Bitboard) LSB() Bitboard {
	return (b & -b)
}

func (b *Bitboard) PopLSB() Bitboard {
	lsb := b.LSB()
	*b = *b - lsb
	return lsb
}

func (b *Bitboard) ShiftPawns(player Player) {
	switch player {
	case WHITE:
		b.ShiftNorth()
	case BLACK:
		b.ShiftSouth()
	}
}

func (b *Bitboard) ShiftNorth() {
	*b = *b << 8
}
func (b *Bitboard) ShiftSouth() {
	*b = *b >> 8
}
func (b *Bitboard) ShiftWest() {
	*b = (*b & (^FILE_A_BB)) >> 1
}
func (b *Bitboard) ShiftEast() {
	*b = (*b & (^FILE_H_BB)) << 1
}

func (b Bitboard) Index() int {
	lsb := b.LSB()
	return BoardConstants.BitToIndex[lsb]
}

func (b Bitboard) Rank() int {
	return b.Index() >> 3
}
func (b Bitboard) File() int {
	return b.Index() & 7
}

func Overlap(b1, b2 Bitboard) bool {
	overlap := b1 & b2
	return overlap > 0
}

func AFileToRank(b Bitboard) Bitboard {
	b = b & FILE_A_BB
	for i := 0; i < 7; i++ {
		b = b | (b >> 7)
	}
	return b & RANK_1_BB
}
func RankToAFile(b Bitboard) Bitboard {
	b = b & RANK_1_BB
	for i := 0; i < 7; i++ {
		b = b | (b << 7)
	}
	return b & FILE_A_BB
}
func URDiagonalToRank(b Bitboard) Bitboard {
	new_bb := (b & FILE_A_BB) |
		(b&FILE_B_BB)>>RANK_SHIFT_1 |
		(b&FILE_C_BB)>>RANK_SHIFT_2 |
		(b&FILE_D_BB)>>RANK_SHIFT_3 |
		(b&FILE_E_BB)>>RANK_SHIFT_4 |
		(b&FILE_F_BB)>>RANK_SHIFT_5 |
		(b&FILE_G_BB)>>RANK_SHIFT_6 |
		(b&FILE_H_BB)>>RANK_SHIFT_7

	return new_bb
}

func RankToURDiagonal(b Bitboard) Bitboard {
	new_bb := (b & FILE_A_BB) |
		(b&FILE_B_BB)<<RANK_SHIFT_1 |
		(b&FILE_C_BB)<<RANK_SHIFT_2 |
		(b&FILE_D_BB)<<RANK_SHIFT_3 |
		(b&FILE_E_BB)<<RANK_SHIFT_4 |
		(b&FILE_F_BB)<<RANK_SHIFT_5 |
		(b&FILE_G_BB)<<RANK_SHIFT_6 |
		(b&FILE_H_BB)<<RANK_SHIFT_7

	return new_bb
}

func DRDiagonalToRank(b Bitboard) Bitboard {
	new_bb := (b & FILE_H_BB) |
		(b&FILE_G_BB)>>RANK_SHIFT_1 |
		(b&FILE_F_BB)>>RANK_SHIFT_2 |
		(b&FILE_E_BB)>>RANK_SHIFT_3 |
		(b&FILE_D_BB)>>RANK_SHIFT_4 |
		(b&FILE_C_BB)>>RANK_SHIFT_5 |
		(b&FILE_B_BB)>>RANK_SHIFT_6 |
		(b&FILE_A_BB)>>RANK_SHIFT_7

	return new_bb
}

func RankToDRDiagonal(b Bitboard) Bitboard {
	new_bb := (b & FILE_H_BB) |
		(b&FILE_G_BB)<<RANK_SHIFT_1 |
		(b&FILE_F_BB)<<RANK_SHIFT_2 |
		(b&FILE_E_BB)<<RANK_SHIFT_3 |
		(b&FILE_D_BB)<<RANK_SHIFT_4 |
		(b&FILE_C_BB)<<RANK_SHIFT_5 |
		(b&FILE_B_BB)<<RANK_SHIFT_6 |
		(b&FILE_A_BB)<<RANK_SHIFT_7

	return new_bb
}

func ConvertToURDiagonal(b Bitboard, idx int) Bitboard {
	rank, file := IndexToRankFile(idx)
	diag_idx := (rank - file) & 15
	if diag_idx == 0 {
		b = b & MAIN_DIAGONAL
	} else if diag_idx < 8 {
		b = b >> (RANK_SHIFT_1 * diag_idx)
		b = b & MAIN_DIAGONAL
	} else if diag_idx > 8 {
		b = b << (RANK_SHIFT_1 * (16 - diag_idx))
		b = b & MAIN_DIAGONAL
	}

	return b
}
func ReverseConvertToURDiagonal(b Bitboard, idx int) Bitboard {
	rank, file := IndexToRankFile(idx)
	diag_idx := (rank - file) & 15
	if diag_idx == 0 {
		b = b & MAIN_DIAGONAL
	} else if diag_idx < 8 {
		b = b & MaskRows[8-diag_idx]
		b = b << (RANK_SHIFT_1 * diag_idx)
	} else if diag_idx > 8 {
		b = b >> (RANK_SHIFT_1 * (16 - diag_idx))
	}
	return b
}

func ConvertToDRDiagonal(b Bitboard, idx int) Bitboard {
	rank, file := IndexToRankFile(idx)
	diag_idx := (rank + file) ^ 7
	if diag_idx == 0 {
		b = b & DR_DIAGONAL
	} else if diag_idx < 8 {
		b = b << (RANK_SHIFT_1 * diag_idx)
		b = b & DR_DIAGONAL
	} else if diag_idx > 8 {
		b = b >> (RANK_SHIFT_1 * (16 - diag_idx))
		b = b & DR_DIAGONAL
	}
	return b
}

func ReverseConvertToDRDiagonal(b Bitboard, idx int) Bitboard {
	rank, file := IndexToRankFile(idx)
	diag_idx := (rank + file) ^ 7
	if diag_idx == 0 {
		b = b & DR_DIAGONAL
	} else if diag_idx < 8 {
		b = b >> (RANK_SHIFT_1 * diag_idx)
	} else if diag_idx > 8 {
		b = b & MaskRows[diag_idx-8]
		b = b << (RANK_SHIFT_1 * (16 - diag_idx))
	}
	return b
}
