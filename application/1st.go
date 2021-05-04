package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

//it must be capital letters
func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   32,
	}
	p2 := person{
		First: "miss",
		Last:  "moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}

	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
