package main

import (
	"flag"
	"fmt"
	"time"

	quiz "github.com/czeslavo/gophercises/quiz/lib"
)

func main() {
	csvFilename, timeLimit := parseArgs()
	problems := quiz.GetProblemsFromCsvFile(csvFilename)

	var (
		correctAnswers    = 0
		answerChan        = make(chan string)
		timeLimitDuration = time.Duration(timeLimit) * time.Second
	)

	for i, problem := range problems {
		timer := time.NewTimer(timeLimitDuration)
		fmt.Printf("Problem #%d: %s?\n", i, problem.Question)

		go quiz.GetUserAnswer(answerChan)
		select {
		case <-timer.C:
		case answer := <-answerChan:
			correctAnswers += quiz.Score(answer, problem.Answer)
		}
	}

	fmt.Printf("You scored %d out of %d!\n", correctAnswers, len(problems))
}

func parseArgs() (csvFilename string, timeLimit int) {
	csvFilenameFlag := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimitFlag := flag.Int("limit", 5, "the time limit for the quiz in seconds")
	flag.Parse()
	return *csvFilenameFlag, *timeLimitFlag
}
