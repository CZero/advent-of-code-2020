package main

import (
	"bufio"
	"fmt"
	"os"
)

var totalAnswered []int              // The number of answers per group
var groupAnswered int                // The number of answers in this group
var answered = make(map[string]bool) // A map containing the answered questions

func main() {
	answers, _ := readLines("input.txt")
	for _, line := range answers {
		if line != "" { // This is the end of a group
			for _, answer := range line {
				if !answered[string(answer)] {
					groupAnswered++
					answered[string(answer)] = true
				}
			}
		} else {
			endGroup()
		}
	}
	endGroup() // After last line
	fmt.Printf("Sum: %d", sumTotalAnswered())
}

// endGroup appends the unique answers of the group to the totalAnswered range and then resets the group and answered counts
func endGroup() {
	totalAnswered = append(totalAnswered, groupAnswered)
	groupAnswered = 0
	answered = make(map[string]bool)
}

func sumTotalAnswered() (sum int) {
	for _, number := range totalAnswered {
		sum = sum + number
	}
	return sum
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
