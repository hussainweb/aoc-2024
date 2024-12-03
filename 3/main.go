package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var mulInstructions = regexp.MustCompile(`(mul\(([\d]{1,3}),([\d]{1,3})\)|do(n\'t)?\(\))`)

	// Read file input.txt
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var instructions = mulInstructions.FindAllStringSubmatch(string(dat), -1)
	if instructions == nil {
		panic("No strings found")
	}

	var sum = 0
	var add = true
	for _, inst := range instructions {
		fmt.Println(inst, add)
		if inst[0] == "do()" {
			add = true
			continue
		}
		if inst[1] == "don't()" {
			add = false
			continue
		}

		if !add {
			continue
		}

		p1, e1 := strconv.Atoi(inst[2])
		panicErr(e1)
		p2, e2 := strconv.Atoi(inst[3])
		panicErr(e2)
		sum += p1 * p2
	}

	fmt.Println("Sum: ", sum)
}
