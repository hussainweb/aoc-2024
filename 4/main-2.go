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

	var lines [][]rune
	// Split the file into lines
	for _, line := range strings.Split(string(dat), "\n") {
		if len(line) == 0 {
			continue
		}
		lines = append(lines, []rune(line))
	}

	totalLines := len(lines)
	totalCols := len(lines[0])

	count := 0

	for r, line := range lines {
		for c, char := range line {
			if char != 'M' {
				continue
			}

			enoughSpaceRight := c < totalCols-2
			enoughSpaceLeft := c > 1
			enoughSpaceUp := r > 1
			enoughSpaceDown := r < totalLines-2

			// Search diagonals.
			// Top right.
			if enoughSpaceUp && enoughSpaceRight {
				if (lines[r-1][c+1] == 'A' && lines[r-2][c+2] == 'S') && ((lines[r][c+2] == 'M' && lines[r-2][c] == 'S') || (lines[r][c+2] == 'S' && lines[r-2][c] == 'M')) {
					count++
				}
			}
			// Top left.
			if enoughSpaceUp && enoughSpaceLeft {
				if lines[r-1][c-1] == 'A' && lines[r-2][c-2] == 'S' && ((lines[r][c-2] == 'M' && lines[r-2][c] == 'S') || (lines[r][c-2] == 'S' && lines[r-2][c] == 'M')) {
					count++
				}
			}
			// Bottom right.
			if enoughSpaceDown && enoughSpaceRight {
				if lines[r+1][c+1] == 'A' && lines[r+2][c+2] == 'S' && ((lines[r][c+2] == 'M' && lines[r+2][c] == 'S') || (lines[r][c+2] == 'S' && lines[r+2][c] == 'M')) {
					count++
				}
			}
			// Bottom left.
			if enoughSpaceDown && enoughSpaceLeft {
				if lines[r+1][c-1] == 'A' && lines[r+2][c-2] == 'S' && ((lines[r][c-2] == 'M' && lines[r+2][c] == 'S') || (lines[r][c-2] == 'S' && lines[r+2][c] == 'M')) {
					count++
				}
			}
		}
	}

	fmt.Println(count >> 1)
}
