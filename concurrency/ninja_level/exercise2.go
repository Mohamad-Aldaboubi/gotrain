package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type human interface {
	speak()
}

func (s *person) speak() {
	fmt.Println(s.first, s.last, s.age)

}
func saysomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		first: "james",
		last:  "bond",
		age:   45,
	}

	saysomething(&p1)

}
