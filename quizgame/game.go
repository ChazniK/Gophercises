package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func playGame(quizGame []quiz, timeLimit int) {
	totalQuestions := len(quizGame)
	score := 0

	reader := bufio.NewReader(os.Stdin)
	timer := time.NewTimer(time.Second * time.Duration(timeLimit))

	for _, qAndA := range quizGame {
		fmt.Print(qAndA.question + " = ")

		ansChan := make(chan string)
		go func() {
			answer, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			ansChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", score, totalQuestions)
			return
		case answer := <-ansChan:
			if strings.Trim(answer, " \n") == qAndA.answer {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", score, totalQuestions)
}
