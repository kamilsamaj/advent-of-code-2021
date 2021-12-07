package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strconv"
	"strings"
)

type boardItem struct {
	value   int
	checked bool
}

type board [5][5]boardItem

func (b *board) markItem(num int) (noMarked int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].value == num {
				noMarked++
				b[i][j].checked = true
			}
		}
	}
	return noMarked
}

func (b *board) hasBingo() bool {
	// check lines
	for i := 0; i < 5; i++ {
		lineAllTrue := true
		for j := 0; j < 5; j++ {
			if !b[i][j].checked {
				lineAllTrue = false
				break
			}
		}
		if lineAllTrue {
			return true
		}
	}

	// check columns
	for i := 0; i < 5; i++ {
		columnAllTrue := true
		for j := 0; j < 5; j++ {
			if !b[j][i].checked {
				columnAllTrue = false
				break
			}
		}
		if columnAllTrue {
			return true
		}
	}
	return false
}

func (b *board) getScore(multiplier int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b[i][j].checked {
				sum += b[i][j].value
			}
		}
	}
	return sum * multiplier
}

func lineToInts(line string) (res [5]boardItem, err error) {
	for i, numStr := range strings.Fields(line) {
		val, err := strconv.Atoi(numStr)
		if err != nil {
			return res, fmt.Errorf("lineToInts: couldn't parse line %s", line)
		}
		res[i] = boardItem{value: val, checked: false}
	}
	return
}

func getBoards(lines []string) (boards []board, err error) {
	lineCounter := 0 // line counter to fill in the b
	var b board
	for _, line := range lines {
		if line == "" {
			continue
		}
		b[lineCounter], err = lineToInts(line)
		if err != nil {
			return boards, fmt.Errorf("%v", err)
		}
		lineCounter++

		// full b has been parsed
		if lineCounter == 5 {
			lineCounter = 0
			boards = append(boards, b)
			b = board{}
		}
	}
	return
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/4/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// process the input
	lines := strings.Split(string(input), "\n")

	// get the random number input
	var randNumbers []int // input with random numbers
	for _, num := range strings.Split(lines[0], ",") {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalln(err)
		}
		randNumbers = append(randNumbers, numInt)
	}

	// get the boards
	boards, err := getBoards(lines[1:])
	if err != nil {
		log.Fatalln(err)
	}

	// Task 1: Bingo - which board wins first?
	// try to check each number on check board and always check if bingo was found
	boards1 := make([]board, len(boards)) // these lines will be reduced based on the most common bits
	copy(boards1, boards)
out: // this is a label to break from the for loops
	for j, num := range randNumbers {
		for i := range boards1 {
			_ = boards1[i].markItem(num)
			if boards1[i].hasBingo() {
				fmt.Printf("Bingo found in board #%d after searching through %d numbers:\n%v\n", i, j, boards1[i])
				fmt.Printf("Task 1: Final score: %d\n\n", boards1[i].getScore(num))
				break out
			}
		}
	}

	// Task 2: Bingo - which board wins last?
	// we can use []boards directly
	var i int
	var worstBoard struct {
		b      *board
		winsAt int
	}
	{
	}

	// the logic is reverse here - go through all boards and apply all number until the board gets a bingo
	// then just remember the worst board
	for i = range boards {
		boardFinished := false
		for j, num := range randNumbers {
			_ = boards[i].markItem(num)
			if boards[i].hasBingo() {
				boardFinished = true
				if j > worstBoard.winsAt {
					// we have a new worst board
					worstBoard.b = &boards[i]
					worstBoard.winsAt = j
				}
				//fmt.Printf("Bingo found in board #%d after searching through %d numbers:\n%v\n", i, j, boards[i])
				//fmt.Printf("Task 1: Final score: %d", boards[i].getScore(num))
				break
			}
		}
		// if all numbers were applied and the board doesn't win, it's yours
		if !boardFinished {
			fmt.Println("this board never closes: ", boards[i])
			worstBoard.b = &boards[i]
			worstBoard.winsAt = len(randNumbers) - 1
			break
		}
	}

	fmt.Printf("Task 2: The worst board found after searching through %d numbers:\n%v\n", worstBoard.winsAt, *worstBoard.b)
	fmt.Printf("Task 2: Final score: %d", worstBoard.b.getScore(randNumbers[worstBoard.winsAt]))
}
