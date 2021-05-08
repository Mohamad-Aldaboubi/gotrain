package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	//recieve
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

/* another solution with buffer
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 5)

	//send
	func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	//recieve
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}*/
