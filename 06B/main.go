package main

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	totalAnswered   int // The number of answers per group
	questionAnswers map[string]int
	answers         []string
}

var groups []group

func main() {
	answers, _ := readLines("input.txt")
	groupid := 0
	for _, line := range answers {
		if line != "" { // This would be the end of a group
			parseGroupanswers(line, groupid)
		} else {
			groupid++ // The next group starts
		}
	}
	unanimousVotes := countUnanimous()

	fmt.Printf("Unanimous votes: %d", unanimousVotes)
}

// parseGroupanswers checks the answers in a line and counts the voters and the totals. It initalizes when fresh.
func parseGroupanswers(line string, groupid int) {
	if groups == nil || len(groups) <= groupid {
		groups = append(groups, group{totalAnswered: 0, answers: []string{}})
		groups[groupid].questionAnswers = make(map[string]int)
	}
	groups[groupid].totalAnswered++ // A line is a persone
	for _, answer := range line {
		if groups[groupid].questionAnswers[string(answer)] == 0 {
			groups[groupid].answers = append(groups[groupid].answers, string(answer))
		}
		groups[groupid].questionAnswers[string(answer)]++
	}
	return
}

// countUnanimous walks through the groups and compares the number of times an answer was given with the total number of people / lines within the group.
// When those are the same it's a unanimous vote and it's counted as such.
func countUnanimous() (unanimousVotes int) {
	for _, group := range groups {
		for _, answer := range group.answers {
			if group.questionAnswers[answer] == group.totalAnswered {
				unanimousVotes++
			}
		}
	}
	return unanimousVotes
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
