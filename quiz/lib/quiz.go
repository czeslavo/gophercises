package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// GetProblemsFromCsvFile reads Problems from CSV file in format 'problem,answer'
func GetProblemsFromCsvFile(filename string) []Problem {

	f, err := os.Open(filename)
	check(err)

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	check(err)

	return parseLines(lines)
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i].Question, problems[i].Answer = line[0], strings.TrimSpace(line[1])
	}

	return problems
}

// GetUserAnswer is a goroutine which scan user's answer and pushes it to the answerChan
func GetUserAnswer(answerChan chan string) {
	var answer string
	fmt.Scanln(&answer)

	answerChan <- answer
}

// Score validates user's answer. Returns 1 on success or 0 on failure.
func Score(userAnswer, correctAnswer string) int {
	if userAnswer == correctAnswer {
		return 1
	}
	return 0
}

// Problem with a question and a correct answer
type Problem struct {
	Question, Answer string
}
