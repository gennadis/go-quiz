package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'problem,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(records)

	score := playQuiz(problems)
	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func playQuiz(problems []problem) int {
	score := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answer := ""
		fmt.Scanf("%s \n", &answer)
		if answer == p.a {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}
	return score
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
