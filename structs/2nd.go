package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type secertagent struct {
	person
	itk bool
}

func main() {

	sa1 := secertagent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32},
		itk: true,
	}

	p2 := person{
		first: "miss",
		last:  "moneybenny",
		age:   27,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
