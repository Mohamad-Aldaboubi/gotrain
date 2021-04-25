package main

import "fmt"

func main() {
	a := (7 == 7)
	b := (7 <= 5)
	c := (7 >= 5)
	d := (7 != 5)
	e := (7 < 5)
	f := (7 > 5)
	fmt.Println(a, b, c, d, e, f)
}
