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

func avg(pos []int) float64 {
	total := 0
	for _, p := range pos {
		total += p
	}
	return float64(len(pos)) / float64(total)
}

func calcMovesArithmProgression(pos []int, val int) int {
	total := 0
	inc := 0
	for _, p := range pos {
		absAnA1 := int(math.Abs(float64(val - p)))
		inc = absAnA1 * (1 + absAnA1) / 2
		total += inc
	}
	return total
}

// cheapestPos calculates the best position for crabs and the total number of moves (fuel) they need to do
// pos []int = sorted positions of crabs
func cheapestPos(positions []int) (bestPos, minFuel int) {
	med := (len(positions) / 2) - 1 // index of the median item
	increases := 0
	bestPos = -1
	minFuel = MaxInt

	// go left and right to see if we didn't find just the local min or max
	// go right
	for i := med; i < len(positions); i++ {
		fuel := calcMoves(positions, positions[i])
		if fuel < minFuel {
			minFuel = fuel
			bestPos = positions[i]
			increases = 0
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
		if p < minFuel {
			minFuel = p
			bestPos = i
		} else {
			increases++
		}
		if increases > 20 {
			break
		}
	}
	return bestPos, minFuel
}

// cheapestArithmPos calculates the best position for crabs and the total number of moves (fuel) they need to do
// pos []int = sorted positions of crabs
// the fuel price increases with every step = arithmetic progression
func cheapestArithmPos(positions []int) (bestPos, minFuel int) {
	average := avg(positions) // index of the median item
	increases := 0
	bestPos = -1
	minFuel = MaxInt

	// go left and right to see if we didn't find just the local min or max
	// go right
	for i := int(average); i < positions[len(positions)-1]; i++ {
		fuel := calcMovesArithmProgression(positions, i)
		if fuel < minFuel {
			minFuel = fuel
			bestPos = i
			increases = 0
		} else {
			increases++
		}
		if increases > 20 {
			break
		}
	}

	// go left
	increases = 0
	for i := int(average) - 1; i >= 0; i-- {
		fuel := calcMovesArithmProgression(positions, i)
		if fuel < minFuel {
			minFuel = fuel
			bestPos = i
		} else {
			increases++
		}
		if increases > 20 {
			break
		}
	}
	return bestPos, minFuel
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/7/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// process the input and load crabs' positions
	// be careful about the linebreak in the last number
	origPosStr := strings.Split(strings.Trim(string(input), "\n"), ",")
	var origPos []int
	for _, s := range origPosStr {
		n, _ := strconv.Atoi(s)
		origPos = append(origPos, n)
	}
	// the ideal position should be somewhere around the median
	sort.Ints(origPos) // sort positions in place

	// Task 1
	bestPos, fuel := cheapestPos(origPos)
	fmt.Printf("Task 1: Best position: %d, fuel: %d\n", bestPos, fuel)

	// Task 2
	bestPos2, fuel2 := cheapestArithmPos(origPos)
	fmt.Printf("Task 2: Best position: %d, fuel: %d\n", bestPos2, fuel2)
}
