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

//https://stackoverflow.com/questions/746171/efficient-algorithm-for-bit-reversal-from-msb-lsb-to-lsb-msb-in-c
func (b Bitboard) Reverse() Bitboard {
	return Bitboard(
		bitReverseLookupTable[(b>>56)&RANK_1_BB]<<56 |
			bitReverseLookupTable[(b>>48)&RANK_1_BB]<<48 |
			bitReverseLookupTable[(b>>40)&RANK_1_BB]<<40 |
			bitReverseLookupTable[(b>>32)&RANK_1_BB]<<32 |
			bitReverseLookupTable[(b>>24)&RANK_1_BB]<<24 |
			bitReverseLookupTable[(b>>16)&RANK_1_BB]<<16 |
			bitReverseLookupTable[(b>>8)&RANK_1_BB]<<8 |
			bitReverseLookupTable[b&RANK_1_BB])
}

func (b Bitboard) VReverse() Bitboard {
	return Bitboard(
		(b&RANK_8_BB)>>56 |
			(b&RANK_7_BB)>>40 |
			(b&RANK_6_BB)>>24 |
			(b&RANK_5_BB)>>8 |
			(b&RANK_4_BB)<<8 |
			(b&RANK_3_BB)<<24 |
			(b&RANK_2_BB)<<40 |
			(b&RANK_1_BB)<<56)
}

var bitReverseLookupTable = []uint64{
	0x00, 0x80, 0x40, 0xC0, 0x20, 0xA0, 0x60, 0xE0, 0x10, 0x90, 0x50, 0xD0, 0x30, 0xB0, 0x70, 0xF0,
	0x08, 0x88, 0x48, 0xC8, 0x28, 0xA8, 0x68, 0xE8, 0x18, 0x98, 0x58, 0xD8, 0x38, 0xB8, 0x78, 0xF8,
	0x04, 0x84, 0x44, 0xC4, 0x24, 0xA4, 0x64, 0xE4, 0x14, 0x94, 0x54, 0xD4, 0x34, 0xB4, 0x74, 0xF4,
	0x0C, 0x8C, 0x4C, 0xCC, 0x2C, 0xAC, 0x6C, 0xEC, 0x1C, 0x9C, 0x5C, 0xDC, 0x3C, 0xBC, 0x7C, 0xFC,
	0x02, 0x82, 0x42, 0xC2, 0x22, 0xA2, 0x62, 0xE2, 0x12, 0x92, 0x52, 0xD2, 0x32, 0xB2, 0x72, 0xF2,
	0x0A, 0x8A, 0x4A, 0xCA, 0x2A, 0xAA, 0x6A, 0xEA, 0x1A, 0x9A, 0x5A, 0xDA, 0x3A, 0xBA, 0x7A, 0xFA,
	0x06, 0x86, 0x46, 0xC6, 0x26, 0xA6, 0x66, 0xE6, 0x16, 0x96, 0x56, 0xD6, 0x36, 0xB6, 0x76, 0xF6,
	0x0E, 0x8E, 0x4E, 0xCE, 0x2E, 0xAE, 0x6E, 0xEE, 0x1E, 0x9E, 0x5E, 0xDE, 0x3E, 0xBE, 0x7E, 0xFE,
	0x01, 0x81, 0x41, 0xC1, 0x21, 0xA1, 0x61, 0xE1, 0x11, 0x91, 0x51, 0xD1, 0x31, 0xB1, 0x71, 0xF1,
	0x09, 0x89, 0x49, 0xC9, 0x29, 0xA9, 0x69, 0xE9, 0x19, 0x99, 0x59, 0xD9, 0x39, 0xB9, 0x79, 0xF9,
	0x05, 0x85, 0x45, 0xC5, 0x25, 0xA5, 0x65, 0xE5, 0x15, 0x95, 0x55, 0xD5, 0x35, 0xB5, 0x75, 0xF5,
	0x0D, 0x8D, 0x4D, 0xCD, 0x2D, 0xAD, 0x6D, 0xED, 0x1D, 0x9D, 0x5D, 0xDD, 0x3D, 0xBD, 0x7D, 0xFD,
	0x03, 0x83, 0x43, 0xC3, 0x23, 0xA3, 0x63, 0xE3, 0x13, 0x93, 0x53, 0xD3, 0x33, 0xB3, 0x73, 0xF3,
	0x0B, 0x8B, 0x4B, 0xCB, 0x2B, 0xAB, 0x6B, 0xEB, 0x1B, 0x9B, 0x5B, 0xDB, 0x3B, 0xBB, 0x7B, 0xFB,
	0x07, 0x87, 0x47, 0xC7, 0x27, 0xA7, 0x67, 0xE7, 0x17, 0x97, 0x57, 0xD7, 0x37, 0xB7, 0x77, 0xF7,
	0x0F, 0x8F, 0x4F, 0xCF, 0x2F, 0xAF, 0x6F, 0xEF, 0x1F, 0x9F, 0x5F, 0xDF, 0x3F, 0xBF, 0x7F, 0xFF,
}
