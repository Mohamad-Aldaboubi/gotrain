package main

import (
	"fmt"
)

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS
var x int

type person struct {
	first string
	last  string
}
type foo int

var y foo

const bar int = 42

func main() {
	p1 := person{
		first: "mohamad",
		last:  "aldaboubi",
	}
	y = 42

	fmt.Println(p1)
	fmt.Printf("%T", int(y))
	fmt.Printf("%T", bar)
	fmt.Println((bar))

}
