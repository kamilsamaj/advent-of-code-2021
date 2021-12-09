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

type Ocean [9]int

func (o *Ocean) growOcean(days int) {
	for i := 0; i < days; i++ {
		origZeros := o[0] // remember the original number of fish with 0 days to reproduction
		for j := 1; j < 9; j++ {
			// all fish go down by 1 day
			o[j-1] = o[j]
		}
		// fish that have 0 days to reproduction create the same number of fish with 8 days and then become fish with 6 days
		o[6] += origZeros
		o[8] = origZeros
	}
}

func (o *Ocean) totalFish() int64 {
	var total int64 = 0
	for i := 0; i < 9; i++ {
		total += int64(o[i])
	}
	return total
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
	var fish []Fish
	for _, numStr := range strings.Split(inpLine, ",") {
		if numStr == "" {
			continue
		}
		num, _ := strconv.Atoi(numStr)
		fish = append(fish, Fish{uint8(num)})
	}

	// grow the fish for 80 days - very naive algorithm that showed up to don't scale
	var fish1 []Fish = make([]Fish, len(fish))
	copy(fish1, fish)
	for i := 0; i < 80; i++ {
		var newfish []Fish
		for j := range fish1 {
			newFish := fish1[j].decr()
			if newFish != nil {
				newfish = append(newfish, *newFish)
			}
		}
		fish1 = append(fish1, newfish...)
	}
	fmt.Println("Task 1: No. fish after 80 days:", len(fish1))

	// Task 2 - we need to reduce the no. fish to categories
	ocean := Ocean([9]int{})
	for _, fish := range fish {
		ocean[fish.timer]++
	}
	ocean.growOcean(256)
	fmt.Println("Task 2: No. fish after 256 days:", ocean.totalFish())
}
