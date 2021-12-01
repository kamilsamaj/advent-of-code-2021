package main

import "testing"

func TestCountIncreases(t *testing.T) {
	// these cases should succeed
	cases := []struct {
		input    string
		expected int64
	}{
		{"1\n2\n3", 2},
		{"3\n2\n1", 0},
		{"-1\n-2\n-3", 0},
		{"-10\n-8\n1\n-20", 2},
	}
	for _, c := range cases {
		res, err := countIncreases(c.input)
		if err != nil {
			t.Errorf("got error: %s", err)
		}
		if res != c.expected {
			t.Errorf("input %s, expected: %d, got: %d", c.input, c.expected, res)
		}
	}
}
