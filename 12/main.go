package main

import (
	"fmt"
	"os"
	"strings"
)

type plot struct {
	crop       rune
	borderT    bool
	borderL    bool
	borderB    bool
	borderR    bool
	numBorders int
	region     int
}

type regionInfo struct {
	r      int
	c      int
	number int

	crop    rune
	size    int
	borders int
}

func newPlot(crop rune) *plot {
	p := plot{crop: crop}
	return &p
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	panicErr(err)

	var farmMap [][]plot
	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		farmLine := make([]plot, 0, len(line))
		for _, char := range line {
			farmLine = append(farmLine, *newPlot(char))
		}

		farmMap = append(farmMap, farmLine)
	}

	processBordersMap(farmMap)
	regions := fillRegions(farmMap)
	sum := 0
	for _, ri := range regions {
		price := ri.size * ri.borders
		fmt.Printf("Region %d with plant %c has a price of %d x %d = %d\n", ri.number, ri.crop, ri.size, ri.borders, price)
		sum += price
	}
	fmt.Println("Total price", sum)
	// drawMap(farmMap, true)
}

func processBordersMap(farmMap [][]plot) {
	for r, line := range farmMap {
		for c := range line {
			p := &farmMap[r][c]

			// Determine borders based on the neighbors (or edges).
			// We need to check top, right, bottom, and left. No diagonals.
			// (r-1,c), (r,c+1), (r+1,c), (r,c-1).
			if r == 0 || (farmMap[r-1][c].crop != p.crop) {
				p.borderT = true
				p.numBorders++
			}
			if c == len(farmMap[0])-1 || farmMap[r][c+1].crop != p.crop {
				p.borderR = true
				p.numBorders++
			}
			if r == len(farmMap)-1 || farmMap[r+1][c].crop != p.crop {
				p.borderB = true
				p.numBorders++
			}
			if c == 0 || farmMap[r][c-1].crop != p.crop {
				p.borderL = true
				p.numBorders++
			}
		}
	}

	// Now loop through the map again. This time, check in straight lines to see consecutive
	// borders and remove them.
	// This means, going row-wise, remove all consecutive top borders until we hit a plot
	// with no top border. Do this for all borders one at a time.
	for r, line := range farmMap {
		for c := range line {
			p := &farmMap[r][c]

			if p.borderT {
				for c2 := c + 1; c2 < len(line); c2++ {
					if !farmMap[r][c2].borderT || farmMap[r][c2].crop != p.crop {
						break
					}
					farmMap[r][c2].borderT = false
					farmMap[r][c2].numBorders--
				}
			}
			if p.borderB {
				for c2 := c + 1; c2 < len(line); c2++ {
					if !farmMap[r][c2].borderB || farmMap[r][c2].crop != p.crop {
						break
					}
					farmMap[r][c2].borderB = false
					farmMap[r][c2].numBorders--
				}
			}
			if p.borderL {
				for r2 := r + 1; r2 < len(farmMap); r2++ {
					if !farmMap[r2][c].borderL || farmMap[r2][c].crop != p.crop {
						break
					}
					farmMap[r2][c].borderL = false
					farmMap[r2][c].numBorders--
				}
			}
			if p.borderR {
				for r2 := r + 1; r2 < len(farmMap); r2++ {
					if !farmMap[r2][c].borderR || farmMap[r2][c].crop != p.crop {
						break
					}
					farmMap[r2][c].borderR = false
					farmMap[r2][c].numBorders--
				}
			}
		}
	}
}

func fillRegions(farmMap [][]plot) []regionInfo {
	currentRegion := 1
	var regions []regionInfo

	for r, line := range farmMap {
		for c := range line {
			p := &farmMap[r][c]

			if p.region != 0 {
				continue
			}

			ri := regionInfo{r: r, c: c, number: currentRegion, crop: p.crop}

			fillSurroundingRegions(farmMap, r, c, p.crop, &ri)
			regions = append(regions, ri)
			// regionCount[currentRegion] = count
			// regionBorders[currentRegion] = borders
			currentRegion++
		}
	}

	return regions
}

func fillSurroundingRegions(farmMap [][]plot, r, c int, crop rune, ri *regionInfo) *regionInfo {
	if r < 0 || c < 0 || r >= len(farmMap) || c >= len(farmMap[0]) {
		return ri
	}

	p := &farmMap[r][c]
	if p.region != 0 || p.crop != crop {
		return ri
	}

	ri.size++
	ri.borders += p.numBorders
	p.region = ri.number

	fillSurroundingRegions(farmMap, r-1, c, crop, ri)
	fillSurroundingRegions(farmMap, r, c+1, crop, ri)
	fillSurroundingRegions(farmMap, r+1, c, crop, ri)
	fillSurroundingRegions(farmMap, r, c-1, crop, ri)
	return ri
}

func drawMap(topoMap [][]plot, borders bool) {
	for _, l := range topoMap {
		if borders {
			for _, v := range l {
				fmt.Print(" ")
				if v.borderT {
					fmt.Print("-")
				} else {
					fmt.Print(" ")
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
		for _, v := range l {
			if borders {
				if v.borderL {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print(string(v.crop))
			// fmt.Print(v.region)
			if borders {
				if v.borderR {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
		if borders {
			for _, v := range l {
				fmt.Print(" ")
				if v.borderB {
					fmt.Print("-")
				} else {
					fmt.Print(" ")
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}
