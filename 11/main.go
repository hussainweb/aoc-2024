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

	var stones []uint64
	for _, stoneStr := range strings.Split(strings.Trim(string(dat), "\n"), " ") {
		i, e := strconv.ParseUint(stoneStr, 10, 64)
		panicErr(e)
		stones = append(stones, i)
	}

	fmt.Println(stones)
	for i := 0; i < 25; i++ {
		stones = blink(stones)
		fmt.Println("Iteration", i, "Number of stones", len(stones))
		// fmt.Println(stones)
	}

	fmt.Println(len(stones))
}

func blink(stones []uint64) []uint64 {
	newStones := make([]uint64, 0, len(stones)*2)
	for _, s := range stones {
		if s == 0 {
			newStones = append(newStones, 1)
			continue
		}

		ss := strconv.FormatUint(s, 10)
		l := len(ss)
		if l%2 == 0 {
			i1, e1 := strconv.ParseUint(ss[0:l/2], 10, 64)
			panicErr(e1)
			i2, e2 := strconv.ParseUint(ss[l/2:l], 10, 64)
			panicErr(e2)
			newStones = append(newStones, i1, i2)
			continue
		}

		newStones = append(newStones, s*2024)
	}

	return newStones
}
