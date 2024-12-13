package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type machineInfo struct {
	buttonAX int64
	buttonAY int64
	buttonBX int64
	buttonBY int64
	prizeX   int64
	prizeY   int64
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var machines []machineInfo
	for _, block := range strings.Split(string(dat), "\n\n") {
		var machine machineInfo
		for _, line := range strings.Split(strings.Trim(block, "\n "), "\n") {
			lineParts := strings.Split(line, ": ")
			coords := strings.Split(lineParts[1], ", ")
			coordX, e1 := strconv.Atoi(coords[0][2:])
			panicErr(e1)
			coordY, e2 := strconv.Atoi(coords[1][2:])
			panicErr(e2)
			if lineParts[0] == "Button A" {
				machine.buttonAX = int64(coordX)
				machine.buttonAY = int64(coordY)
			} else if lineParts[0] == "Button B" {
				machine.buttonBX = int64(coordX)
				machine.buttonBY = int64(coordY)
			} else if lineParts[0] == "Prize" {
				machine.prizeX = int64(coordX)
				machine.prizeY = int64(coordY)
			}
		}

		machines = append(machines, machine)
	}

	totalCost := int64(0)
	for i, machine := range machines {
		a, b := solveEquation(machine.buttonAX, machine.buttonAY, machine.buttonBX, machine.buttonBY, machine.prizeX, machine.prizeY)
		fmt.Printf("Machine %d - (%d, %d) + (%d, %d) = (%d, %d) has solution (%d, %d)\n", i, machine.buttonAX, machine.buttonAY, machine.buttonBX, machine.buttonBY, machine.prizeX, machine.prizeY, a, b)
		if a > 100 || b > 100 {
			fmt.Println("Skipping machine", i)
			continue
		}

		totalCost += (3 * a) + b
	}

	fmt.Println(totalCost)
}

func solveEquation(ax, ay, bx, by, px, py int64) (int64, int64) {
	// Solve system of equations:
	// ax * t + bx * s = px
	// ay * t + by * s = py

	// Using Cramer's rule: https://en.wikipedia.org/wiki/Cramer%27s_rule
	det := ax*by - ay*bx
	if det == 0 {
		return 0, 0 // No unique solution
	}

	t := (px*by - py*bx)
	s := (ax*py - ay*px)

	if t%det != 0 || s%det != 0 {
		return 0, 0 // No integer solution
	}

	return t / det, s / det
}
