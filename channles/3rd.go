package main

//understanding channels
//it run by buffer
import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

}
