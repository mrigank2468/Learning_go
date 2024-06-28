package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("time", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	_ = csvFileName
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFileName)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("fail to provide the csv file")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answeCh:=make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answeCh<-answer
		}()
		select {
		case <-timer.C:
			fmt.Print("You got ", correct, " out of ", len(problems), " correct \n")
			return
		case answer := <-answeCh:
			if answer == p.answer {
				correct++
			}
		}
	}
	fmt.Print("You got ", correct, " out of ", len(problems), " correct \n")
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}
