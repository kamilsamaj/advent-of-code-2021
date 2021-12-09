package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

func calcMoves(pos []int, val int) int {
	total := 0
	inc := 0
	for _, p := range pos {
		inc = int(math.Abs(float64(val - p)))
		total += inc
	}
	return total
}

// cheapestPos calculates the best position for crabs and the total number of moves (fuel) they need to do
// pos []int = sorted positions of crabs
func cheapestPos(positions []int) (bestPos, minNoMoves int) {
	med := (len(positions) / 2) - 1 // index of the median item
	increases := 0
	bestPos = -1
	minNoMoves = MaxInt

	// go left and right to see if we didn't find just the local min or max
	// go right
	for i := med; i < len(positions); i++ {
		noMoves := calcMoves(positions, positions[i])
		if noMoves < minNoMoves {
			minNoMoves = noMoves
			bestPos = positions[i]
		} else {
			increases++
		}
		if increases > 20 {
			break
		}
	}

	// go left
	increases = 0
	for i := med - 1; i >= 0; i-- {
		p := calcMoves(positions, i)
		if p < minNoMoves {
			minNoMoves = p
			bestPos = i
		} else {
			increases++
		}
		if increases > 20 {
			break
		}
	}
	return bestPos, minNoMoves
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/7/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// process the input and load crabs' positions
	origPosStr := strings.Split(strings.Trim(string(input), "\n"), ",")
	var origPos []int
	for _, s := range origPosStr {
		n, _ := strconv.Atoi(s)
		fmt.Println(n)
		origPos = append(origPos, n)
	}
	// the ideal position should be somewhere around the median
	sort.Ints(origPos) // sort positions in place
	bestPos, noMoves := cheapestPos(origPos)
	fmt.Printf("Best position: %d, no. moves: %d", bestPos, noMoves)
}
