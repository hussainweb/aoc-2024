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

	rules := make(map[int][]int)
	var pageLines [][]int

	for _, line := range strings.Split(string(dat), "\n") {
		if line == "" {
			continue
		}

		if strings.ContainsRune(line, '|') {
			rule := strings.Split(line, "|")
			r0, err := strconv.Atoi(rule[0])
			panicErr(err)
			r1, err := strconv.Atoi(rule[1])
			panicErr(err)

			rules[r0] = append(rules[r0], r1)
		} else {
			var pageLine []int
			for _, page := range strings.Split(line, ",") {
				num, err := strconv.Atoi(page)
				panicErr(err)

				pageLine = append(pageLine, num)
			}
			pageLines = append(pageLines, pageLine)
		}
	}

	sum := 0
	for _, pages := range pageLines {
		if !isValidPageLine(pages, rules) {
			fixedPages := fixOrder(pages, rules)
			sum += fixedPages[len(fixedPages)/2]
		}
	}

	// fmt.Println(rules)
	fmt.Println(sum)
}

func isValidPageLine(pages []int, rules map[int][]int) bool {
	for i, page := range pages {
		for _, followingPage := range rules[page] {
			// Check all the elements before this one in the pages array and if
			// any of them are equal to the following page, then the page line
			// is invalid.
			for j := 0; j < i; j++ {
				if pages[j] == followingPage {
					return false
				}
			}
		}
	}
	return true
}

func fixOrder(pages []int, rules map[int][]int) []int {
	for i, page := range pages {
		for _, followingPage := range rules[page] {
			// Check all the elements before this one in the pages array and if
			// any of them are equal to the following page, then swap them so that
			// the rule is satisfied.
			for j := 0; j < i; j++ {
				if pages[j] == followingPage {
					pages[j], pages[i] = pages[i], pages[j]
				}
			}
		}
	}

	// The page list may be violating multiple rules, so keep checking this.
	if !isValidPageLine(pages, rules) {
		return fixOrder(pages, rules)
	}
	return pages
}
