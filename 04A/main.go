package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	passportsList, _ := readLines("input.txt")
	lastLine := len(passportsList) - 1
	validPassports := 0
	passport := make(map[string]string)
	for i, line := range passportsList {
		if line != "" { // Not an empty line, so extract all the values and put them in "passport"
			passport = extractValuePairs(line, passport)
		}
		if i == lastLine || line == "" { // We have all the fields for this passport. Validate it.
			if checkValidPassport(passport) {
				validPassports++
			}
			// fmt.Printf("Passport: %v - Valid:%t - #Valids:%d\n", passport, checkValidPassport(passport), validPassports)
			passport = make(map[string]string)
		}
	}
	fmt.Printf("Valid passports found: %d", validPassports)
}

func extractValuePairs(line string, passport map[string]string) map[string]string {
	pairs := strings.Split(line, " ")
	for _, pair := range pairs {
		passport[pair[0:3]] = pair[4:]
	}
	return passport
}

// checkValidPassport checks if the passport has all the required fields
func checkValidPassport(passport map[string]string) bool {
	return passport["byr"] != "" && passport["pid"] != "" && passport["ecl"] != "" && passport["eyr"] != "" && passport["hcl"] != "" && passport["hgt"] != "" && passport["iyr"] != ""
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
