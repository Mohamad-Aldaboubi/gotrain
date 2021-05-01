package main

import (
	"fmt"
)

//factorial with my answer
func main() {

	n := factorial(4)
	fmt.Println(n)
}

func factorial(n int) int {
	for i := (n - 1); i > 0; i-- {
		n *= i
	}
	return n

}
