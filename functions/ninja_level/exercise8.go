package main

import (
	"fmt"
)

func anything() func() int {

	return func() int {

		return 1998

	}

}

func main() {
	x := anything()
	fmt.Printf("%T\n", x)

	fmt.Println(x())
	fmt.Println(anything()())
}
