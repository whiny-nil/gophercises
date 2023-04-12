package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const questionField = 0
const answerField = 1

var problemFile = flag.String("f", "problems.csv", "a problem file to use")

func main() {
	flag.Parse()

	file, err := os.Open(*problemFile)
	if err != nil {
		log.Fatalln(err)
	}

	var testChan = make(chan struct{})
	go takeTest(file, testChan)
	var timerChan = make(chan struct{})
	go runTimer(timerChan)

	select {
	case <-testChan:
		return
	case <-timerChan:
		return
	}
}

func takeTest(file *os.File, testChan chan struct{}) {
	reader := csv.NewReader(file)
	rows, _ := reader.ReadAll()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type the answer and press Enter.  Ready?  Go!")

	correct := 0

	for _, row := range rows {
		fmt.Println(row[questionField])
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println("An error occurred reading your answer")
		}
		userAnswer := strings.TrimSpace(scanner.Text())

		if userAnswer == row[answerField] {
			correct += 1
		}
	}

	fmt.Printf("You got %d/%d\n", correct, len(rows))

	testChan <- struct{}{}
}

func runTimer(timerChan chan struct{}) {
	time.Sleep(30 * time.Second)

	fmt.Println("Sorry, you ran out of time")
	timerChan <- struct{}{}
}
