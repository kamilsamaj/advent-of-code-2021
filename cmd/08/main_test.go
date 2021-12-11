package main

import (
	"reflect"
	"strings"
	"testing"
)

const testInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`

func TestFindNums(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input:    testInput,
			expected: 26,
		},
	}

	for _, c := range cases {
		lines := strings.Split(strings.Trim(string(c.input), "\n"), "\n")
		m := findSimpleNums(lines)
		totalNums := countTotals(m)
		if totalNums != c.expected {
			t.Errorf("Count Total Number of 1, 4, 7, 8: expected: %d, got: %d",
				c.expected, totalNums)
		}
	}
}

func TestDiffSignals(t *testing.T) {
	cases := []struct {
		a, b     string
		expected [2][]rune
	}{
		{
			a:        "acd",
			b:        "ac",
			expected: [2][]rune{{'d'}, nil},
		},
	}
	for _, c := range cases {
		leftInA, leftInB := diffSignals(c.a, c.b)
		if !reflect.DeepEqual(c.expected[0], leftInA) || !reflect.DeepEqual(c.expected[1], leftInB) {
			t.Errorf("a: %s, b: %s, got: [%v %v], expected: %v",
				c.a, c.b, leftInA, leftInB, c.expected)
		}
	}

}

func TestFindAllNums(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input:    testInput,
			expected: 61229,
		},
	}
	for _, c := range cases {
		lines := strings.Split(strings.Trim(string(c.input), "\n"), "\n")

		if res := findAllNums(lines); res != c.expected {
			t.Errorf("expected: %d, got: %d", c.expected, res)
		}
	}
}
