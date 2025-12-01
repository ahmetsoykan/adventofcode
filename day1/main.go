package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Numbers [100]int = [100]int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
}

type Ring struct {
	CurrIndex      int
	PointTo        int
	ZeroCrossing   int
	ZeroPointStops int
}

func main() {
	// Read the inputs.txt file
	file, err := os.Open("inputs.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	ring := &Ring{
		CurrIndex:      50,
		PointTo:        50,
		ZeroPointStops: 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		// Parse direction (R or L) and number
		direction := line[0]
		numberStr := line[1:]
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Printf("Error parsing number from line '%s': %v\n", line, err)
			continue
		}

		fmt.Printf("Direction: %c, Number: %d\n", direction, number)

		// Apply the movement based on direction
		switch direction {
		case 'R':
			ring.MoveToRight(number)
			fmt.Printf("Moved right %d steps, now at: %+v\n", number, ring)
		case 'L':
			ring.MoveToLeft(number)
			fmt.Printf("Moved left %d steps, now at: %+v\n", number, ring)
		default:
			fmt.Printf("Unknown direction: %c\n", direction)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("Final ring state: %+v\n", ring)
}

func (r *Ring) MoveToLeft(movingIndex int) {

	if movingIndex > 100 {
		loops := (movingIndex / 100) % 10
		r.ZeroCrossing = r.ZeroCrossing + loops
	}

	movingIndex = movingIndex % 100

	if r.CurrIndex-movingIndex < 0 {

		if r.CurrIndex != 0 {
			r.ZeroCrossing++
		}
		r.CurrIndex = (r.CurrIndex - movingIndex + 100) % 100
	} else {
		r.CurrIndex = r.CurrIndex - movingIndex
	}

	r.PointTo = Numbers[r.CurrIndex]

	if r.CurrIndex == 0 {
		r.ZeroPointStops++
	}

}

func (r *Ring) MoveToRight(movingIndex int) {

	if movingIndex > 100 {
		loops := (movingIndex / 100) % 10
		r.ZeroCrossing = r.ZeroCrossing + loops
	}

	movingIndex = movingIndex % 100

	if r.CurrIndex+movingIndex >= 100 {
		if r.CurrIndex+movingIndex != 100 {
			r.ZeroCrossing++
		}

		r.CurrIndex = (r.CurrIndex + movingIndex) % 100
	} else {
		r.CurrIndex = r.CurrIndex + movingIndex
	}

	if r.CurrIndex == 0 {
		r.ZeroPointStops++
	}

	r.PointTo = Numbers[r.CurrIndex]
}
