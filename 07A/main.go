package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bag struct {
	name     string
	contents map[string]int
}

var (
	bags       []Bag
	bagsFound  []string
	lookingFor = "shiny gold"
)

func main() {
	start := time.Now()
	rules, _ := readLines("input.txt")
	for _, rule := range rules {
		parseBag(rule)
	}
	for _, bag := range bags {
		lookInBag(bag, lookingFor)
	}
	fmt.Printf("Number of possabilities: %d\n", len(bagsFound))
	duration := time.Since(start)
	fmt.Printf("Duration: %s", duration)

}

func lookInBag(openBag Bag, lookingFor string) {
	if len(openBag.contents) == 0 { // A bag without contents doesn't need searching.
		return
	}
	for bag := range openBag.contents {
		if bag != lookingFor { // Only recurse when this isn't the bag we're looking for.
			for _, newBag := range bags {
				if newBag.name == bag && !contains(bagsFound, newBag.name) { // Check if we haven't found it before.
					lookInBag(newBag, lookingFor)
				}
			}
		} else {
			if !contains(bagsFound, openBag.name) {
				bagsFound = append(bagsFound, openBag.name)
				for _, bag := range bags {
					if !contains(bagsFound, bag.name) { // Only recurse it if we haven't found it before.
						lookInBag(bag, openBag.name)
					}
				}
			}
		}
	}
}

func contains(bagsFound []string, bag string) bool {
	for _, examine := range bagsFound {
		if examine == bag {
			return true
		}
	}
	return false
}

func parseBag(line string) {
	bagRule := strings.Split(line, " contain ")                                                      // Break line in 2 parts: mainBag and contents
	mainBag := strings.Split(bagRule[0], " ")[0:2][0] + " " + strings.Split(bagRule[0], " ")[0:2][1] // Remove bag or bags at the end
	bag := Bag{name: mainBag}
	bag.contents = make(map[string]int)
	rules := strings.Split(bagRule[1], ", ")
	for _, rule := range rules {
		if rule != "no other bags." {
			words := strings.Split(rule, " ")[0:3]
			containBag := words[1] + " " + words[2]
			num, _ := strconv.Atoi(words[0])
			bag.contents[containBag] = num
		}
	}
	bags = append(bags, bag)
	return
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
