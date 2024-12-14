package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type robotState struct {
	px int
	py int
	vx int
	vy int
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var robots []robotState
	coordRE := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		num := coordRE.FindStringSubmatch(line)
		px, e1 := strconv.Atoi(num[1])
		panicErr(e1)
		py, e2 := strconv.Atoi(num[2])
		panicErr(e2)
		vx, e3 := strconv.Atoi(num[3])
		panicErr(e3)
		vy, e4 := strconv.Atoi(num[4])
		panicErr(e4)
		robots = append(robots, robotState{px, py, vx, vy})
	}

	maxX := 101
	maxY := 103

	seconds := 100

	for s := 0; s < seconds; s++ {
		for r := range robots {
			robots[r] = IterateRobot(robots[r], maxX, maxY)
		}
	}

	quadrantRobots := make([]int, 5)
	for i, r := range robots {
		quadrant := checkQuadrant(r, maxX, maxY)
		fmt.Printf("Robot %d is in quadrant %d\n", i, quadrant)
		quadrantRobots[quadrant] += 1
	}
	fmt.Println(quadrantRobots)
	fmt.Println("Score", quadrantRobots[1]*quadrantRobots[2]*quadrantRobots[3]*quadrantRobots[4])
}

func IterateRobot(robot robotState, maxX, maxY int) robotState {
	newX := (robot.px + robot.vx)
	if newX < 0 {
		newX += maxX
	}
	newX = newX % maxX
	newY := (robot.py + robot.vy)
	if newY < 0 {
		newY += maxY
	}
	newY = newY % maxY
	return robotState{newX, newY, robot.vx, robot.vy}
}

func checkQuadrant(robot robotState, maxX, maxY int) int {
	midX := maxX / 2
	midY := maxY / 2
	if robot.px == midX || robot.py == midY {
		return 0
	}

	if robot.px < midX && robot.py < midY {
		return 1
	}
	if robot.px < midX && robot.py > midY {
		return 2
	}
	if robot.px > midX && robot.py < midY {
		return 3
	}
	if robot.px > midX && robot.py > midY {
		return 4
	}

	return 0
}
