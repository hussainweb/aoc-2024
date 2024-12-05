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
	for _, line := range strings.Split(string(dat), "\n") {
		if len(line) == 0 {
			continue
		}

		var levels []int
		for _, elem := range strings.Split(line, " ") {
			e, err := strconv.Atoi(elem)
			panicErr(err)

			levels = append(levels, e)
		}

		if isSafeReport(levels) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isSafeReport(levels []int) bool {
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
				return false
			}
		}
	}

	return true
}
