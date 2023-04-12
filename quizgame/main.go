package main

import "flag"

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	quizGame := readFile(*csvFileName)
	playGame(quizGame, *timeLimit)
}
