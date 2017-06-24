package main

import (
	"fmt"
)

func main() {
	i := new(int)
	defer fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println(*i)
		*i++
	}
}
