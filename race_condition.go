package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var balance int
var breakPoint bool

func main() {
	rand.Seed(time.Now().Unix())
	balance = 1000yash

	wg := new(sync.WaitGroup)

	done := make(chan bool)
	transChannel := make(chan int)

	wg.Add(1)

	for i := 0; i < 100; i++ {
		go func(j int) {
			tranAmount := rand.Intn(25)
			transChannel <- tranAmount
			if j == 99 {
				done <- true
				close(transChannel)
				wg.Done()
			}
		}(i)
	}

	for {
		if breakPoint {
			break
		}

		select {
		case amt := <-transChannel:
			if amt > 0 {
				// balance -= tranAmount
				balance -= amt
			}
		case status := <-done:
			if status {
				fmt.Println("Finished!")
				breakPoint = true
				close(done)
			}
		}
	}

	wg.Wait()

	fmt.Println("Final Amount left:" + strconv.Itoa(balance))
}
