package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"regexp"
	"strings"
)

type polymer struct {
	mappings map[string]string
	val      string
}

func (p *polymer) load(lines []string) {
	p.mappings = make(map[string]string)
	for i, line := range lines {
		trimmedLine := strings.Trim(line, "\n")
		if trimmedLine == "" {
			continue
		}
		if i == 0 {
			p.val = trimmedLine
			continue
		}

		r := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
		match := r.FindStringSubmatch(line)
		p.mappings[match[1]] = match[2]
	}
}

func task1(lines []string, noSteps int) int {
	/*
		Outline of the algorithm
		- allocate []string big enough, grow when needed
		- iterate with +1 step
		- have a function that just tries to insert
			- either it returns the same string or extended
		- always insert with last first item skipped
	*/
	var p polymer
	p.load(lines)
	for i := 0; i < noSteps; i++ {
		var newPolymer strings.Builder
		for pos := 0; pos < (len(p.val) - 1); pos++ {
			tuple := p.val[pos : pos+2]
			if v, ok := p.mappings[tuple]; ok {
				// insert the expanded polymer - skip the last one
				newPolymer.WriteString(fmt.Sprint(string(p.val[pos]), v))
			}
		}
		// we skipped the last piece from the polymer, insert it too
		newPolymer.WriteString(string(p.val[len(p.val)-1]))
		p.val = newPolymer.String()
	}
	return 0
}

func task2(lines []string) int {
	return 0
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/14/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	fmt.Println("Task 1:", task1(lines, 10))
	fmt.Println("Task 2: Code to escape")
	task2(lines)
}
