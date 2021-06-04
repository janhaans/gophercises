package main

import (
	"flag"
	"fmt"

	"github.com/janhaans/quiz-game/game"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	quiz := game.GetQuiz(*csvPtr, *limitPtr)
	quiz.Play()

	fmt.Printf("\nTotal questions = %d\n", len(quiz.Problems))
	fmt.Printf("Correct answers = %d\n", quiz.Score)

}
