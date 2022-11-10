package engine

import "testing"

func TestRankAndFile(t *testing.T) {
	for i := 0; i < 64; i++ {
		rank, file := rank_and_file(i)
		if file != i%8 {
			t.Fatalf(`File %v != %v`, file, i%8)
		}
		if rank != (i-file)/8 {
			t.Fatalf(`Rank %v != %v`, rank, (i-file)/8)
		}
	}
}
