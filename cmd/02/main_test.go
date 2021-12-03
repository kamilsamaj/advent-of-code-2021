package main

import "testing"

func TestCalculateVector(t *testing.T) {
	cases := []struct {
		cmd      string
		expected Vector
	}{
		{"forward 7", Vector{7, 0}},
		{"down 4", Vector{0, 4}},
		{"up 9", Vector{0, -9}},
	}

	for _, c := range cases {
		v, err := calculateVector(c.cmd)
		if err != nil {
			t.Errorf("%v", err)
		}
		if v != c.expected {
			t.Errorf("cmd: '%s', expected: %v, got: %v", c.cmd, c.expected, v)
		}
	}
}

func TestCalculatePosition(t *testing.T) {
	cases := []struct {
		bytes    []byte
		expected int64
	}{
		{[]byte(`forward 3
down 3
down 4
forward 1
forward 1
forward 7
`),
			84},
		{[]byte(`forward 2
down 10
down 4
up 2
up 8
forward 6`),
			32,
		},
	}

	for _, c := range cases {
		p, err := calculatePosition(&c.bytes)
		if err != nil {
			t.Errorf("got error: %s", err)
		}
		if p != c.expected {
			t.Errorf("%v", err)
		}
	}
}

func TestCalculatePositionWithAim(t *testing.T) {
	cases := []struct {
		bytes    []byte
		expected int64
	}{
		{[]byte(`forward 5
down 5
forward 8
up 3
down 8
forward 2
`),
			900}, // f = 23  d = 77+66+126=269
		{[]byte(`forward 3
down 1
down 3
down 6
down 1
forward 7
forward 6
down 5
down 2
forward 7`),
			6187,
		},
	}
	for _, c := range cases {
		p, err := calculatePositionWithAim(&c.bytes)
		if err != nil {
			t.Errorf("got error: %s", err)
		}
		if p != c.expected {
			t.Errorf("cmd: '%s', expected: %v, got: %v", string(c.bytes), c.expected, p)
		}
	}
}
