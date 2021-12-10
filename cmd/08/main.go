package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strings"
)

type entry struct {
	signals []string
	outputs []string
}

func (e *entry) loadFromInput(input string) {
	so := strings.Split(input, " | ")
	e.signals = strings.Fields(so[0])
	e.outputs = strings.Fields(so[1])
}

func (e *entry) findSimpleNums() map[int]int {
	res := make(map[int]int)
	for _, n := range []int{2, 3, 4, 8} {
		res[n] = 0
	}

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

func updateTotals(m map[int]int, instance map[int]int) {
	for k, v := range instance {
		m[k] += v
	}
}

func findSimpleNums(lines []string) map[int]int {
	totals := make(map[int]int)
	for _, n := range []int{2, 3, 4, 8} {
		totals[n] = 0
	}
	for _, line := range lines {
		var e entry
		e.loadFromInput(line)
		instance := e.findSimpleNums()
		updateTotals(totals, instance)
	}
	return totals
}

func countTotals(m map[int]int) int {
	totalNums := 0
	for _, v := range m {
		totalNums += v
	}
	return totalNums
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
}
