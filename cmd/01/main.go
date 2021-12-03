package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strconv"
	"strings"
)

func countIncreases(input string) (increases int64, err error) {
	var prevNumber, currNumber int64
	for i, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}
		if i == 0 {
			prevNumber, err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				return 0, err
			}
			continue
		}
		currNumber, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		if currNumber > prevNumber {
			increases++
		}
		prevNumber = currNumber
	}
	return increases, nil
}

func deleteEmptyItems(src []string) []string {
	var result []string
	for _, str := range src {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func slidingWindowIncreases(input string) (increases int64, err error) {
	var prevWindowSum, currWindowSum int64
	lines := deleteEmptyItems(strings.Split(input, "\n"))

	if len(lines) < 4 {
		return -1, fmt.Errorf("cannot create at least 2 sliding windows")
	}

	for i := 3; i < len(lines); i++ {
		for j := 0; j < 3; j++ {
			a, err := strconv.ParseInt(lines[i-j], 10, 64)
			if err != nil {
				return -1, err
			}
			b, err := strconv.ParseInt(lines[i-j-1], 10, 64)
			if err != nil {
				return -1, err
			}
			currWindowSum += a
			prevWindowSum += b
		}
		if currWindowSum > prevWindowSum {
			increases++
		}
		currWindowSum = 0
		prevWindowSum = 0
	}
	return increases, nil
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/1/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// count increases per each measurement
	increases, err := countIncreases(string(input))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Regular increases: ", increases)

	// count increases in the sliding window of size 3
	slidingWinIncreases, err := slidingWindowIncreases(string(input))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Sliding window (3 measurements) increases: ", slidingWinIncreases)
}
