package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x, y int
}

type movement struct {
	xv, yv int // X velocity and Y velocity
}

func main() {
	totaltreeshugged := 1
	var velocities = []movement{
		movement{xv: 1, yv: 1},
		movement{xv: 3, yv: 1},
		movement{xv: 5, yv: 1},
		movement{xv: 7, yv: 1},
		movement{xv: 1, yv: 2},
	}
	for _, mov := range velocities {
		pos := position{x: 0, y: 0}
		treeshugged := 0
		slope, _ := readLines("input.txt")
		for pos.y < len(slope) {
			if checkTree(pos, slope[pos.y]) {
				treeshugged++
			}
			pos = changePosition(pos, mov)
		}
		totaltreeshugged *= treeshugged
		fmt.Printf("With velocity %v, trees hugged: %d (total trees hugged: %d)\n", mov, treeshugged, totaltreeshugged)
	}
}

// checkTree will look for a tree and up the counter
func checkTree(pos position, slope string) bool {
	slopewidth := len(slope)
	positionOnSlope := pos.x % slopewidth
	return string(slope[positionOnSlope]) == "#"
}

// changePosition will change the position according to the movement.
func changePosition(pos position, mov movement) position {
	pos.x += mov.xv
	pos.y += mov.yv
	return pos
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
