package main

import (
	"fmt"
	"math"
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

	sum := uint64(0)
	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		equation := strings.Split(line, ": ")
		lhs, err := strconv.ParseUint(equation[0], 10, 64)
		panicErr(err)

		var eqRhs []int
		for _, page := range strings.Split(equation[1], " ") {
			num, err := strconv.Atoi(page)
			panicErr(err)

			eqRhs = append(eqRhs, num)
		}

		fmt.Println(lhs, "=", eqRhs)
		found := tryPermutations(lhs, eqRhs)
		if found {
			fmt.Println("Found", lhs)
			sum += lhs
		}
	}

	fmt.Println("Sum:", sum)
}

func tryPermutations(lhs uint64, rhs []int) bool {
	// We have 2^(count-1) permutations possible. Create an array to hold all possible results.
	places := pow(3, len(rhs)-1)
	var permutations []uint64 = make([]uint64, places)

	// Fill our array with the first number in our list. Subsequent numbers will be added or
	// multiplied to generate all permutations, like so. For example, a list of 4 numbers will have
	// 0 x (0 1 2 3 4 5 6 7) 2^0
	// 1 * (0 1 2 3) + (4 5 6 7) length of operands: 2^2
	// 2 * (0 1) + (2 3) * (4 5) + (6 7) length of operands: 2^1
	// 3 * (0) + (1) * (2) + (3) * (4) + (5) * (6) + (7) length of operands: 2^0
	// --
	// For part 2, we have 3 operators, so we change the power base. All the above calculations
	// now become 3^something.
	for x := range places {
		permutations[x] = uint64(rhs[0])
	}

	operator := 0 // Mutiply
	for i := 1; i < len(rhs); i++ {
		j := 0
		// Iterate over each permutation
		for j < places {
			// Determine the length of each operand list and either multiply or add.
			ln := pow(3, len(rhs)-i-1)
			for k := 0; k < ln; k++ {
				switch operator {
				case 0:
					permutations[j+k] *= uint64(rhs[i])
				case 1:
					permutations[j+k] += uint64(rhs[i])
				case 2:
					o1 := strconv.FormatUint(permutations[j+k], 10)
					o2 := strconv.Itoa(rhs[i])
					permutations[j+k], _ = strconv.ParseUint(o1+o2, 10, 64)
				}
			}
			// Advance to the next operand list and switch the operation.
			j += ln
			operator++
			if operator > 2 {
				operator = 0
			}
		}
	}

	for _, res := range permutations {
		if res == lhs {
			return true
		}
	}

	return false
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
