package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"math"
	"strings"
)

const bitSize = 12

func calcMostCommonBits(input *string) (resArr [bitSize]uint8, err error) {
	var resArrSums [bitSize]int
	var i int
	for _, line := range strings.Split(*input, "\n") {
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

func main() {
	inputUrl := "https://adventofcode.com/2021/day/3/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Task 1 - calculate the most and the least common bits
	inputStr := string(input)
	bits, err := calcMostCommonBits(&inputStr)
	if err != nil {
		log.Fatalln(err)
	}
	gamma := convertIntBitsToInt64(bits)
	epsilon := convertIntBitsToInt64(invertBits(bits))
	fmt.Printf("Task 1: %d", int64(gamma*epsilon))
}
