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
	pos := position{x: 0, y: 0}
	mov := movement{xv: 3, yv: 1}
	treeshugged := 0
	slope, _ := readLines("input.txt")
	for pos.y < len(slope) {
		if checkTree(pos, slope[pos.y]) {
			treeshugged++
		}
		pos = changePosition(pos, mov)
	}
	fmt.Printf("Trees hugged: %d", treeshugged)
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
