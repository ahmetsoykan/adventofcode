package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Content struct {
	paper    bool
	selected bool
}

var contents [][]Content

func main() {

	// read the inputs.txt file
	file, err := os.Open("inputs.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowIndex := 0

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		// Data population into structure
		// Initialize a new row for this line
		var currentRow []Content

		// Check each character in the line for '@'
		for _, char := range line {
			if char == '@' {
				content := Content{
					paper:    true,
					selected: false,
				}

				currentRow = append(currentRow, content)
			} else {
				content := Content{
					paper:    false,
					selected: false,
				}
				currentRow = append(currentRow, content)
			}
		}

		// Add the completed row to contents
		contents = append(contents, currentRow)

		rowIndex++
	}

	// rows 0,1,2
	// columns 0,1,2
	// 1,1 - 00, 01, 02, 10, 11, 12, 20, 21, 22

	// rows 0,1,2
	// columns 4,5,6
	// 1,5 - 04, 05, 06, 14, 15, 16, 24, 25, 26

	// Check against all paper found
	for i := 0; i < len(contents); i++ {
		for y := 0; y < len(contents[i]); y++ {
			if contents[i][y].paper {
				contents[i][y].check(i, y, len(contents)-1, len(contents[i])-1)
			}
		}
	}

	numOfSuitablePapers := 0
	// print all suitable papers
	for i := 0; i < len(contents); i++ {
		for y := 0; y < len(contents[i]); y++ {
			if contents[i][y].paper && contents[i][y].selected {
				numOfSuitablePapers++
			}
		}
	}

	// remove papers by forklift
	for i := 0; i < len(contents); i++ {
		for y := 0; y < len(contents[i]); y++ {
			if contents[i][y].paper && contents[i][y].selected {
				contents[i][y].paper = false
				contents[i][y].selected = false
			}
		}
	}
	fmt.Printf("[part1] Number of papers can be accessed by a forklift: %d\n", numOfSuitablePapers)

	// While loop to continue until no more papers can be accessed by elves
	numOfSuitablePapersbyElves := 1 // Initialize to 1 to enter the loop
	iterationCount := 0

	numOfSuitablePapersbyElvesTotal := 0

	for numOfSuitablePapersbyElves > 0 {
		iterationCount++
		// fmt.Printf("\n--- Iteration %d ---\n", iterationCount)

		// Check against all paper found
		for i := 0; i < len(contents); i++ {
			for y := 0; y < len(contents[i]); y++ {
				if contents[i][y].paper {
					contents[i][y].check(i, y, len(contents)-1, len(contents[i])-1)
				}
			}
		}

		// Count suitable papers for this iteration
		numOfSuitablePapersbyElves = 0
		for i := 0; i < len(contents); i++ {
			for y := 0; y < len(contents[i]); y++ {
				if contents[i][y].paper && contents[i][y].selected {
					numOfSuitablePapersbyElves++
					contents[i][y].paper = false
				}
			}
		}
		numOfSuitablePapersbyElvesTotal = numOfSuitablePapersbyElvesTotal + numOfSuitablePapersbyElves
		// fmt.Printf("Papers accessible by elves in this iteration: %d\n", numOfSuitablePapersbyElves)
	}

	// fmt.Printf("\n[part2] Completed after %d iterations\n", iterationCount)
	fmt.Printf("[part2] Final number of papers accessible by elves and first removal by forklifts: %d\n", numOfSuitablePapers+numOfSuitablePapersbyElvesTotal)

}

func (c *Content) check(currRow, currCol, numOfRows, numOfColumns int) {

	paperCountAround := 0

	lowerBound := currRow - 1
	upperBound := currRow + 1

	rightBound := currCol + 1
	leftBound := currCol - 1

	// safety checks for -1 indexes
	if lowerBound < 0 {
		lowerBound = 0
	}
	if leftBound < 0 {
		leftBound = 0
	}
	// safety checks if we're checking non-existing adjancents
	if upperBound > numOfRows {
		upperBound = numOfRows
	}
	if rightBound > numOfColumns {
		rightBound = numOfColumns
	}

	for i := lowerBound; i <= upperBound; i++ {
		for y := leftBound; y <= rightBound; y++ {

			if (i == currRow) && (y == currCol) {
				// nothing to do
			} else if contents[i][y].paper {
				paperCountAround++
			} else {
				// nothing to do
			}
		}
	}

	if paperCountAround < 4 {
		c.selected = true
	}

}
