package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strconv"
	"strings"
)

// entry represents a single line of input
type entry struct {
	signals        []string       // inputs on the left side
	outputs        []string       // outputs on the right side
	numToSignalMap map[int]string // the final output integer to the string representation
	numToRuneMap   map[int]rune   // segment to wire code mapping
}

// loadFromInput will process a line and load signals on the left side and outputs on the right side
func (e *entry) loadFromInput(input string) {
	so := strings.Split(input, " | ")
	e.signals = strings.Fields(so[0])
	e.outputs = strings.Fields(so[1])
}

// countSimpleNums returns a map with the number of occurrences of nums. 1, 7, 4, 8
func (e *entry) countSimpleNums() map[int]int {
	res := make(map[int]int)
	// 2 segments = num 1
	// 3 segments = num 7
	// 4 segments = num 4
	// 7 segments = num 8
	for _, o := range e.outputs {
		switch len(o) {
		case 2:
			res[2]++
		case 4:
			res[4]++
		case 3:
			res[3]++
		case 7:
			res[7]++
		default:
			{
				// we don't know
			}
		}
	}
	return res
}

// findSimpleNums returns a map of signal mapping to a number 1, 4, 7 and 8
func (e *entry) findSimpleNums() {
	if e.numToSignalMap == nil {
		e.numToSignalMap = make(map[int]string)
	}
	if e.numToRuneMap == nil {
		e.numToRuneMap = make(map[int]rune)
	}
	for _, sig := range e.signals {
		sortedSig := internal.SortString(sig)
		switch len(sortedSig) {
		case 2:
			e.numToSignalMap[1] = sortedSig
		case 4:
			e.numToSignalMap[4] = sortedSig
		case 3:
			e.numToSignalMap[7] = sortedSig
		case 7:
			e.numToSignalMap[8] = sortedSig
		default:
			{
				// we don't know
			}
		}
	}
}

// findAllNums is the main function for task 2
// It loads simple numbers and then deduces all missing numbers, one by one
func (e *entry) findAllNums() {
	/* Here are the numbered segments that will hold a wire rune

	 	  000
		1     2
		1     2
		  333
		4     5
		4	  5
	      666

	Principle of finding all numbers:
	---------------------------------
	- always sort the input signal for less variance
	- find the easy numbers 1, 4, 7, 8
	- find the segment 0 from the difference of num 7 vs num 1
	- find number 6 = the only 6-segment not fully including segments from num 1
	- find the segment 2 as the difference of 8 and 6
	- find the segment 5 as the intersection of 1 and 6

	- identity 5-segment nums:
		num. 3 - has both seg. 2 and 5
		num. 5 - missing seg. 2
		num. 2 - missing seg. 5
	- identify 6-segment nums:
		num. 9 - just 1 different to num. 3 - gives segment 4
		num. 0 - only one left
	- identifying of all segments is not required for this task
	*/
	e.findSimpleNums() // this maps 1, 4, 7, 8 to their string representation
	e.numToRuneMap[0] = findRuneForSegment(e.numToSignalMap[7],
		e.numToSignalMap[1]) // find the segment 0

	// find num 6 = diff with no. 1 returns one element
	for _, sig := range e.signals {
		if len(sig) == 6 {
			diffA, _ := diffSignals(e.numToSignalMap[1], sig)
			if diffA != nil {
				// this must be no. 6
				e.numToSignalMap[6] = internal.SortString(sig)
				break
			}
		}
	}

	// find the segment 2 as difference of 8 and 6
	e.numToRuneMap[2] = findRuneForSegment(e.numToSignalMap[8],
		e.numToSignalMap[6])

	// find the segment #5
	isect := intersectSignals(e.numToSignalMap[1], e.numToSignalMap[6])
	if len(isect) != 1 {
		log.Fatalln("Unexpected signal intersection")
	}
	e.numToRuneMap[5] = isect[0]

	// identify 5-segment nums
	for _, sig := range e.signals {
		if len(sig) != 5 {
			continue
		}
		s := internal.SortString(sig)
		// num 3 fully includes no. 1
		if len(intersectSignals(s, e.numToSignalMap[1])) == 2 {
			e.numToSignalMap[3] = s
			continue
		}
		// num 5
		if isect5 := intersectSignals(s, fmt.Sprint(string(e.numToRuneMap[5]))); isect5 != nil {
			e.numToSignalMap[5] = s
			continue
		}
		// num 2
		if isect2 := intersectSignals(s, fmt.Sprint(string(e.numToRuneMap[2]))); isect2 != nil {
			e.numToSignalMap[2] = s
			continue
		}
	}

	// identify 6-segment numbers
	for _, sig := range e.signals {
		if len(sig) != 6 {
			continue
		}
		s := internal.SortString(sig)
		// we already have num 6, continue
		if s == e.numToSignalMap[6] {
			continue
		} else if len(intersectSignals(s, e.numToSignalMap[3])) == 5 {
			e.numToSignalMap[9] = s
		} else {
			e.numToSignalMap[0] = s
		}
	}
}

