package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type biggies struct {
	index int
	digit int
}

var (
	joltages []int
	result   = 0

	expectedDigitCount = 12
	biggieNumbers      = make([]biggies, expectedDigitCount)
)

func main() {

	// read the inputs.txt file
	file, err := os.Open("inputs.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentJoltage := ""

		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		fmt.Println(line)
		startingIndex := 0

		for i := 0; i < len(biggieNumbers); i++ {

			// Extract digits one by one from the line
			// Finding the biggest digit and its index

			currentBiggestDigit := 0
			currentBiggestIndex := 0

			for index := startingIndex; index < len(line)-(len(biggieNumbers)-(i+1)); index++ {
				char := line[index]
				if char >= '0' && char <= '9' {
					digit := int(char - '0') // Convert character to digit
					currentBiggestDigit, currentBiggestIndex = updateCurrentBiggestDigit(currentBiggestDigit, currentBiggestIndex, digit, index)
					fmt.Printf("[%d] currentBiggestDigit: %d, currentBiggestIndex: %d\n", i, currentBiggestDigit, currentBiggestIndex)
				}
			}

			startingIndex = currentBiggestIndex + 1
			currentJoltage = currentJoltage + fmt.Sprintf("%d", currentBiggestDigit)

			fmt.Printf("joltages: %s\n", currentJoltage)

			if len(currentJoltage) == expectedDigitCount {
				_joltage, _ := strconv.Atoi(currentJoltage)
				joltages = append(joltages, _joltage)
			}
		}

	}

	for _, num := range joltages {
		result = result + num
	}

	fmt.Printf("result: %d\n", result)

}

func updateCurrentBiggestDigit(biggestDigit, biggestIndex, curr, currIndex int) (int, int) {
	if curr > biggestDigit {
		return curr, currIndex
	}
	return biggestDigit, biggestIndex
}
