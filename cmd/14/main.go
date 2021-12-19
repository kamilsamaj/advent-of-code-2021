package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"regexp"
	"strings"
)

type polymer struct {
	rules         map[string]string
	originalInput string
	resultCache   map[string]map[string]int // for example: resultCache["NN_2"] = {"N": 2, "B": 4, "C": 1...}
}

const MaxInt = int((^uint(0)) >> 1)

func (p *polymer) load(lines []string, noSteps int) {
	p.rules = make(map[string]string)
	for i, line := range lines {
		trimmedLine := strings.Trim(line, "\n")
		if trimmedLine == "" {
			continue
		}
		if i == 0 {
			p.originalInput = trimmedLine
			continue
		}

		r := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
		match := r.FindStringSubmatch(line)
		p.rules[match[1]] = match[2]
	}
}

// getTaskResult finds the most and the least common character and returns their result
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

// expand gets a tuple of characters and recursively counts what they add
func (p *polymer) expand(tuple string, noSteps int, characterCounts map[string]int) map[string]int {
	if noSteps == 0 {
		return nil
	}
	var updated = make(map[string]int) // what characters this tuple adds (doesn't include the tuple itself)

	cacheKeyName := fmt.Sprint(tuple, "_", noSteps) // same results = same tuple with the same number of missing steps

	// if cache already exists, return is right away
	if counts, ok := p.resultCache[cacheKeyName]; ok {
		// cache doesn't contain the tuple itself to avoid duplicates when expanding the tree
		for k, v := range counts {
			characterCounts[k] += v
		}
		return counts
	}

	insertedChar := p.rules[tuple]
	characterCounts[insertedChar]++ // global counters
	updated[insertedChar]++         // cache for the tuple expansion

	// recursively expand the other steps
	r1 := p.expand(string(tuple[0])+insertedChar, noSteps-1, characterCounts)
	r2 := p.expand(insertedChar+string(tuple[1]), noSteps-1, characterCounts)
	for k, v := range r1 {
		updated[k] += v
	}
	for k, v := range r2 {
		updated[k] += v
	}

	p.resultCache[cacheKeyName] = updated // save the result to the cache
	return updated
}

func expandPolymer(lines []string, noSteps int) int {
	/*
		The algorithm expands each tuple of character by `noSteps` and only updates the character counters
		It needs a cache that remembers the expansion (=what a subtree adds)
	*/
	var p polymer
	p.resultCache = make(map[string]map[string]int)
	p.load(lines, noSteps)
	var characterCounts = make(map[string]int)

	// recursively expand the tuples but just update the characterCounts - don't expand the string
	for i := 0; i < len(p.originalInput)-1; i++ {
		characterCounts[string(p.originalInput[i])]++ // global results
		p.expand(p.originalInput[i:i+2], noSteps, characterCounts)
	}
	characterCounts[string(p.originalInput[len(p.originalInput)-1])]++ // last item is missed
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
	fmt.Println("Task 1:", expandPolymer(lines, 10))
	fmt.Println("Task 2:", expandPolymer(lines, 40))
}
