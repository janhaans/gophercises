//package game creates and plays a quiz in wich questions are asked
package game

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//Problem type has a question and answer
type Problem struct {
	Question string
	Answer   string
}

//Quiz type has problems, manages the score and sets a limit to the duration of the quiz
type Quiz struct {
	Problems []Problem
	Score    int
	Limit    int
}

//GetQuiz receives csv file name that has questions and answers and limit that represents the duration of the quiz in seconds
//and returns a pointer to Quiz
func GetQuiz(file string, limit int) *Quiz {
	quiz := Quiz{Limit: limit}

	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	for _, record := range records {
		problem := Problem{Question: record[0], Answer: record[1]}
		quiz.Problems = append(quiz.Problems, problem)
	}

	return &quiz
}

//Play the quiz
func (q *Quiz) Play() {
	done := make(chan bool)
	go func() {
		input := bufio.NewReader(os.Stdin)

		fmt.Print("Press Enter to start the quiz")
		_, err := input.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		go timer(done, q.Limit)

		for i, problem := range q.Problems {
			fmt.Printf("Question %d: %s = ", i, problem.Question)
			playerAnswer, err := input.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			if strings.TrimSpace(playerAnswer) == problem.Answer {
				q.Score += 1
			}
		}

		done <- true
	}()

	<-done
}

//timer ends the quiz when the time is gone
func timer(done chan bool, limit int) {
	time.Sleep(time.Duration(limit) * time.Second)
	done <- true
}
