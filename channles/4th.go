package main

//receive and send channels
import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) // sent

	fmt.Println("------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//specific to general doesnt assign
	//c=cr
	//c=cs
	//specific to general doesnt convert
	//fmt.Printf("c\t%T\n",(chan int)(cr))
	//fmt.Printf("c\t%T\n",(chan int)(cs))

}
