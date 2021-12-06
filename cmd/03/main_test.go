package main

import (
	"strings"
	"testing"
)

func TestCalcMostCommonBits(t *testing.T) {
	cases := []struct {
		input    string
		expected [bitSize]uint8
	}{
		{
			`101010000100
100001010100
111100000101
010000000010
001101100010
100110110101
110100101101`,
			[bitSize]uint8{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		},
	}

	for _, c := range cases {
		resArr, err := calcMostCommonBits(strings.Split(c.input, "\n"))
		if err != nil {
			t.Errorf("%s", err)
		}
		if resArr != c.expected {
			t.Errorf("expected: %v, got %v", c.expected, resArr)
		}
	}

}

func TestConvertIntBitsToInt64(t *testing.T) {
	if convertIntBitsToInt64([bitSize]uint8{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}) != 2308.0 {
		t.Errorf("{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0} should result into 2308")
	}
}

func TestInvertBits(t *testing.T) {
	if invertBits([bitSize]uint8{1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}) != [bitSize]uint8{0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1} {
		t.Errorf("error in bits inversion")
	}
}
