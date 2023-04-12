package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type quiz struct {
	question string
	answer   string
}

func readFile(fileName string) []quiz {
	var qAndA []quiz
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	reader := csv.NewReader(file)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ques := strings.Trim(rec[0], " ")
		ans := strings.Trim(rec[1], " ")
		qAndA = append(qAndA, quiz{question: ques, answer: ans})
	}
	file.Close()
	return qAndA
}
