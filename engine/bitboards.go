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