// findRuneForSegment identifies a wire signal (rune) for a difference of 2 numbers
func findRuneForSegment(bigger, smaller string) rune {
	leftInA, _ := diffSignals(bigger, smaller)
	if len(leftInA) != 1 {
		log.Fatalln("Unexpected signal comparison")
	}
	return leftInA[0]
}

// diffSignals compares two signals a - b and returns what's left in a (= not contained in b) and what's left in b
func diffSignals(a, b string) (leftInA, leftInB []rune) {
	for _, c := range a {
		r := string(c)
		if !strings.Contains(b, r) {
			leftInA = append(leftInA, c)
		}
	}

	for _, c := range b {
		if !strings.Contains(a, string(c)) {
			leftInB = append(leftInB, c)
		}
	}

	return leftInA, leftInB
}

// diffSignals compares two signals a - b and returns what's left in a (= not contained in b) and what's left in b
func intersectSignals(a, b string) (intersection []rune) {
	for _, c := range a {
		r := string(c)
		if strings.Contains(b, r) {
			intersection = append(intersection, c)
		}
	}
	return intersection
}

// updateTotals is a helper function that updates the 'm' (in-place) map with entry 'e'
func updateTotals(m map[int]int, e map[int]int) {
	for k, v := range e {
		m[k] += v
	}
}

// findSimpleNums is the main function for Task 1.
// It finds all the simple numbers, counts how many times they appeared in the output and returns the total
func findSimpleNums(lines []string) map[int]int {
	totals := make(map[int]int)
	for _, n := range []int{2, 3, 4, 8} {
		totals[n] = 0
	}
	for _, line := range lines {
		var e entry
		e.loadFromInput(line)
		instance := e.countSimpleNums()
		updateTotals(totals, instance)
	}
	return totals
}

// countTotals is a helper function that calculates the total of values from a map
func countTotals(m map[int]int) int {
	totalNums := 0
	for _, v := range m {
		totalNums += v
	}
	return totalNums
}

// findAllNums is the main function for Task 2
// It identifies all numbers and returns the total all values in the output lines
func findAllNums(lines []string) (total int) {
	for _, line := range lines {
		e := entry{}
		e.loadFromInput(line)
		e.findAllNums()

		// reverse the e.numToSignalMap
		signalToNumMap := make(map[string]int)
		for k, v := range e.numToSignalMap {
			signalToNumMap[v] = k
		}
		outStr := ""
		for _, out := range e.outputs {
			sortedOut := internal.SortString(out)
			outStr = fmt.Sprint(outStr, signalToNumMap[sortedOut])
		}
		inc, _ := strconv.Atoi(outStr)
		total += inc
	}
	return total
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/8/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")

	// Task 1: Find the number of occurrences of 1, 4, 7 and 8s
	m := findSimpleNums(lines)
	totalNums := countTotals(m)
	fmt.Println("Task 1: total number of 1, 4, 7 and 8s:", totalNums)

	// Task 2: Find all numbers:
	total := findAllNums(lines)

	fmt.Println("Task 2: Total output value:", total)
}
