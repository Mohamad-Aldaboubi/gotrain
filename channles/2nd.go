package main

//understanding channels
//it run by go routine
import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
