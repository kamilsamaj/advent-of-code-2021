package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strings"
)

// grid stores numbers loaded from the input
// Because the input is loaded by lines (by y-axis), first index is y, second is x
type grid [][]int

type point struct {
	x, y int
}

// loadLines loads lines from the input file and store them in the grid
func (g *grid) loadLines(lines []string) {
	*g = make([][]int, len(lines))
	for i, line := range lines {
		(*g)[i] = make([]int, len(line))
		for j, sym := range line {
			(*g)[i][j] = int(sym - '0')
		}
	}
}

// getPoint returns a value stored in the grid.
// It abstracts the fact the grid is indexed as [y, x], because of the way of loading input
func (g *grid) getPoint(p point) int {
	return (*g)[p.y][p.x]
}

// compRight returns true if the value at point p is smaller than the value right to it
func (g *grid) compRight(p point) bool {
	if g.getPoint(p) < g.getPoint(point{p.x + 1, p.y}) {
		return true
	}
	return false
}

func (g *grid) compLeft(p point) bool {
	if g.getPoint(p) < g.getPoint(point{p.x - 1, p.y}) {
		return true
	}
	return false
}

func (g *grid) compUp(p point) bool {
	if g.getPoint(p) < g.getPoint(point{p.x, p.y - 1}) {
		return true
	}
	return false
}

func (g *grid) compDown(p point) bool {
	if g.getPoint(p) < g.getPoint(point{p.x, p.y + 1}) {
		return true
	}
	return false
}

func (g *grid) findMins() []point {
	var mins []point
	noLines := len(*g)
	for i, line := range *g {

		lineLen := len(line) // number of ints per line
		for j, _ := range (*g)[i] {
			p := point{j, i}
			isMin := true
			if j < (lineLen - 1) {
				isMin = isMin && g.compRight(p)
			}
			if j > 0 {
				isMin = isMin && g.compLeft(p)
			}
			if i > 0 {
				isMin = isMin && g.compUp(p)
			}
			if i < (noLines - 1) {
				isMin = isMin && g.compDown(p)
			}

			if isMin {
				mins = append(mins, p)
			}
		}
	}
	return mins
}

// task1 solves the first task
func task1(lines []string) (sum int) {
	g := grid{}
	g.loadLines(lines)
	mins := g.findMins()

	totalRiskLevel := 0
	for _, p := range mins {
		totalRiskLevel += g.getPoint(p) + 1
	}
	return totalRiskLevel
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/9/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	fmt.Println("Task 1: Sum of risk levels:", task1(lines))
}
