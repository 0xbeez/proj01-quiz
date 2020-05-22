package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format q,a")
	timeLimit := flag.Int("limit", 10, "The time limit in seconds")

	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Failed to open csv file: %s", *csvFileName)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to parse csv file: %s", *csvFileName)
		os.Exit(1)
	}
	problems := parseLines(lines)
	//fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =", i+1, p.q)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d/%d\n", correct, len(problems))
			return
		case answer := <-answerChannel:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("You scored %d/%d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
