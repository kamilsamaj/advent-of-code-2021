package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const boardSize = 1000

type Board [boardSize][boardSize]Point

func (b *Board) drawLine(line Line) bool {
	if line.start.x == line.end.x {
		// vertical line
		startIndex := int(math.Min(float64(line.start.y), float64(line.end.y))) // start from the smaller index
		noSteps := int(math.Abs(float64(line.end.y-line.start.y))) + 1
		for i := 0; i < noSteps; i++ {
			b[line.start.x][startIndex+i].val++
		}
		return true
	} else if line.start.y == line.end.y {
		// horizontal line
		startIndex := int(math.Min(float64(line.start.x), float64(line.end.x))) // start from the smaller index
		noSteps := int(math.Abs(float64(line.end.x-line.start.x))) + 1
		for i := 0; i < noSteps; i++ {
			b[startIndex+i][line.start.y].val++
		}
	} else {
		return false
	}
	return true
}

type Point struct {
	x, y int
	val  int
}

type Line struct {
	start, end Point
}

func NewLine(x1, y1, x2, y2 string) Line {
	x1i, _ := strconv.Atoi(x1)
	x2i, _ := strconv.Atoi(x2)
	y1i, _ := strconv.Atoi(y1)
	y2i, _ := strconv.Atoi(y2)
	return Line{start: Point{x1i, y1i, 0}, end: Point{x2i, y2i, 0}}
}

func (b *Board) countPoints(minVal int) int {
	noPoints := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j].val >= minVal {
				noPoints++
			}
		}
	}
	return noPoints
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/5/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// process the input
	inputLines := strings.Split(string(input), "\n")
	//var lines []Line
	b := Board{}
	r := regexp.MustCompile(`(?P<x1>\d{1,4}),(?P<y1>\d{1,4}) -> (?P<x2>\d{1,4}),(?P<y2>\d{1,4})`) // precompile the regexp
	for _, inpLine := range inputLines {
		if inpLine == "" {
			continue
		}
		fields := r.FindStringSubmatch(inpLine)
		line := NewLine(fields[1], fields[2], fields[3], fields[4])

		// check for vertical or horizontal lines happens in the b.drawLine()

		b.drawLine(line)
		//lines = append(lines, line)
	}
	fmt.Println("Task 1: OverLap score:", b.countPoints(2))
}
