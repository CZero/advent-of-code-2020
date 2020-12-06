package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type components struct {
	min    int
	max    int
	letter string
	pass   string
}

func main() {
	passwordlist, _ := readLines("input.txt")
	correctpasses := len(passwordlist)
	for _, passwordentry := range passwordlist {
		parts := disect(passwordentry)
		times := strings.Count(parts.pass, parts.letter)
		if times > parts.max || times < parts.min {
			fmt.Printf("Wrong! %s is in %s %d times, while min:%d and max:%d\n", parts.letter, parts.pass, times, parts.min, parts.max)
			correctpasses--
		}
	}
	fmt.Printf("This leaves %d correct passes.", correctpasses)
}

// disect breaks the string into a struct and returns it.
func disect(input string) components {
	min, _ := strconv.Atoi(strings.Split(input, "-")[0])
	max, _ := strconv.Atoi(strings.Split(strings.Split(input, " ")[0], "-")[1])
	letter := string(strings.Split(input, " ")[1][0])
	pass := strings.Split(input, " ")[2]
	return components{min, max, letter, pass}
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
