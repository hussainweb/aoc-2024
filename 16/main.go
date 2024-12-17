package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(input string) (maze [][]rune, startPos point, endPos point) {
	for y, line := range strings.Split(input, "\n") {
		maze = append(maze, []rune(line))

		for x, char := range line {
			switch char {
			case 'S':
				startPos = point{x, y}
			case 'E':
				endPos = point{x, y}
			}
		}
	}

	return
}

func navigateMaze(maze [][]rune, startPos point, endPos point, direction int, initScore int, blockedMaze [][]rune, mazeCallback func([][]rune)) (score int) {
	if startPos.x == endPos.x && startPos.y == endPos.y {
		return initScore
	}

	maze[startPos.y][startPos.x] = '*'
	score = initScore + 1

	directions := getTraversableDirections(maze, blockedMaze, startPos, direction)
	if len(directions) == 0 {
		return -1
	}

	lastScore := -1
	for _, nextDir := range directions {
		nextPos := getNextPos(startPos, nextDir)
		wMaze := maze
		if len(directions) > 1 {
			wMaze = copyMaze(maze)
		}
		mazeCallback(wMaze)

		turnScore := 0
		if nextDir != direction {
			turnScore = 1000
		}

		dirScore := navigateMaze(wMaze, nextPos, endPos, nextDir, score+turnScore, blockedMaze, mazeCallback)
		if dirScore == -1 {
			// If this path can't take us to end, then let's block it for real.
			// TODO: Since a path can be blocked in multiple directions, use a bitmap rather than a rune.
			blockedMaze[nextPos.y][nextPos.x] = rune(nextDir)
			continue
		}

		if lastScore == -1 || dirScore < lastScore {
			lastScore = dirScore
		}
	}

	score = lastScore

	return
}

func getNextPos(pos point, direction int) point {
	switch direction {
	case 0:
		return point{pos.x, pos.y - 1}
	case 1:
		return point{pos.x + 1, pos.y}
	case 2:
		return point{pos.x, pos.y + 1}
	case 3:
		return point{pos.x - 1, pos.y}
	}

	panic("Invalid direction")
}

func canTraverse(maze [][]rune, blockedMaze [][]rune, pos point, direction int) (bool, point) {
	nextPos := getNextPos(pos, direction)

	// Since our maze is bordered with walls, we don't really need this for now.
	// if nextPos.x < 0 || nextPos.x >= len(maze[0]) || nextPos.y < 0 || nextPos.y >= len(maze) {
	// 	return false
	// }

	return maze[nextPos.y][nextPos.x] == '.' && blockedMaze[nextPos.y][nextPos.x] != rune(direction), nextPos
}

func getTraversableDirections(maze [][]rune, blockedMaze [][]rune, pos point, direction int) (directions []int) {
	possibleDirections := []int{direction, (direction + 1) % 4, (direction + 3) % 4}
	for _, i := range possibleDirections {
		traversable, _ := canTraverse(maze, blockedMaze, pos, i)
		if traversable {
			directions = append(directions, i)
		}
	}

	return
}

func drawMaze(maze [][]rune) {
	for _, line := range maze {
		for _, char := range line {
			fmt.Print(string(char))
		}
		println()
	}
}

func copyMaze(maze [][]rune) (copy [][]rune) {
	for _, line := range maze {
		copy = append(copy, append([]rune{}, line...))
	}

	return
}

func main() {
	input, err := os.ReadFile("input.txt")
	panicErr(err)

	maze, startPos, endPos := parseInput(string(input))

	maze[endPos.y][endPos.x] = '.'

	blockedMaze := copyMaze(maze)

	fmt.Print("\033[H\033[2J")
	drawMaze(maze)

	score := navigateMaze(maze, startPos, endPos, 1, 0, blockedMaze, func(maze [][]rune) {
		fmt.Printf("\033[0;0H")
		drawMaze(maze)
		time.Sleep(1 * time.Millisecond)
	})
	fmt.Println("Score", score, "                      ")
}
