package main

import (
	"fmt"
)

func main() {

	defer func() {
		fmt.Println("hi")

	}()
	func(s string) {
		fmt.Println(s)

	}("hello")

}
