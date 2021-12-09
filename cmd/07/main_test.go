package main

import "testing"

func TestCheapestPos(t *testing.T) {
	cases := []struct {
		pos      []int
		expected [2]int // best position, no. moves
	}{
		{
			[]int{1, 2, 3},
			[2]int{2, 2},
		},
		{
			[]int{6, 6, 1},
			[2]int{6, 5},
		},
		{
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			[2]int{2, 37},
		},
	}
	for _, c := range cases {
		pos, noMoves := cheapestPos(c.pos)

		if pos != c.expected[0] || noMoves != c.expected[1] {
			t.Errorf("got: [%d %d], expected: %v", pos, noMoves, c.expected)
		}
	}
}
