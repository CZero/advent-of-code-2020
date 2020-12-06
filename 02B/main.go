package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type components struct {
	posa   int
	posb   int
	letter string
	pass   string
}

func main() {
	passwordlist, _ := readLines("input.txt")
	correctpasses := len(passwordlist)
	for _, passwordentry := range passwordlist {
		parts := disect(passwordentry)
		switch {
		case string(parts.pass[parts.posa-1]) == parts.letter && string(parts.pass[parts.posb-1]) != parts.letter: // Correct, so do nothing
		case string(parts.pass[parts.posa-1]) != parts.letter && string(parts.pass[parts.posb-1]) == parts.letter: // Correct, so do nothing
		default: // Wrong!
			fmt.Printf("%s is invalid!\n", passwordentry)
			correctpasses--
		}
	}
	fmt.Printf("This leaves %d correct passes.", correctpasses)
}

// disect breaks the string into a struct and returns it.
func disect(input string) components {
	posa, _ := strconv.Atoi(strings.Split(input, "-")[0])
	posb, _ := strconv.Atoi(strings.Split(strings.Split(input, " ")[0], "-")[1])
	letter := string(strings.Split(input, " ")[1][0])
	pass := strings.Split(input, " ")[2]
	return components{posa, posb, letter, pass}
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
