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

func (s secertagent) speak() {
	fmt.Println("iam", s.first, s.last, "-the secertagent speak")
}

func (p person) speak() {
	fmt.Println("iam", p.first, p.last, "-the person speak")
}

var x int

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("i was passed into barrrrrr", h.(person).first)
	case secertagent:
		fmt.Println("i was passed into barrrrrr", h.(secertagent).first)
	}
	fmt.Println("i was passed into bar", h)
}

type hotdog int

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
	p1 := person{
		first: "dr",
		last:  "yes",
	}
	fmt.Println(s1)
	s1.speak()
	fmt.Println(s2)
	s2.speak()
	fmt.Println(p1)
	bar(s1)
	bar(s2)
	bar(p1)
	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
