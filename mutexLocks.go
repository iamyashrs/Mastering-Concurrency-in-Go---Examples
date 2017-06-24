package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	current := 0
	iters := 100

	wg := new(sync.WaitGroup)
	wg.Add(iters)

	var mutex = &sync.Mutex{}

	for i := 0; i < iters; i++ {
		go func() {
			mutex.Lock()
			current++
			fmt.Println(current)
			mutex.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("WHAT")
	wg.Wait()
}
