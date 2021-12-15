package main

import (
	"fmt"
	"github.com/kamilsamaj/advent-of-code-2021/internal"
	"log"
	"strings"
	"unicode"
)

type graph struct {
	nodes          map[string][]string // map node names with its neighbors
	visitedPaths   map[string]bool
	finalizedPaths map[string]bool
}

func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func (g *graph) load(lines []string) {
	g.nodes = make(map[string][]string)
	g.visitedPaths = make(map[string]bool)
	g.finalizedPaths = make(map[string]bool)
	for _, line := range lines {
		n := strings.Split(line, "-")
		n0 := n[0]
		n1 := n[1]

		if _, ok := g.nodes[n0]; ok {
			if !containsString(g.nodes[n0], n1) {
				g.nodes[n0] = append(g.nodes[n0], n1)
			}
		} else {
			g.nodes[n0] = []string{n1}
		}

		if _, ok := g.nodes[n1]; ok {
			if !containsString(g.nodes[n1], n0) {
				g.nodes[n1] = append(g.nodes[n1], n0)
			}
		} else {
			g.nodes[n1] = []string{n0}
		}
	}
}

func (g *graph) visitNeighbors(node string, previouslyVisited string) {
	// for each neighbor
	for _, n := range (*g).nodes[node] {
		if n == "start" {
			continue
		} else if n == "end" {
			finPath := fmt.Sprint(previouslyVisited, n)
			if _, ok := g.finalizedPaths[finPath]; !ok {
				g.finalizedPaths[finPath] = true
			}
		} else {
			if (unicode.IsLower(rune(n[0])) && !strings.Contains(previouslyVisited, fmt.Sprint(",", n, ","))) ||
				unicode.IsUpper(rune(n[0])) {
				g.visitedPaths[fmt.Sprint(previouslyVisited, n, ",")] = true
				g.visitNeighbors(n, fmt.Sprint(previouslyVisited, n, ","))
			}
		}
	}
}

func task1(lines []string) int {
	g := graph{}
	g.load(lines)
	g.visitedPaths["start,"] = true
	g.visitNeighbors("start", "start,")
	return len(g.finalizedPaths)
}

func main() {
	inputUrl := "https://adventofcode.com/2021/day/12/input"
	input, err := internal.GetInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// be careful about the linebreak in the last number
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")
	fmt.Println("Task 1: Finalized paths:", task1(lines))
}
