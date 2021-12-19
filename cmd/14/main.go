package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"regexp"
	"strings"
)

type polymer struct {
	cache map[string]string
	val   string
}

const MaxInt = int((^uint(0)) >> 1)

func (p *polymer) load(lines []string, noSteps int) {
	p.cache = make(map[string]string)
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
		p.cache[match[1]] = match[2]
	}
}

func (p *polymer) getTaskResult(characterCounts map[string]int) int {
	type res struct {
		val    int
		letter string
	}
	min := res{MaxInt, ""}
	max := res{0, ""}

	for k, v := range characterCounts {
		if v > max.val {
			max.val = v
			max.letter = k
		}
		if v < min.val {
			min.val = v
			min.letter = k
		}
	}
	return max.val - min.val
}

func (p *polymer) expand(tuple string, noSteps int, characterCounts map[string]int) {
	if noSteps == 0 {
		return
	}
	insertedChar := p.cache[tuple]
	characterCounts[insertedChar]++

	p.expand(string(tuple[0])+insertedChar, noSteps-1, characterCounts)
	p.expand(insertedChar+string(tuple[1]), noSteps-1, characterCounts)
}

func expandPolymer(lines []string, noSteps int) int {
	/*
		Outline of the algorithm
		- try to cache results
		- iterate on a single pair as deep as you can on a tuple
	*/
	var p polymer
	p.load(lines, noSteps)
	var characterCounts = make(map[string]int)

	// recursively expand the tuples but just update the characterCounts - don't expand the string
	for i := 0; i < len(p.val)-1; i++ {
		characterCounts[string(p.val[i])]++
		p.expand(p.val[i:i+2], noSteps, characterCounts)
		fmt.Println(i)
	}
	characterCounts[string(p.val[len(p.val)-1])]++ // last item is missed
	// calculate result
	return p.getTaskResult(characterCounts)
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/14/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	//fmt.Println("Task 1:", expandPolymer(lines, 10))
	fmt.Println("Task 2:", expandPolymer(lines, 30))
}
