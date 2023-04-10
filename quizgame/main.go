package main

import  (
        "flag"
        "fmt"
        "os"
        "bufio"
        "log"
        "strings"
        )

func main() {
    initialiseFlags()
    fileName := "problems.csv"
    quizGame := readFile(fileName)
    totalQuestions := len(quizGame)
    score := 0

    reader := bufio.NewReader(os.Stdin) 

    for _, qAndA := range quizGame {
        fmt.Print(qAndA.question + " = ")
        
        ans, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        if strings.Trim(ans, " \n") == strings.Trim(qAndA.answer, " \n") {
            score++
        }
    }
    fmt.Printf("You score %d out of %d", score, totalQuestions)
}


func initialiseFlags() {
    flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
    flag.Int("limit", 30, "the time limit for the quiz in seconds")
    flag.Parse()
}
