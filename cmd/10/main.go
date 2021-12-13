package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strings"
)

func task1(lines []string) int {
	var m = make(map[rune]int)
	totalErrorScore := 0
	m['('] = 0
	m['{'] = 0
	m['['] = 0
	m['<'] = 0

	for i, line := range lines {
		for j, r := range line {
			errScore := 0
			switch r {
			case ')':
				m['(']--
				errScore = 3
			case '}':
				m['{']--
				errScore = 1197
			case ']':
				m['[']--
				errScore = 57
			case '>':
				m['<']--
				errScore = 25137
			case '(', '{', '[', '<':
				m[r]++
			default:
				log.Fatalln("unknown character:", string(r))
			}
			if m[r] < 0 {
				fmt.Printf("line no. %d '%s' is invalid. Found %v at position %d",
					i, line, r, j)
				totalErrorScore += errScore
			}
		}
	}
	return totalErrorScore
}

func task2(lines []string) int {
	return 0
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/10/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	fmt.Println("Task 1: Sum of error scores:", task1(lines))
	//fmt.Println("Task 2: Sum of 3 largest basins:", task2(lines))

}
