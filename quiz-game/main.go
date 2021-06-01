package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds (defaul 30)")
	flag.Parse()
	end := make(chan bool)

	records := read(*csvPtr)

	numQuestions := len(records)
	score := 0

	go play(records, &score, *limitPtr, end)
	if <-end {
		fmt.Printf("\nTotal questions = %d\n", numQuestions)
		fmt.Printf("Correct answers = %d\n", score)
	}

}

func read(file string) [][]string {
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

	return records
}

func play(records [][]string, score *int, limit int, c chan bool) {
	input := bufio.NewReader(os.Stdin)

	fmt.Println("Press Enter to start the quiz")
	_, err := input.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	go timer(c, limit)

	for i, record := range records {
		question := record[0]
		answer := record[1]
		fmt.Printf("Question %d: %s = ", i, question)
		playerAnswer, err := input.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if strings.TrimSpace(playerAnswer) == answer {
			*score += 1
		}
	}

	c <- true
}

func timer(c chan bool, limit int) {
	time.Sleep(time.Duration(limit) * time.Second)
	c <- true
}
