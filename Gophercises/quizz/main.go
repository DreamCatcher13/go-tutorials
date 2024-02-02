package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func readCsv(filePath string) [][]string { // returns slice of slices
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		os.Exit(1)
	}
	defer f.Close() // will close file at the end of a function

	csvReader := csv.NewReader(f)       // create a reader that reads from "f"
	records, err := csvReader.ReadAll() // read all field into slice of strings
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
		os.Exit(1)
	}

	return records
}

func parseQuest(lines [][]string) []problem {
	quizz := make([]problem, len(lines))
	for i, line := range lines {
		quizz[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return quizz
}

func main() {
	file := flag.String("file", "problems.csv", "path to file with quiz")
	timeLimit := flag.Int("time", 30, "time limit for the quiz in seconds")
	flag.Parse()

	r := readCsv(*file)
	q := parseQuest(r)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range q {
		fmt.Printf("Problem #%d: %v = ", i+1, p.q)
		answerCh := make(chan string) //creating a channel
		go func() {                   // declaring anonim func
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}() // calling the func

		select {
		case <-timer.C: // timer Fired
			fmt.Printf("\nYou scored %d out of %d", correct, len(q))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(q))
}
