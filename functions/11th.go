package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println(x)
	x++
	{
		y := 42
		fmt.Println(y)
	}
	//fmt.println(y)(outside scope)
	fmt.Println()
	foo()
}

func foo() {
	fmt.Println("hello")

}
