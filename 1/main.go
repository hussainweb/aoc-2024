package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Read file input.txt
	dat, err := os.ReadFile("input.txt")
	panicErr(err)
	lines := strings.Split(string(dat), "\n")

	var list1 []int
	var list2 []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		val := strings.Split(line, "   ")
		if len(val) != 2 {
			panic("Invalid line")
		}

		p1, e1 := strconv.Atoi(val[0])
		panicErr(e1)
		list1 = append(list1, p1)
		p2, e2 := strconv.Atoi(val[1])
		panicErr(e2)
		list2 = append(list2, p2)

		fmt.Println(line)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i, l1 := range list1 {
		diff := l1 - list2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	fmt.Println("Sum: ", sum)
}
