package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"math/rand"
)

type Book struct {
	Cache map[string]BookEntry `json:"cache"`
}

type BookEntry struct {
	AcceptableMoves []BookMove `json:"move_list"`
}

type BookMove struct {
	Move    string `json:"move"`
	Percent int    `json:"percent"`
}

func NewBook(filepath string) *Book {

	var b *Book
	var err error
	if filepath == "" {
		b = new(Book)
		b.Cache = make(map[string]BookEntry)
	} else {
		b, err = LoadBookFromFile(filepath)
		if err != nil {
			panic(err)
		}
	}

	return b
}

func LoadBookFromFile(filepath string) (*Book, error) {

	var b Book
	b.Cache = make(map[string]BookEntry)

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		b.SaveToFile(filepath)
	} else if err != nil {
		panic(err)
	}

	jsonBytes, err := os.ReadFile(filepath)

	if err != nil {
		return &Book{}, fmt.Errorf("Error reading file: %v", err)
	}

	err = json.Unmarshal(jsonBytes, &b)
	if err != nil {
		return &Book{}, fmt.Errorf("Error parsing json file: %v", err)
	}

	return &b, nil
}

func (b *Book) SaveToFile(filepath string) error {
	js, err := json.Marshal(b)
	if err != nil {
		return fmt.Errorf("Error marshalling book: %v", err)
	}

	err = os.WriteFile(filepath, js, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}
	return nil
}

func (b *Book) GetNextMove(fen string) (string, error) {

	entry, ok := b.Cache[fen]

	if !ok {
		return "", fmt.Errorf("No entry found for fen %v", fen)
	}

	/*
		Weighted random example

		number = 65
		book move percents = [30, 20, 50]

		iter 1: number !< percent
		number = 35

		iter 2: number !< percet
		number = 15

		iter 3: number < percent
		return last


		if number [0-30], first entry
		if number (30-50], second entry
		else: third entry
	*/

	number := rand.Float64() * 100
	for _, book_move := range entry.AcceptableMoves {

		p_val := float64(book_move.Percent)

		if number < p_val {
			return book_move.Move, nil
		} else {
			number -= p_val
		}
	}

	return "", fmt.Errorf("No valid move found")
}
