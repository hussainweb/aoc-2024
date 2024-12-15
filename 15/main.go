package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(str string) (whMap [][]rune, moves []rune, rpr, rpc int) {
	parts := strings.Split(str, "\n\n")
	for r, line := range strings.Split(parts[0], "\n") {
		whMap = append(whMap, []rune(line))
		for c, char := range line {
			if char == '@' {
				rpr, rpc = r, c
			}
		}
	}

	moves = []rune(strings.ReplaceAll(parts[1], "\n", ""))
	return
}

func resizeMap(whMap [][]rune) [][]rune {
	var newWhMap [][]rune

	for _, line := range whMap {
		mapLine := make([]rune, len(line)*2)
		for c, char := range line {
			newC := c * 2
			switch char {
			case '#':
				mapLine[newC] = '#'
				mapLine[newC+1] = '#'

			case '@':
				mapLine[newC] = '@'
				mapLine[newC+1] = '.'

			case 'O':
				mapLine[newC] = '['
				mapLine[newC+1] = ']'

			case '.':
				mapLine[newC] = '.'
				mapLine[newC+1] = '.'
			}
		}
		newWhMap = append(newWhMap, mapLine)
	}
	return newWhMap
}

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	whMap, moves, rpr, rpc := parseInput(string(dat))

	whMap = resizeMap(whMap)
	// The robot is at a different position now.
	rpc *= 2

	fmt.Print("\033[H\033[2J")
	for i, move := range moves {
		fmt.Printf("\033[0;0H")
		fmt.Printf("Move %d/%d: %c\n\n", i+1, len(moves), move)

		_, rpr, rpc = moveCell(whMap, rpr, rpc, move)
		drawMap(whMap)
		time.Sleep(20 * time.Millisecond)
	}
	sum := getGpsSum(whMap)
	fmt.Println("Sum", sum)
}

func drawMap(whMap [][]rune) {
	for _, line := range whMap {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func moveCell(whMap [][]rune, rpr, rpc int, move rune) (bool, int, int) {
	newR := rpr
	newC := rpc

	rOff := 0
	cOff := 0
	switch move {
	case '^':
		newR--
		rOff = -1

	case 'v':
		newR++
		rOff = 1

	case '<':
		newC--
		cOff = -1

	case '>':
		newC++
		cOff = 1
	}

	if newR < 0 || newR >= len(whMap) || newC < 0 || newC >= len(whMap[rpr]) {
		return false, rpr, rpc
	}

	// Have we hit a wall?
	if whMap[newR][newC] == '#' {
		return false, rpr, rpc
	}

	// Empty space? Just move ahead.
	if whMap[newR][newC] == '.' {
		whMap[newR][newC] = whMap[rpr][rpc]
		whMap[rpr][rpc] = '.'
		return true, newR, newC
	}

	// Box? Try to move it.
	if whMap[newR][newC] == '[' || whMap[newR][newC] == ']' {
		if newR == rpr {
			// If we are moving in the same row, we can move the box.
			success, _, _ := moveCell(whMap, newR+rOff, newC+cOff, move)
			if success {
				// If it was a success, there is an empty space now and we can safely move.
				moveCell(whMap, newR, newC, move)
				whMap[newR][newC] = whMap[rpr][rpc]
				whMap[rpr][rpc] = '.'
				return true, newR, newC
			}
		} else {
			// Moving up or down is much more difficult.
			// Work on a copy of the map, so we don't mess up the original.
			whMapCopy := make([][]rune, len(whMap))
			for r, line := range whMap {
				whMapCopy[r] = make([]rune, len(line))
				copy(whMapCopy[r], line)
			}

			otherC := newC + 1
			if whMap[newR][newC] == ']' {
				otherC = newC - 1
			}
			success1, _, _ := moveCell(whMapCopy, newR, newC, move)
			success2, _, _ := moveCell(whMapCopy, newR, otherC, move)
			if success1 && success2 {
				// If it was a success, our copy is successful. Copy it to the main map and
				// there is an empty space now and we can safely move.
				for r, line := range whMapCopy {
					copy(whMap[r], line)
				}
				whMap[newR][newC] = whMap[rpr][rpc]
				whMap[rpr][rpc] = '.'
				return true, newR, newC
			}
		}
	}

	return false, rpr, rpc
}

func getGpsSum(whMap [][]rune) (sum uint) {
	for r, line := range whMap {
		for c, char := range line {
			if char == '[' {
				sum += uint(r*100 + c)
			}
		}
	}

	return
}
