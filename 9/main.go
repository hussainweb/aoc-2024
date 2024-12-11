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
	len, compactLen := determineLength(diskMap)

	diskLayout := make([]int, len)
	compactLayout := make([]int, compactLen)

	fillDiskLayout(diskMap, diskLayout)

	fillAndCompactLayout(diskLayout, compactLayout)

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
	i := 0
	j := len(diskLayout) - 1
	for ; i < len(compactLayout); i++ {
		compactLayout[i] = diskLayout[i]
		if diskLayout[i] == -1 {
			for diskLayout[j] == -1 {
				j--
			}
			compactLayout[i] = diskLayout[j]
			diskLayout[j] = -2
			j--
		}
	}
}

func calcChecksum(layout []int) uint64 {
	sum := uint64(0)
	for i, id := range layout {
		sum += uint64(i * id)
	}
	return sum
}
