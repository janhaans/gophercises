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

	play(*csvPtr)
}

func play(csvFileName string) {
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
	for i, record := range records {
		question := record[0]
		goodAnswer := record[1]
		for {
			fmt.Printf("Question %d: %s = ", i, question)
			answer, err := input.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			if strings.TrimSpace(answer) != goodAnswer {
				fmt.Println("Try again")
			} else {
				break
			}
		}
	}

	fmt.Println("Congratulations you have finished the quiz!")
}
