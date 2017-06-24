package main

import (
	"fmt"
	"strings"
)

var (
	inputString, outputString string
	length, lettersDone       int
	inputBytes                []byte
	applicationDone           bool
)

func sendLetters(gq chan string) {
	for i := 0; i < length; i++ {
		gq <- string(inputBytes[i])
	}
}

func capitalize(gq chan string, sq chan string) {
	for {
		if lettersDone >= length {
			applicationDone = true
			break
		}
		select {
		case let := <-sq:
			outputString += let
			lettersDone++
		case letter := <-gq:
			letter = strings.ToUpper(letter)
			sq <- letter
			// outputString += letter
		}
	}
}

func main() {
	applicationDone = false

	gQueue := make(chan string)
	sQueue := make(chan string)

	fmt.Println("Lets start..")

	lettersDone = 0

	inputString = "we are who we are"
	inputBytes = []byte(inputString)

	length = len(inputBytes)

	go sendLetters(gQueue)
	capitalize(gQueue, sQueue)

	close(gQueue)
	close(sQueue)

	for {
		if applicationDone == true {
			fmt.Println("Done")
			fmt.Println(outputString)
			break
		}
	}
}
