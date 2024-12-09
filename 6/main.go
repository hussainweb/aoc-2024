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

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var guardMap [][]rune
	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		guardMap = append(guardMap, []rune(line))
	}

	count := walkGuard(guardMap)
	fmt.Println(count)
	drawMap(guardMap)
}

func walkGuard(guardMap [][]rune) int {
	x := -1
	y := -1
	xInc := 0
	yInc := 0
	count := 0

	var line []rune
	var char rune
wgfunc:
	for y, line = range guardMap {
		for x, char = range line {
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				switch char {
				case '^':
					yInc = -1
				case '>':
					xInc = 1
				case '<':
					xInc = -1
				case 'v':
					yInc = 1
				}
				break wgfunc
			}
		}
	}

	if x == -1 || y == -1 {
		panic("No starting position found")
	}

	guardMap[y][x] = 'X'
	count++

	for {
		x += xInc
		y += yInc
		if x < 0 || y < 0 || y >= len(guardMap) || x >= len(guardMap[y]) {
			break
		}

		if guardMap[y][x] == 'X' {
			continue
		}

		if guardMap[y][x] != '#' {
			guardMap[y][x] = 'X'
			count++
			continue
		}

		x -= xInc
		y -= yInc

		// Rotate the direction
		// (0,-1) => (1,0) => (0,1) => (-1,0)
		if xInc == 0 && yInc == -1 {
			xInc = 1
			yInc = 0
		} else if xInc == 1 && yInc == 0 {
			xInc = 0
			yInc = 1
		} else if xInc == 0 && yInc == 1 {
			xInc = -1
			yInc = 0
		} else if xInc == -1 && yInc == 0 {
			xInc = 0
			yInc = -1
		} else {
			panic("Invalid direction")
		}
	}

	return count
}

func drawMap(guardMap [][]rune) {
	for _, line := range guardMap {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}
