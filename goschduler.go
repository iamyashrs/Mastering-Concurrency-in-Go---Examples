package main

import (
	"fmt"
	"runtime"
)

func listThread() int {
	threads := runtime.GOMAXPROCS(0)
	return threads
}

func main() {
	runtime.GOMAXPROCS(0)
	fmt.Println(listThread())

	fmt.Println(runtime.NumCPU())
}
