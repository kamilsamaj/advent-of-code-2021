package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold struct {
	axisName  string
	foldIndex int
}

type grid struct {
	points []point
	folds  []fold
}

func (g *grid) load(lines []string) {
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "fold") {
			r := regexp.MustCompile(`^fold along (\w+)=(\d{1,5})$`)
			fields := r.FindStringSubmatch(line)
			axisName := fields[1]
			foldIndex, _ := strconv.Atoi(fields[2])
			g.folds = append(g.folds, fold{axisName, foldIndex})
		} else {
			coord := strings.Split(strings.Trim(line, "\n"), ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			g.points = append(g.points, point{x, y})
		}
	}
}

func (g *grid) fold(axis string, index int) map[point]bool {
	var newPoints map[point]bool = make(map[point]bool)

	for _, p := range g.points {
		if axis == "x" {
			if p.x >= index {
				newPoints[point{p.x - 2*(p.x-index), p.y}] = true
			} else {
				newPoints[p] = true
			}
		} else if axis == "y" {
			if p.y >= index {
				newPoints[point{p.x, p.y - 2*(p.y-index)}] = true
			} else {
				newPoints[p] = true
			}
		} else {
			log.Fatalln("i don't know anything about axis", axis)
		}
	}
	return newPoints
}

func task1(lines []string) int {
	g := grid{}
	g.load(lines)
	newPoints := g.fold(g.folds[0].axisName, g.folds[0].foldIndex)
	return len(newPoints)
}

func task2(lines []string) int {
	return 0
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/13/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	fmt.Println("Task 1:", task1(lines))
	fmt.Println("Task 2:", task2(lines))
}
