package main

import (
	"strings"
	"testing"
)

const input = `2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestTask1(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := task1(lines); res != 15 {
		t.Errorf("expected: 15, got: %d", res)
	}
}

func TestTask2(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := task2(lines); res != 1134 {
		t.Errorf("expected: 1134, got: %d", res)
	}
}
