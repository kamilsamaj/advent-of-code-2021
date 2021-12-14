package main

import (
	"strings"
	"testing"
)

const input = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func TestTask1(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := task1(lines, 100); res != 1656 {
		t.Errorf("expected: 1656 flashes, got: %d", res)
	}
}
