package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	return checkByr(passport["byr"]) && checkPid(passport["pid"]) && checkEcl(passport["ecl"]) && checkEyr(passport["eyr"]) && checkHcl(passport["hcl"]) && checkHgt(passport["hgt"]) && checkIyr(passport["iyr"])
}

// checkByr (Birth Year) - four digits; at least 1920 and at most 2002.
func checkByr(byr string) bool {
	year, _ := strconv.Atoi(byr)
	return year >= 1920 && year <= 2002
}

// checkIyr (Issue Year) - four digits; at least 2010 and at most 2020.
func checkIyr(iyr string) bool {
	year, _ := strconv.Atoi(iyr)
	return year >= 2010 && year <= 2020
}

// checkEyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func checkEyr(eyr string) bool {
	year, _ := strconv.Atoi(eyr)
	return year >= 2020 && year <= 2030
}

// checkHgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func checkHgt(hgt string) bool {
	if hgt == "" || len(hgt) < 3 {
		return false
	}
	l := len(hgt)
	// fmt.Println(string(hgt[l-2:]))
	n, _ := strconv.Atoi(string(hgt[:l-2]))
	switch {
	case string(hgt[l-2:]) == "cm":
		return n >= 150 && n <= 193
	case string(hgt[l-2:]) == "in":
		return n >= 59 && n <= 76
	}
	return false
}

// checkHcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func checkHcl(hcl string) bool {
	if hcl == "" {
		return false
	}
	toegestaan := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if string(hcl[0]) == "#" && len(hcl) == 7 {
		for _, teken := range hcl[1:] {
			if !strings.Contains(toegestaan, string(teken)) {
				return false
			}
		}
		return true
	}
	return false
}

// checkEcl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func checkEcl(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
}

// checkPid (Passport ID) - a nine-digit number, including leading zeroes.
func checkPid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	n, _ := strconv.Atoi(pid)
	return n > 0
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
