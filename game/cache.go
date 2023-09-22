package game

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/connoraubry/chessbot-go/engine"
)

type Cache struct {
	Cache             map[string]PositionEval
	HalfmoveThreshold int
	Verbose           bool
}

type PositionEval struct {
	Score         int         `json:"score"`
	TimesAccessed int         `json:"times_accessed"`
	LastHalfmove  int         `json:"last_halfmove"`
	BestMove      engine.Move `json:"best_move"`
	DepthAnalyzed int         `json:"depth_analyzed"`
}

func NewCache(threshold int, verbose bool) *Cache {
	c := new(Cache)
	c.Cache = make(map[string]PositionEval)
	c.HalfmoveThreshold = threshold
	c.Verbose = verbose
	return c
}

func LoadCacheFromJSONFile(filepath string) *map[string]PositionEval {
	c := make(map[string]PositionEval)

	jsonBytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(fmt.Errorf("Error reading file: %v", err))
	}

	err = json.Unmarshal(jsonBytes, &c)
	if err != nil {
		panic(fmt.Errorf("Error parsing json file: %v", err))
	}

	return &c

}

/*
Returns amount of entries flushed
*/
func (c *Cache) Flush(halfmove int) int {

	flushCount := 0
	for fen, values := range c.Cache {
		if values.LastHalfmove < halfmove-c.HalfmoveThreshold {
			delete(c.Cache, fen)
			flushCount += 1
		}
	}
	return flushCount
}

func (c *Cache) Len() int {
	return len(c.Cache)
}

func (c *Cache) Lookup(fen string) (PositionEval, bool) {
	pos, ok := c.Cache[fen]
	if ok {
		pos.TimesAccessed += 1
		c.Cache[fen] = pos
	}
	return pos, ok
}

func (c *Cache) GetScore(fen string) (int, bool) {
	pos, ok := c.Lookup(fen)
	return pos.Score, ok
}

func (c *Cache) GetBestMove(fen string) (engine.Move, bool) {
	pos, ok := c.Lookup(fen)
	return pos.BestMove, ok
}

func (c *Cache) Update(fen string, newPosition PositionEval) {
	existingPosition, ok := c.Lookup(fen)

	if ok {
		if newPosition.DepthAnalyzed > existingPosition.DepthAnalyzed {
			c.Cache[fen] = newPosition
		}
	} else {
		c.Cache[fen] = newPosition
	}
}
