package main

import (
	"fmt"
	"os"
	"strings"
)

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

type vertex struct {
	height  int
	visited bool
	marked  bool
}

type point struct {
	l, c int
}

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var topoMap [][]vertex
	var trailHeads []point
	for l, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		topoLine := make([]vertex, len(line))

		for c, height := range line {
			h := int(height - 48)

			if h == 0 {
				p := point{l, c}
				trailHeads = append(trailHeads, p)
			}

			topoLine[c] = vertex{h, false, false}
		}

		topoMap = append(topoMap, topoLine)
	}

	fmt.Println(trailHeads)
	drawMap(topoMap)

	sum := 0
	for _, p := range trailHeads {
		workMap := copyMap(topoMap)
		var peaks []point
		peaks = visitVertex(workMap, p.l, p.c, peaks)

		fmt.Println("From", p, "we found", len(peaks), "peaks")

		sum += len(peaks)
	}

	fmt.Println("Sum", sum)
}

func visitVertex(topoMap [][]vertex, l int, c int, peaks []point) []point {
	topoMap[l][c].visited = true
	// fmt.Println("Visiting", l, c, "with height", topoMap[l][c].height)
	for y := l - 1; y <= l+1; y++ {
		if y < 0 || y >= len(topoMap) {
			continue
		}

		for x := c - 1; x <= c+1; x++ {
			if x < 0 || x >= len(topoMap[0]) {
				continue
			}

			if !(x == c || y == l) || (x == c && y == l) {
				// Can't visit self or diagonals.
				continue
			}

			v := &topoMap[y][x]
			if v.height == topoMap[l][c].height+1 {
				if v.visited {
					// We have already been to this point.
					continue
				}

				if v.height == 9 && !v.marked {
					peaks = append(peaks, point{x, y})
					v.marked = true
					v.visited = true
					continue
				}

				peaks = visitVertex(topoMap, y, x, peaks)
			}
		}
	}
	return peaks
}

func copyMap(topoMap [][]vertex) [][]vertex {
	var newTopoMap [][]vertex
	for _, line := range topoMap {
		newTopoMap = append(newTopoMap, append([]vertex{}, line...))
	}
	return newTopoMap
}

func drawMap(topoMap [][]vertex) {
	for _, l := range topoMap {
		for _, v := range l {
			fmt.Print(v.height)
		}
		fmt.Println()
	}
}
