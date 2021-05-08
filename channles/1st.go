package main

//understanding channels
//it dont run
import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
