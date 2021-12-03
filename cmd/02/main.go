package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strconv"
	"strings"
)

// Vector represents the "forward" as "x" and "depth" as "y"
type Vector struct {
	x int64
	y int64
}

func calculateVector(cmd string) (Vector, error) {
	// identify the command and update the correct coordinate
	if strings.HasPrefix(cmd, "forward") {
		posDelta, err := strconv.ParseUint(strings.Fields(cmd)[1], 10, 64)
		if err != nil {
			return Vector{}, err
		}
		return Vector{int64(posDelta), 0}, nil
	} else if strings.HasPrefix(cmd, "down") {
		posDelta, err := strconv.ParseUint(strings.Fields(cmd)[1], 10, 64)
		if err != nil {
			return Vector{}, err
		}
		return Vector{0, int64(posDelta)}, nil
	} else if strings.HasPrefix(cmd, "up") {
		posDelta, err := strconv.ParseUint(strings.Fields(cmd)[1], 10, 64)
		if err != nil {
			return Vector{}, err
		}
		return Vector{0, -int64(posDelta)}, nil
	} else {
		return Vector{}, fmt.Errorf("could not parse the cmd = '%s'", cmd)
	}
}

func calculatePosition(input *[]byte) (position int64, err error) {
	var totalPos Vector

	// split the byte array by the line break '\n' = 0x0A
	var cmdBytes []byte

	for _, b := range *input {
		// read the input until you find the line break
		if b == 0x0a {
			cmdStr := string(cmdBytes)
			cmdBytes = []byte{}
			v, err := calculateVector(cmdStr)
			if err != nil {
				return 0, err
			}
			totalPos.x += v.x
			totalPos.y += v.y

		} else {
			// keep building the command []byte array
			cmdBytes = append(cmdBytes, b)
		}
	}

	// handle a case when the input is not terminated by a line break
	if len(cmdBytes) > 0 {
		v, err := calculateVector(string(cmdBytes))
		if err != nil {
			return 0, err
		}
		totalPos.x += v.x
		totalPos.y += v.y
	}
	return totalPos.x * totalPos.y, nil
}

func calculatePositionWithAim(input *[]byte) (position int64, err error) {
	var totalPos Vector
	var aim int64

	// split the byte array by the line break '\n' = 0x0A
	var cmdBytes []byte

	for _, b := range *input {
		// read the input until you find the line break
		if b == 0x0a {
			cmdStr := string(cmdBytes)
			cmdBytes = []byte{}
			v, err := calculateVector(cmdStr)
			if err != nil {
				return 0, err
			}
			// if only the "y" was set, update the aim, but don't change "forward" or "depth" metrics
			if v.x == 0 {
				aim += v.y
			} else {
				totalPos.x += v.x
				totalPos.y += v.x * aim
			}
		} else {
			// keep building the command []byte array
			cmdBytes = append(cmdBytes, b)
		}
	}

	// handle a case when the input is not terminated by a line break
	if len(cmdBytes) > 0 {
		v, err := calculateVector(string(cmdBytes))
		if err != nil {
			return 0, err
		}
		// if only the "y" was set, update the aim, but don't change "forward" or "depth" metrics
		if v.x != 0 {
			totalPos.x += v.x
			totalPos.y += v.x * aim
		}
	}
	return totalPos.x * totalPos.y, nil
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/2/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// Task 1 - calculate position of the submarine
	position, err := calculatePosition(&input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Task 1: Position = %d\n", position)

	// Task 2 - calculate the position with aim
	position, err = calculatePositionWithAim(&input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Task 2: Position = %d\n", position)
}
