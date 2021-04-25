package main

import (
	"fmt"
)

//الكونست رقم ثابت م بتغير
const (
	a int     = 42
	b float64 = 42.78
	c string  = "james bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
