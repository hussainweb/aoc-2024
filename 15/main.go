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

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	whMap, moves, rpr, rpc := parseInput(string(dat))

	for i, move := range moves {
		fmt.Print("\033[H\033[2J")
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

	switch move {
	case '^':
		newR--

	case 'v':
		newR++

	case '<':
		newC--

	case '>':
		newC++
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
	success, _, _ := moveCell(whMap, newR, newC, move)
	if success {
		whMap[newR][newC] = whMap[rpr][rpc]
		whMap[rpr][rpc] = '.'
		return true, newR, newC
	}

	return false, rpr, rpc
}

func getGpsSum(whMap [][]rune) (sum uint) {
	for r, line := range whMap {
		for c, char := range line {
			if char == 'O' {
				sum += uint(r*100 + c)
			}
		}
	}

	return
}
