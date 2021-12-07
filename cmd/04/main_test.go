package main

import "testing"

func TestLineToInts(t *testing.T) {
	cases := []struct {
		line     string
		expected [5]boardItem
	}{
		{
			"26 68  3 95 59",
			[5]boardItem{
				{26, false},
				{68, false},
				{3, false},
				{95, false},
				{59, false},
			},
		},
	}
	for _, c := range cases {
		nums, err := lineToInts(c.line)

		if err != nil {
			t.Error(err)
		}
		if nums != c.expected {
			t.Errorf("expected: %v, got: %v", c.expected, nums)
		}
	}
}
