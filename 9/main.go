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

	diskMap := strings.Trim(string(dat), "\n")
	len, _ := determineLength(diskMap)

	diskLayout := make([]int, len)
	compactLayout := make([]int, len)

	fillDiskLayout(diskMap, diskLayout)
	// fmt.Println(diskLayout)

	fillAndCompactLayout(diskLayout, compactLayout)

	// fmt.Println(compactLayout)
	fmt.Println(calcChecksum(compactLayout))
}

func determineLength(diskMap string) (int, int) {
	sum := 0
	compactSum := 0
	for i, r := range diskMap {
		sum += int(r) - 48
		if i%2 == 0 {
			compactSum += int(r) - 48
		}
	}
	return sum, compactSum
}

func fillDiskLayout(diskMap string, diskLayout []int) {
	j := 0
	fileId := 0
	for i, r := range diskMap {
		span := int(r) - 48
		for k := 0; k < span; k++ {
			if i%2 == 0 {
				diskLayout[j+k] = fileId
			} else {
				diskLayout[j+k] = -1
			}
		}
		if i%2 == 0 {
			fileId++
		}
		j += span
	}
}

func fillAndCompactLayout(diskLayout []int, compactLayout []int) {
	copy(compactLayout, diskLayout)

	j := len(compactLayout) - 1
	for j > 0 {
		if compactLayout[j] < 0 {
			j--
			continue
		}

		// Found a file. Figure out how big it is.
		sF, eF := findSpan(compactLayout, j)
		fileLen := eF - sF + 1
		// fmt.Println("For file", compactLayout[j], "at", sF, eF)

		// Now find a space for this file starting at the left.
		i := 0
		for ; i < sF; i++ {
			if compactLayout[i] >= 0 {
				// We are only concerned with finding a space here. Skip the file.
				continue
			}

			// Determine the span of this empty space.
			sS, eS := findSpan(compactLayout, i)
			spaceLen := eS - sS + 1
			// fmt.Println("i=", i, "j=", j, "space=", sS, eS, "file=", sF, eF)

			if spaceLen >= fileLen {
				// First, copy the file to the empty space
				for k := sS; k < sS+fileLen; k++ {
					compactLayout[k] = compactLayout[j]
				}
				// Next, clear the previous file space.
				for k := sF; k <= eF; k++ {
					compactLayout[k] = -2
				}
				break
			}
			i = eS
		}

		j = sF - 1
	}
}

func calcChecksum(layout []int) uint64 {
	sum := uint64(0)
	for i, id := range layout {
		if id < 0 {
			continue
		}
		sum += uint64(i * id)
	}
	return sum
}

func findSpan(layout []int, idx int) (int, int) {
	s := idx
	e := idx

	elem := layout[s]
	for layout[s] == elem {
		s--
		if s < 0 {
			break
		}
	}
	for layout[e] == elem {
		e++
		if e >= len(layout) {
			break
		}
	}

	// Correct the indexes
	s++
	e--

	return s, e
}
