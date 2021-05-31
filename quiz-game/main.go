package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	numQuestions, numAnswers := play(*csvPtr)
	fmt.Printf("Total questions = %d\n", numQuestions)
	fmt.Printf("Correct answers = %d\n", numAnswers)

}

func play(csvFileName string) (numQuestions, numAnswers int) {
	f, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	input := bufio.NewReader(os.Stdin)

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	numQuestions = len(records)

	for i, record := range records {
		question := record[0]
		answer := record[1]
		fmt.Printf("Question %d: %s = ", i, question)
		playerAnswer, err := input.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if strings.TrimSpace(playerAnswer) == answer {
			numAnswers += 1
		}
	}

	return numQuestions, numAnswers
}
