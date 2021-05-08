package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//recieve
	bar(c)
}

//send
func foo(c chan<- int) {
	c <- 43
}

//recieve
func bar(c <-chan int) {
	fmt.Println(<-c)

}
