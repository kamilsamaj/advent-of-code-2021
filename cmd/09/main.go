package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"sort"
	"strings"
)

// item holds a value and a basinId for a given point in the grid
type item struct {
	val     int
	basinId int
}

// grid stores all numbers loaded from the input
type grid [][]item

type point struct {
	x, y int
}

// loadLines loads lines from the input file and store them in the grid
func (g *grid) loadLines(lines []string) {
	*g = make([][]item, len(lines))
	for i, line := range lines {
		(*g)[i] = make([]item, len(line))
		for j, sym := range line {
			(*g)[i][j].val = int(sym - '0')
		}
	}
}

// size returns a horizontal and vertical size of a grid
func (g *grid) size() (x int, y int) {
	if len(*g) > 0 {
		return len((*g)[0]), len(*g)
	} else {
		return 0, 0
	}
}

// getItem returns a value stored in the grid.
// It abstracts the fact the grid is indexed as [y, x], because of the way of loading input
func (g *grid) getItem(p point) item {
	return (*g)[p.y][p.x]
}

// compRight returns true if the value at point p is smaller than the value right to it
func (g *grid) compRight(p point) bool {
	return g.getItem(p).val < g.getItem(point{p.x + 1, p.y}).val
}

// compLeft compares a value of a point with a point on the left
func (g *grid) compLeft(p point) bool {
	return g.getItem(p).val < g.getItem(point{p.x - 1, p.y}).val
}

// compUp compares a value of a point with a point up
func (g *grid) compUp(p point) bool {
	return g.getItem(p).val < g.getItem(point{p.x, p.y - 1}).val
}

// compDown  compares a value of a point with a point below
func (g *grid) compDown(p point) bool {
	return g.getItem(p).val < g.getItem(point{p.x, p.y + 1}).val
}

// findMins finds all points on a grid that are local minimums (smallest between its neighbors)
func (g *grid) findMins() []point {
	var mins []point
	noItems := len(*g)
	for y, items := range *g {

		lineLen := len(items) // number of ints per items
		for x := range (*g)[y] {
			p := point{x, y}
			isMin := true
			if x < (lineLen - 1) {
				isMin = isMin && g.compRight(p)
			}
			if x > 0 {
				isMin = isMin && g.compLeft(p)
			}
			if y > 0 {
				isMin = isMin && g.compUp(p)
			}
			if y < (noItems - 1) {
				isMin = isMin && g.compDown(p)
			}

			if isMin {
				mins = append(mins, p)
			}
		}
	}
	return mins
}

// markNeighbors is a recursive DFS algorithm that goes through neighbors of a give item.
// It marks all connected items with the same id
func (g *grid) markNeighbors(p point, id int) {
	xSize, ySize := g.size()
	if it := g.getItem(p); it.val == 9 || it.basinId == id {
		// nine is a border and the DFS algorithm ends here
		// node has been already updated
		return
	}
	if bId := g.getItem(p).basinId; bId != 0 && bId != id {
		log.Fatalln("this shouldn't happen", p, id, bId)
	}

	(*g)[p.y][p.x].basinId = id // mark yourself

	// recursively update neighbors in all directions
	if p.x < (xSize - 1) {
		// update right neighbor
		g.markNeighbors(point{p.x + 1, p.y}, id)
	}
	if p.x > 0 {
		// update left neighbor
		g.markNeighbors(point{p.x - 1, p.y}, id)
	}
	if p.y > 0 {
		// update neighbor up
		g.markNeighbors(point{p.x, p.y - 1}, id)
	}
	if p.y < (ySize - 1) {
		// update neighbor below
		g.markNeighbors(point{p.x, p.y + 1}, id)
	}
}

// findBasins goes through each element in the grid and recursively finds its neighbors
// It stops on items that are already marked or of number nines
func (g *grid) findBasins() {
	nextBasinId := 1
	// try to iterate through all items in the grid => you can't skip anything
	for y := range *g {
		for x := range (*g)[y] {
			it := g.getItem(point{x, y})
			if it.val == 9 || it.basinId != 0 {
				continue
			}
			g.markNeighbors(point{x, y}, nextBasinId)
			nextBasinId++
		}
	}
}

// getBasinSizes returns a map of a basin ID and its size. Basin ID 0 is skipped (number nine fields)
func (g *grid) getBasinSizes() (totals map[int]int) {
	totals = make(map[int]int)
	for y := range *g {
		for x := range (*g)[y] {
			it := g.getItem(point{x, y})
			if it.basinId == 0 {
				continue
			}
			totals[it.basinId]++
		}
	}
	return totals
}

// task1 solves the first task
func task1(lines []string) (sum int) {
	g := grid{}
	g.loadLines(lines)
	mins := g.findMins()

	totalRiskLevel := 0
	for _, p := range mins {
		totalRiskLevel += g.getItem(p).val + 1
	}
	return totalRiskLevel
}

// task1 solves the second task
func task2(lines []string) (sum int) {
	g := grid{}
	g.loadLines(lines)
	g.findBasins()
	totals := g.getBasinSizes()

	// sort totals by the basin size (value)
	type kv struct {
		Key   int
		Value int
	}

	var sortedSlice []kv // slice to sort
	for k, v := range totals {
		sortedSlice = append(sortedSlice, kv{k, v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	totalSize := 1
	for _, kv := range sortedSlice[:3] {
		totalSize *= kv.Value
	}

	return totalSize
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
	fmt.Println("Task 2: Sum of 3 largest basins:", task2(lines))

}
