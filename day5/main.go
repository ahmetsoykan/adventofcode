package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rangesCompleted bool
var ranges = []Range{}
var ids = []string{}

type Range struct {
	Lower, Upper int
}

func main() {
	// read the inputs.txt file
	file, err := os.Open("inputs.txt")
	if err != nil {
		rangesCompleted = true
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			rangesCompleted = true
			continue
		}

		if !rangesCompleted {
			lower, _ := strconv.Atoi(strings.Split(line, "-")[0])
			upper, _ := strconv.Atoi(strings.Split(line, "-")[1])

			ranges = append(ranges, Range{
				Lower: lower,
				Upper: upper,
			})
		} else {
			ids = append(ids, line)
		}
	}

	numOfFreshID := 0
	for _, id := range ids {

		numb, _ := strconv.Atoi(id)

		// PART 1
		isFresh := part1(numb)

		if isFresh {
			numOfFreshID++
		}
	}

	// PART 1
	fmt.Println("PART 1")
	fmt.Println(numOfFreshID)

	// PART 2
	fmt.Println("PART 2")
	part2()
}

func part1(id int) bool {
	for _, r := range ranges {
		if id >= r.Lower && id <= r.Upper {
			return true
		}
	}
	return false
}

func part2() {

	changed := true
	for changed {
		changed = false
		for i, r := range ranges {
			for j, o := range ranges {

				if i == j {
					continue
				}

				if r.Lower > o.Upper || r.Upper < o.Lower {
					// no potential overlap
					// do nothing
					continue
				}

				if r.Lower >= o.Lower && r.Upper <= o.Upper {
					// Contained entirely.
					// lets remove smaller one
					ranges = removeRangeAtIndex(ranges, i)
					changed = true
					break
				}

				if r.Lower <= o.Lower && r.Upper >= o.Lower && r.Upper <= o.Upper {
					// extend lower to r.Lower for o.Lower
					// and remove smaller range, which is r
					ranges[j].Lower = r.Lower
					ranges = removeRangeAtIndex(ranges, i)
					changed = true
					break
				}
				if r.Upper >= o.Upper && r.Lower >= o.Lower && r.Lower <= o.Upper {
					// extend upper to r.Upper for o.Upper
					// and remove smaller range, which is r
					ranges[j].Upper = r.Upper
					ranges = removeRangeAtIndex(ranges, i)
					changed = true
					break
				}

			}
		}
	}
	countPart2 := 0
	for _, r := range ranges {
		countPart2 += 1 + r.Upper - r.Lower
	}
	fmt.Println(countPart2)
}

// Helper functions to remove items from []Range slice
func removeRangeAtIndex(slice []Range, index int) []Range {
	if index < 0 || index >= len(slice) {
		return slice // Index out of bounds, return original slice
	}
	return append(slice[:index], slice[index+1:]...)
}
