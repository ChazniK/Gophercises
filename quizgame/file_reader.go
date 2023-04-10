package main

import  (
        "os"
        "io"
        "log"
        "encoding/csv"
        )


type quiz struct {
    question string
    answer string
}

func readFile(fileName string) []quiz {
    var qAndA [] quiz
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
        qAndA = append(qAndA, quiz{question: rec[0], answer: rec[1]})
    }
    file.Close()
    return qAndA
}

