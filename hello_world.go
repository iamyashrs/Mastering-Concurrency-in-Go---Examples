package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	i    int
	max  int
	text string
}

func runTask(t *task, gg *sync.WaitGroup) {
	for t.i < t.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(t.text)
		t.i++
	}
	gg.Done()
}

func main() {
	goGroup := new(sync.WaitGroup)

	hello := new(task)
	world := new(task)

	hello.i = 0
	hello.max = 3
	hello.text = "hello"

	world.i = 0
	world.max = 5
	world.text = "world"

	go runTask(hello, goGroup)
	go runTask(world, goGroup)

	goGroup.Add(1)
	goGroup.Wait()
}
