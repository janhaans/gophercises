package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type QuizItem struct {
	Question string
	Answer   string
}

type Quiz struct {
	Items []QuizItem
	Score int
	Limit int
}

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
		quizItem := QuizItem{Question: record[0], Answer: record[1]}
		quiz.Items = append(quiz.Items, quizItem)
	}

	return &quiz
}

func (q *Quiz) Play(c chan bool) {
	input := bufio.NewReader(os.Stdin)

	fmt.Println("Press Enter to start the quiz")
	_, err := input.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	go timer(c, q.Limit)

	for i, item := range q.Items {
		fmt.Printf("Question %d: %s = ", i, item.Question)
		playerAnswer, err := input.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if strings.TrimSpace(playerAnswer) == item.Answer {
			q.Score += 1
		}
	}

	c <- true
}

func timer(c chan bool, limit int) {
	time.Sleep(time.Duration(limit) * time.Second)
	c <- true
}
