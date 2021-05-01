package main

import (
	"fmt"
)

//func (receiver) identifier(parameters) (returns) { code }
type person struct {
	first string
	last  string
}
type secertagent struct {
	person
	ltk bool
}

func main() {
	s1 := secertagent{
		person: person{
			"james",
			"bond",
		},

		ltk: true,
	}
	s2 := secertagent{
		person: person{
			"abo",
			"shoki",
		},

		ltk: true,
	}
	fmt.Println(s1)
	s1.speak()
	fmt.Println(s2)
	s2.speak()
}

func (s secertagent) speak() {
	fmt.Println("iam", s.first, s.last)
}
