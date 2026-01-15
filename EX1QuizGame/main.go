package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func ReadCSVFile(FilePath string) ([][]string, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatal("Unable to read file from:", FilePath, "Error: ", err)
	}
	defer file.Close()

	CSVReader := csv.NewReader(file)
	Records, err := CSVReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for file:", FilePath, "Error: ", err)
	}
	fmt.Print("Successfully read data from file:", FilePath, "\n")

	return Records, nil
}

func MakeQuiz(Records [][]string) map[string]string {
	quiz := make(map[string]string)
	for _, record := range Records {
		question := record[0]
		answer := record[1]
		quiz[question] = answer
	}
	return quiz

}

func main() {
	Records, _ := ReadCSVFile("data.csv")
	quiz := MakeQuiz(Records)

	timer := time.NewTimer(5 * time.Second)

	var score int
	scanner := bufio.NewScanner(os.Stdin)

	//count down timer ถ้าเวลาหมดให้ตัดจบเลย
	go func() {
		<-timer.C //เป็นตัวรับ สัญญาณจาก timer ว่าหมดเวลาแล้ว
		fmt.Printf("\nTime's up! You scored %d out of %d\n", score, len(quiz))
		os.Exit(0)
	}()

	for question, ans := range quiz {
		fmt.Printf("Question %s:", question+"")
		userAnswer := scanner.Text()
		if scanner.Scan() {
			userAnswer = scanner.Text()
			if userAnswer == ans {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", score, len(quiz))

}
