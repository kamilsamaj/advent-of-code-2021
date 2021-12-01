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

func TestSlidingWindowIncreases(t *testing.T) {
	// these cases should succeed
	cases := []struct {
		input    string
		expected int64
	}{
		{"1\n2\n3\n4\n5\n6", 3},
		{"3\n2\n1\n0\n-1\n-2", 0},
		{"-1\n-2\n-3\n-2\n-1\n0", 2},
		{"10\n8\n1\n20", 1},
	}
	for _, c := range cases {
		res, err := slidingWindowIncreases(c.input)
		if err != nil {
			t.Errorf("got error: %s", err)
		}
		if res != c.expected {
			t.Errorf("input %s, expected: %d, got: %d", c.input, c.expected, res)
		}
	}
}
