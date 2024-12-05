package main

import (
	"fmt"
	"os"
	"strconv"
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

	safeReports := 0
	for r, line := range strings.Split(string(dat), "\n") {
		if len(line) == 0 {
			continue
		}

		var levels []int
		for _, elem := range strings.Split(line, " ") {
			e, err := strconv.Atoi(elem)
			panicErr(err)

			levels = append(levels, e)
		}
		fmt.Println(r, levels)

		if isSafeReport(levels, true) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isSafeReport(levels []int, tryRemoving bool) bool {
	incrementing := false
	decrementing := false

	for i, e := range levels {
		if i > 0 {
			if !incrementing && !decrementing {
				if e > levels[i-1] {
					incrementing = true
				} else if e < levels[i-1] {
					decrementing = true
				}
			}

			// We now know the sequence. Let's check if the next element
			// is in the right order. If not, the report is not safe.
			diff := e - levels[i-1]
			if decrementing {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				if !tryRemoving {
					return false
				}

				// Try removing each element and see if the report is safe
				for j := 0; j < len(levels); j++ {
					removed := removeElement(levels, j)
					if isSafeReport(removed, false) {
						return true
					}
				}

				// Removals didn't work. Give up.
				return false
			}
		}
	}

	return true
}

func removeElement(levels []int, index int) []int {
	var copyLevels []int = make([]int, len(levels))
	copy(copyLevels, levels)
	return append(copyLevels[:index], copyLevels[index+1:]...)
}
