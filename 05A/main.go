package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seatingMap, _ := readLines("input.txt")
	numrows := 128
	numcols := 8
	highestSeat := 0
	for _, boardingpass := range seatingMap {
		cols, rows := initializeRowcols(numcols), initializeRowcols(numrows)
		for i := 0; i != len(boardingpass); i++ {
			switch {
			case string(boardingpass[i]) == "F" || string(boardingpass[i]) == "B":
				rows = cutmap(rows, string(boardingpass[i]))
			case string(boardingpass[i]) == "L" || string(boardingpass[i]) == "R":
				cols = cutmap(cols, string(boardingpass[i]))
			}
		}
		seatId := rows[0]*8 + cols[0]
		fmt.Printf("Row: %d, Column: %d, Seatid: %d\n", rows[0], cols[0], seatId)
		if seatId > highestSeat {
			highestSeat = seatId
		}
	}
	fmt.Println("Highest seat:", highestSeat)
}

func cutmap(rowcols []int, letter string) []int {
	half := len(rowcols) / 2
	switch {
	case letter == "F" || letter == "L":
		return rowcols[:half]
	case letter == "B" || letter == "R":
		return rowcols[half:]
	}
	return rowcols
}

func initializeRowcols(num int) []int {
	var initialized []int
	for i := 0; i < num; i++ {
		initialized = append(initialized, i)
	}
	return initialized
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
