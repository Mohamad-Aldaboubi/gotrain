package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p1 person) speak() {

	fmt.Println(p1.first, p1.last, p1.age)
}
func main() {
	x := person{
		first: "mohamad",
		last:  "aldaboubi",
		age:   23,
	}
	x.speak()

}
