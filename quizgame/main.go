package main

import  (
        "flag"
        "fmt"
        "io"
        )

func main() {
    initialiseFlags()
    fileName := "problems.csv"
    quizGame := readFile(fileName)
}


func initialiseFlags() {
    flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
    flag.Int("limit", 30, "the time limit for the quiz in seconds")
    flag.Parse()
}
