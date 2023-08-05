package game

type Cache struct {
	cache             map[string]PositionEval
	halfmoveThreshold int
	verbose           bool
}

type PositionEval struct {
	score         int
	timesAccessed int
	lastHalfmove  int
	bestMove      string
	depthAnalyzed int
}

func NewCache(threshold int, verbose bool) *Cache {
	c := new(Cache)
	c.cache = make(map[string]PositionEval)
	c.halfmoveThreshold = threshold
	c.verbose = verbose
	return c
}

/*
Returns amount of entries flushed
*/
func (c *Cache) Flush(halfmoveThreshold int) int {

	flushCount := 0
	for fen, values := range c.cache {
		if values.lastHalfmove < halfmoveThreshold {
			delete(c.cache, fen)
			flushCount += 1
		}
	}
	return flushCount
}

func (c *Cache) Lookup(fen string) (PositionEval, bool) {
	pos, ok := c.cache[fen]
	return pos, ok
}

func (c *Cache) GetScore(fen string) (int, bool) {
	pos, ok := c.Lookup(fen)
	return pos.score, ok
}

func (c *Cache) Update(fen string, newPosition PositionEval) {
	existingPosition, ok := c.Lookup(fen)

	if ok {
		if newPosition.depthAnalyzed > existingPosition.depthAnalyzed {
			c.cache[fen] = newPosition
		}
	} else {
		c.cache[fen] = newPosition
	}
}
