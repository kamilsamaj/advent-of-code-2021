package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"math"
	"strings"
)

const bitSize = 12

func calcMostCommonBits(lines []string) (resArr [bitSize]uint8, err error) {
	var resArrSums [bitSize]int
	var i int
	for _, line := range lines {
		for j, c := range line {
			resArrSums[j] += int(c) - 48 // '0' = 48, '1' = 49
		}
		i++
	}

	// if sum is > i/2, set '1'
	for k := 0; k < bitSize; k++ {
		if float64(resArrSums[k]) < float64(i)/2 {
			resArr[k] = 0
		} else {
			resArr[k] = 1
		}
	}
	return
}

func invertBits(bits [bitSize]uint8) [bitSize]uint8 {
	var res [bitSize]uint8
	for i := 0; i < bitSize; i++ {
		if bits[i] == 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}

func convertIntBitsToInt64(bits [bitSize]uint8) float64 {
	res := float64(0)
	for i := 0; i < bitSize; i++ {
		res += float64(bits[i]) * math.Pow(2.0, 12-1-float64(i))
	}
	return res
}

func calcFactor(lines []string, invert bool) ([bitSize]uint8, error) {
	oxgnLines := make([]string, len(lines)) // these lines will be reduced based on the most common bits
	copy(oxgnLines, lines)
	for i := 0; i < bitSize; i++ {
		commBits, err := calcMostCommonBits(oxgnLines)
		if invert {
			commBits = invertBits(commBits)
		}
		if err != nil {
			log.Fatalln(err)
		}
		// purge entries that don't match the bit
		var filteredLines []string
		for _, l := range oxgnLines {
			if l == "" {
				continue
			}
			if (l[i] - 48) == commBits[i] {
				// copy the line to the new array
				filteredLines = append(filteredLines, l)
			}
		}
		oxgnLines = filteredLines // swap the filtered array for the original
		if len(oxgnLines) <= 1 {
			break
		}
	}
	var res [bitSize]uint8
	if len(oxgnLines) == 1 {
		for i := 0; i < bitSize; i++ {
			res[i] = oxgnLines[0][i] - 48
		}
		return res, nil
	} else {
		return res, fmt.Errorf("unexpected result from calcOxygenFactor: %v", oxgnLines)
	}
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/3/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Task 1 - calculate the most and the least common bits
	lines := strings.Split(string(input), "\n")
	bits, err := calcMostCommonBits(lines)
	if err != nil {
		log.Fatalln(err)
	}
	gamma := convertIntBitsToInt64(bits)
	epsilon := convertIntBitsToInt64(invertBits(bits))
	fmt.Printf("Task 1: %d\n", int64(gamma*epsilon))

	// Task 2 - find the oxygen generator and CO2 scrubbing factors
	oxgnLine, err := calcFactor(lines, false)
	if err != nil {
		log.Fatalln(err)
	}

	co2Line, err := calcFactor(lines, true)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Task 2: Oxygen line: %v, CO2 scrubbing line: %v, result: %d",
		oxgnLine,
		co2Line,
		int64(convertIntBitsToInt64(oxgnLine)*convertIntBitsToInt64(co2Line)))
}
