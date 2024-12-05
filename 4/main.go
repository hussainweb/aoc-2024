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
			if char != 'X' {
				continue
			}

			enoughSpaceRight := c < totalCols-3
			enoughSpaceLeft := c > 2
			enoughSpaceUp := r > 2
			enoughSpaceDown := r < totalLines-3

			// Search in each direction.
			// First, go right.
			if enoughSpaceRight {
				if line[c+1] == 'M' && line[c+2] == 'A' && line[c+3] == 'S' {
					count++
				}
			}
			// Next, go left.
			if enoughSpaceLeft {
				if line[c-1] == 'M' && line[c-2] == 'A' && line[c-3] == 'S' {
					count++
				}
			}
			// Go top.
			if enoughSpaceUp {
				if lines[r-1][c] == 'M' && lines[r-2][c] == 'A' && lines[r-3][c] == 'S' {
					count++
				}
			}
			// Go bottom.
			if enoughSpaceDown {
				if lines[r+1][c] == 'M' && lines[r+2][c] == 'A' && lines[r+3][c] == 'S' {
					count++
				}
			}

			// Now diagonals.
			// Top right.
			if enoughSpaceUp && enoughSpaceRight {
				if lines[r-1][c+1] == 'M' && lines[r-2][c+2] == 'A' && lines[r-3][c+3] == 'S' {
					count++
				}
			}
			// Top left.
			if enoughSpaceUp && enoughSpaceLeft {
				if lines[r-1][c-1] == 'M' && lines[r-2][c-2] == 'A' && lines[r-3][c-3] == 'S' {
					count++
				}
			}
			// Bottom right.
			if enoughSpaceDown && enoughSpaceRight {
				if lines[r+1][c+1] == 'M' && lines[r+2][c+2] == 'A' && lines[r+3][c+3] == 'S' {
					count++
				}
			}
			// Bottom left.
			if enoughSpaceDown && enoughSpaceLeft {
				if lines[r+1][c-1] == 'M' && lines[r+2][c-2] == 'A' && lines[r+3][c-3] == 'S' {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
