package main

import (
	"bufio"
	"fmt"
	"os"
)

type seat struct {
	row int
	col int
}

var (
	numrows     = 128
	numcols     = 8
	highestSeat = 0
	seatIdsSeen = make(map[int]bool)
	seatsSeen   = make(map[seat]bool)
)

func main() {
	seatingMap, _ := readLines("input.txt")
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
		seatsSeen[seat{rows[0], cols[0]}] = true
		seatIdsSeen[seatId] = true
		if seatId > highestSeat {
			highestSeat = seatId
		}
	}
	fmt.Println("Highest seat:", highestSeat)
	freeseats := checkAllTheSeats()
	for _, seat := range freeseats {
		seatId := seat.row*8 + seat.col
		if checkFreeSeat(seatId) {
			println("Free seat is: ", seatId)
		}
	}
}

func checkAllTheSeats() (freeseats []seat) {
	cols, rows := initializeRowcols(numcols), initializeRowcols(numrows)
	for _, row := range rows {
		for _, col := range cols {
			if !seatsSeen[seat{row, col}] {
				freeseats = append(freeseats, seat{row, col})
			}
		}
	}
	return freeseats
}

func checkFreeSeat(seatid int) bool {
	return seatIdsSeen[seatid-1] && seatIdsSeen[seatid+1]
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
