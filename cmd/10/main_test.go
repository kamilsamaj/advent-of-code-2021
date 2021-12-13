package main

import (
	"strings"
	"testing"
)

const input = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func TestTask1(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := task1(lines); res != 26397 {
		t.Errorf("expected: 26397, got: %d", res)
	}
}

func TestTask2(t *testing.T) {
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	if res := task2(lines); res != 288957 {
		t.Errorf("expected: 288957, got: %d", res)
	}
}
