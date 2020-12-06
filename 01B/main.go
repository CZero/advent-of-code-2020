package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// De opdracht is: Vindt de drie regels

func main() {
	report, _ := readLines("input.txt")
	for i, first := range report {
		for n, second := range report {
			if n != i {
				for m, third := range report {
					if m != n && m != i {
						firstn, _ := strconv.Atoi(first)
						secondn, _ := strconv.Atoi(second)
						thirdn, _ := strconv.Atoi(third)
						if firstn+secondn+thirdn == 2020 {
							fmt.Printf("%s + %s + %s = 2020!\n%s * %s *%s = %d\n", first, second, third, first, second, third, firstn*secondn*thirdn)
						}
					}
				}
			}
		}
	}
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

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
