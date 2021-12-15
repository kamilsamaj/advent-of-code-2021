package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"sort"
	"strings"
)

func containsRune(s []rune, r rune) bool {
	for _, sym := range s {
		if r == sym {
			return true
		}
	}
	return false
}

func task1(lines []string) int {
	var m = make(map[rune]rune)
	totalErrorScore := 0
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	m['>'] = '<'

	var openingSymbols, closingSymbols, readSymbols []rune
	for k, v := range m {
		closingSymbols = append(closingSymbols, k)
		openingSymbols = append(openingSymbols, v)
	}
	errScore := 0
	for _, line := range lines {
		for _, r := range line {
			if containsRune(openingSymbols, r) {
				// if 'r' is an opening bracket, push it to the read symbols LIFO
				readSymbols = append(readSymbols, r)
				continue
			} else if containsRune(closingSymbols, r) {
				if len(readSymbols) > 0 && readSymbols[len(readSymbols)-1] == m[r] {
					readSymbols = readSymbols[:len(readSymbols)-1] // pop the last symbol in the LIFO
				} else {
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
					totalErrorScore += errScore
					break
				}
			}
		}
	}
	return totalErrorScore
}

func task2(lines []string) int {
	// map closing symbol to the opening symbol
	var m = make(map[rune]rune)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	m['>'] = '<'

	// and map the opening symbol to the closing one too
	var mRev = make(map[rune]rune)
	for k, v := range m {
		mRev[v] = k
	}

	// final scores for a missing closing
	var scores = make(map[rune]int)
	scores[')'] = 1
	scores[']'] = 2
	scores['}'] = 3
	scores['>'] = 4

	var lineScores []int // scores of complete lines
	var openingSymbols, closingSymbols, readSymbols []rune
	for k, v := range m {
		closingSymbols = append(closingSymbols, k)
		openingSymbols = append(openingSymbols, v)
	}

	for _, line := range lines {
		errorScore := 0
		readSymbols = []rune{} // LIFO buffer for read characters
		lineFullyRead := true
		for _, r := range line {
			if containsRune(openingSymbols, r) {
				// if 'r' is an opening bracket, push it to the read symbols LIFO
				readSymbols = append(readSymbols, r)
				continue
			} else if containsRune(closingSymbols, r) {
				if len(readSymbols) > 0 && readSymbols[len(readSymbols)-1] == m[r] {
					readSymbols = readSymbols[:len(readSymbols)-1] // pop the last symbol in the LIFO
				} else {
					lineFullyRead = false
					break // don't consider lines that are incomplete
				}
			}
		}

		if !lineFullyRead {
			continue
		}

		// now finish the line
		for x := len(readSymbols) - 1; x >= 0; x-- {
			symScore := scores[mRev[readSymbols[x]]]
			errorScore = errorScore*5 + symScore
		}
		lineScores = append(lineScores, errorScore)
	}
	sort.Slice(lineScores, func(i, j int) bool {
		return lineScores[i] < lineScores[j]
	})
	return lineScores[len(lineScores)/2]
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
	fmt.Println("Task 2: Median error score of the missing closings:", task2(lines))
}
