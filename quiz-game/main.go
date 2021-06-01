package main

import (
	"flag"
	"fmt"

	"github.com/janhaans/quiz-game/quiz"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds (defaul 30)")
	flag.Parse()
	end := make(chan bool)

	quiz := quiz.GetQuiz(*csvPtr, *limitPtr)
	go quiz.Play(end)

	if <-end {
		fmt.Printf("\nTotal questions = %d\n", len(quiz.Items))
		fmt.Printf("Correct answers = %d\n", quiz.Score)
	}

}
