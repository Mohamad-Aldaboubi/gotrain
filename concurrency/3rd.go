package main

import (
	"fmt"
)

func dosomething(x int) int {
	return x * 5

}

func main() {

	ch := make(chan int)
	go func() {
		ch <- dosomething(5)

	}()
	fmt.Println(<-ch)

	fmt.Println()

}
