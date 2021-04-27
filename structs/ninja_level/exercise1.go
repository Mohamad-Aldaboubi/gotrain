package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favflavors []string
}

func main() {
	x := person{
		first: "ahmad",
		last:  "mohamad",
		favflavors: []string{
			"chocolate",
			"vanilla",
			"strweberry",
		},
	}
	y := person{
		first: "yazan",
		last:  "saddam",
		favflavors: []string{
			"vanilla",
			"strweberry",
			"chocolate",
		},
	}
	fmt.Println(x)
	fmt.Println(y)
}
