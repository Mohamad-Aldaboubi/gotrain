package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	for k, y := range x {
		fmt.Println(k, y)

	}
	fmt.Printf("%T\n", x)
}
