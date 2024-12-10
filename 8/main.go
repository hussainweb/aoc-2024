package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	l, c int
}

func newPoint(l, c int) *Point {
	p := Point{l, c}
	return &p
}

// This is a simple diff. Mainly a helper.
func diffPoints(a, b Point) *Point {
	p := Point{a.l - b.l, a.c - b.c}
	return &p
}

// This determines the mirror point based on the distance between
// 2 given points.
func mirrorDiffPoints(a, b Point) *Point {
	dp := diffPoints(a, b)
	return newPoint(b.l-dp.l, b.c-dp.c)
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var antennaMap [][]rune
	var antennaLoc map[rune][]Point = make(map[rune][]Point)
	for l, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		antennaMap = append(antennaMap, []rune(line))

		for c, char := range line {
			if char == '.' {
				continue
			}

			arr, okay := antennaLoc[char]
			if !okay {
				arr = make([]Point, 0)
			}
			arr = append(arr, *newPoint(l, c))
			antennaLoc[char] = arr
		}
	}

	drawMap(antennaMap)
	for r, points := range antennaLoc {
		fmt.Println("Antenna", string(r), "Points", points)
	}

	// Figure out antinodes
	totalRows := len(antennaMap)
	totalCols := len(antennaMap[0])
	var runeAntinodes map[rune][]Point = make(map[rune][]Point)
	var listAntinodes []Point = make([]Point, 0)
	countAntinodes := 0
	for r, points := range antennaLoc {
		for i, p := range points {
			for j := 0; j < len(points); j++ {
				// Go over each point pair except pairing the point with itself.
				if i == j {
					continue
				}

				dp := mirrorDiffPoints(p, points[j])
				if dp.l < 0 || dp.c < 0 || dp.l >= totalRows || dp.c >= totalCols {
					continue
				}

				// We don't need this map at all, but maybe we need it in part 2.
				runeAntinodes[r] = append(runeAntinodes[r], *newPoint(dp.l, dp.c))

				// Add to listAntinodes
				found := false
				for _, p := range listAntinodes {
					if p.l == dp.l && p.c == dp.c {
						found = true
					}
				}
				if !found {
					listAntinodes = append(listAntinodes, *dp)
					countAntinodes++
				}
			}
		}
	}

	// Update map from runeAntinodes
	for _, points := range runeAntinodes {
		for _, p := range points {
			antennaMap[p.l][p.c] = '#'
		}
	}

	drawMap(antennaMap)
	fmt.Println(countAntinodes)
}

func drawMap(runeMap [][]rune) {
	for _, line := range runeMap {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}

func copyMap(runeMap [][]rune) [][]rune {
	var newRuneMap [][]rune
	for _, line := range runeMap {
		newRuneMap = append(newRuneMap, append([]rune{}, line...))
	}
	return newRuneMap
}
