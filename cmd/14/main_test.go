package main

import (
	"strings"
	"testing"
)

const input = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func TestTask1(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := expandPolymer(lines, 10); res != 1588 {
		t.Errorf("expected: 1588 finalized paths, got: %d", res)
	}
}

func TestTask2(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := expandPolymer(lines, 40); res != 2188189693529 {
		t.Errorf("expected: 2188189693529 finalized paths, got: %d", res)
	}
}
