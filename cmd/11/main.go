package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"math"
	"strings"
)

const arrSize = 10

type octopus struct {
	val        int8
	flashedNow bool
}

type grid [arrSize][arrSize]octopus

func (g *grid) loadFromLines(lines []string) {
	for y, line := range lines {
		for x, r := range line {
			(*g)[x][y] = octopus{val: int8(r - '0')}
		}
	}
}

//
//func (g *grid) formatGridForPrint() string {
//	var sb strings.Builder
//	for i := 0; i < arrSize; i++ {
//		for j := 0; j < arrSize; j++ {
//			sb.WriteString(fmt.Sprint((*g)[j][i].val % 10))
//		}
//		sb.WriteString("\n")
//	}
//	return sb.String()
//}

func (g *grid) addRound() (flashes int) {
	flashesPerRound := 0
	// flash
	for y := 0; y < arrSize; y++ {
		for x := 0; x < arrSize; x++ {
			if it := (*g)[x][y]; it.val >= 9 && !it.flashedNow {
				(*g)[x][y].val++
				(*g)[x][y].flashedNow = true
				g.flashOnNeighbor(x, y)
			} else if (*g)[x][y].flashedNow {
				// skip
			} else {
				(*g)[x][y].val++
			}
		}
	}
	for y := 0; y < arrSize; y++ {
		for x := 0; x < arrSize; x++ {
			if val := (*g)[x][y].val; val == 10 {
				flashesPerRound++
				(*g)[x][y].val = 0
				(*g)[x][y].flashedNow = false
			}
		}
	}
	return flashesPerRound
}

func (g *grid) flashOnNeighbor(x, y int) {
	// bump 3x3 sub-array -> limited by the grid ends
	x1 := int(math.Max(float64(x-1), 0))
	y1 := int(math.Max(float64(y-1), 0))
	x2 := int(math.Min(float64(x+1), arrSize-1))
	y2 := int(math.Min(float64(y+1), arrSize-1))

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if val := (*g)[i][j].val; val < 10 {
				(*g)[i][j].val++
				if val := (*g)[i][j].val; val == 10 && !(*g)[i][j].flashedNow {
					(*g)[i][j].flashedNow = true
					g.flashOnNeighbor(i, j)
				}
			}
		}
	}
}

func task1(lines []string, steps int) int {
	totalFlashes := 0
	g := grid{}
	g.loadFromLines(lines)
	for i := 0; i < 100; i++ {
		totalFlashes += g.addRound()
	}
	return totalFlashes
}

func task2(lines []string) int {
	return 0
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/11/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	fmt.Println("Task 1: Sum of error scores:", task1(lines, 100))
	fmt.Println("Task 2: Median error score of the missing closings:", task2(lines))
}
