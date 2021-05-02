package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeme(p *person) {
	p.first = "mohamad"
	p.last = "aldaboubi"
	p.age = 23
}

func main() {
	x := person{
		first: "james",
		last:  "bond",
		age:   55,
	}
	changeme(&x)
	fmt.Println(x)

}
