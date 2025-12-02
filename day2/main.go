package main

import (
	"fmt"
	"strconv"
	"strings"
)

// var ranges = []string{"2121212118-2121212124"}
var invalidIDs = make(map[int]string)
var ids = []int{}
var result int

var ranges = []string{"874324-1096487", "6106748-6273465", "1751-4283", "294380-348021", "5217788-5252660", "828815656-828846474", "66486-157652", "477-1035", "20185-55252", "17-47", "375278481-375470130", "141-453", "33680490-33821359", "88845663-88931344", "621298-752726", "21764551-21780350", "58537958-58673847", "9983248-10042949", "4457-9048", "9292891448-9292952618", "4382577-4494092", "199525-259728", "9934981035-9935011120", "6738255458-6738272752", "8275916-8338174", "1-15", "68-128", "7366340343-7366538971", "82803431-82838224", "72410788-72501583"}

func main() {

	for _, v := range ranges {

		start, _ := strconv.Atoi(strings.Split(v, "-")[0])
		end, _ := strconv.Atoi(strings.Split(v, "-")[1])

		for i := start; i <= end; i++ {

			if check1(i) {
				ids = append(ids, i)
				fmt.Printf("Invalid ID found: %d\n", i)
			}
		}
	}

	for _, num := range ids {
		result = result + num
	}

	fmt.Printf("result: %d\n", result)

}

// 11-22 has two invalid IDs, 11 and 22.
func check1(i int) bool {

	// Convert number to string to check for patterns
	numStr := strconv.Itoa(i)
	length := len(numStr)

	// Check for repeated patterns
	// Try different pattern lengths from 1 to half the string length
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Check if the string can be formed by repeating a pattern
		if length%patternLen == 0 {
			pattern := numStr[:patternLen]
			isRepeated := true

			// Check if the entire string is made of this pattern repeated
			for start := 0; start < length; start += patternLen {
				if numStr[start:start+patternLen] != pattern {
					isRepeated = false
					break
				}
			}

			if isRepeated {
				// Uncomment for Part1
				// Only detect patterns that repeat exactly twice
				// if length != patternLen*2 {
				// 	continue
				// }

				fmt.Printf("  -> found repeated pattern: '%s' in %s\n", pattern, numStr)
				invalidIDs[i] = fmt.Sprintf("contains repeated pattern: %s", pattern)
				return true
			}
		}
	}

	return false
}
