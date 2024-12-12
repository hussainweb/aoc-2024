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

var stoneCache map[string]int = make(map[string]int)

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var stones []uint64
	for _, stoneStr := range strings.Split(strings.Trim(string(dat), "\n "), " ") {
		i, e := strconv.ParseUint(stoneStr, 10, 64)
		panicErr(e)
		stones = append(stones, i)
	}

	fmt.Println(stones)

	sum := 0
	for _, s := range stones {
		sum += blinkRecursive(s, 75)
	}

	fmt.Println(sum)
}

func blinkRecursive(stone uint64, depth int) int {
	key := strconv.FormatUint(stone, 10) + "-" + strconv.Itoa(depth)
	val, okay := stoneCache[key]
	if okay {
		return val
	}

	if depth == 0 {
		return 1
	}

	if stone == 0 {
		stoneCache[key] = blinkRecursive(1, depth-1)
		return stoneCache[key]
	}

	ss := strconv.FormatUint(stone, 10)
	l := len(ss)
	if l%2 == 0 {
		i1, e1 := strconv.ParseUint(ss[0:l/2], 10, 64)
		panicErr(e1)
		i2, e2 := strconv.ParseUint(ss[l/2:l], 10, 64)
		panicErr(e2)
		stoneCache[key] = blinkRecursive(i1, depth-1) + blinkRecursive(i2, depth-1)
		return stoneCache[key]
	}

	stoneCache[key] = blinkRecursive(stone*2024, depth-1)
	return stoneCache[key]
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
