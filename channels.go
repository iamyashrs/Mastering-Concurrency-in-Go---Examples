package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var (
	inputString, outputString string
	length                    int
)

func addToFinal(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	outputString += letter
	wg.Done()
}

func capitalize(letterChannel chan string, letter string, wg *sync.WaitGroup) {
	lett := strings.ToUpper(letter)
	wg.Done()
	letterChannel <- lett
}

func main() {
	runtime.GOMAXPROCS(2)
	gg := new(sync.WaitGroup)

	inputString = "we are who we are"

	input := []byte(inputString)

	var letChan chan string = make(chan string)

	length = len(input)

	for i := 0; i < length; i++ {
		gg.Add(2)

		go capitalize(letChan, string(input[i]), gg)
		go addToFinal(letChan, gg)

		gg.Wait()
	}

	fmt.Println(outputString)
}
