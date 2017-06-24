package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeSystem struct {
	updates     int
	currentTime time.Time
	lock        sync.RWMutex
}

var Ts TimeSystem

func update(et chan bool) {
	Ts.lock.Lock()
	defer Ts.lock.Unlock()

	Ts.currentTime = time.Now()
	Ts.updates++

	if Ts.updates == 2 {
		et <- false
	}
}

func main() {
	wg := new(sync.WaitGroup)

	Ts.updates = 0
	Ts.currentTime = time.Now()

	timer := time.NewTicker(1 * time.Second)
	writeTimer := time.NewTicker(10 * time.Second)
	endTimer := make(chan bool)

	breakPoint := false

	wg.Add(1)
	for {
		if breakPoint {
			break
		}

		select {
		case <-timer.C:
			fmt.Println(Ts.updates, Ts.currentTime.String())
		case <-writeTimer.C:
			update(endTimer)
		case <-endTimer:
			timer.Stop()
			close(endTimer)
			wg.Done()
			breakPoint = true
			// return
		}
	}

	wg.Wait()
	fmt.Println(Ts.currentTime.String())
}
