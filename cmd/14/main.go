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

func (p *polymer) loadCache(lines []string, noSteps int) {
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
		p.cache[match[1]] = string(match[1][0]) + match[2] + string(match[1][1])
	}
}

func (p *polymer) getTaskResult() int {
	var resMap = make(map[string]int)
	for _, s := range p.val {
		resMap[string(s)]++
	}
	type res struct {
		val    int
		letter string
	}
	min := res{MaxInt, ""}
	max := res{0, ""}

	for k, v := range resMap {
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

func (p *polymer) expand(polStr string) string {
	if len(polStr) == 1 {
		return polStr
	} else if val, ok := p.cache[polStr]; ok {
		return val
	} else {
		middleIndex := len(polStr) / 2
		s1 := polStr[:middleIndex]
		s2 := polStr[middleIndex-1 : middleIndex+1]
		s3 := polStr[middleIndex:]
		p1 := p.expand(s1)
		p2 := p.expand(s2)
		p3 := p.expand(s3)
		if _, ok := p.cache[s1]; !ok {
			p.cache[s1] = p1
		}
		if _, ok := p.cache[s2]; !ok {
			p.cache[s2] = p2
		}
		if _, ok := p.cache[s3]; !ok {
			p.cache[s3] = p3
		}
		mergedPolymer := p1[:len(p1)-1] + p2[:len(p2)-1] + p3
		if _, ok := p.cache[mergedPolymer]; !ok {
			p.cache[polStr] = mergedPolymer
		}
		return mergedPolymer
	}
}

func expandPolymer(lines []string, noSteps int) int {
	/*
		Outline of the algorithm
		- try to cache results
		- iterate on a single pair as deep as you can on a tuple
	*/
	var p polymer
	p.loadCache(lines, noSteps)
	for i := 0; i < noSteps; i++ {
		p.val = p.expand(p.val)
		fmt.Println("iteration", i+1)
	}
	// calculate result
	return p.getTaskResult()
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
	fmt.Println("Task 2:", expandPolymer(lines, 40))
}
