package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strconv"
	"strings"
)

type Fish struct {
	timer uint8
}

func (f *Fish) decr() *Fish {
	if f.timer == 0 {
		f.timer = 6
		newFish := Fish{8}
		return &newFish
	} else {
		f.timer--
		return nil
	}
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/6/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// process the input and load all fish
	inpLine := strings.Split(string(input), "\n")[0]
	var fishes []Fish
	for _, numStr := range strings.Split(inpLine, ",") {
		if numStr == "" {
			continue
		}
		num, _ := strconv.Atoi(numStr)
		fishes = append(fishes, Fish{uint8(num)})
	}

	// grow the fishes for 80 days
	for i := 0; i < 256; i++ {
		var newFishes []Fish
		for j := range fishes {
			newFish := fishes[j].decr()
			if newFish != nil {
				newFishes = append(newFishes, *newFish)
			}
		}
		fmt.Printf("Counted %d days\n", i)
		fishes = append(fishes, newFishes...)
	}
	fmt.Println("Task 1: No. fish after 80 days:", len(fishes))
}
