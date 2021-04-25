package main

import "fmt"

var x int
var y float64

func main() {
	x = 2
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
