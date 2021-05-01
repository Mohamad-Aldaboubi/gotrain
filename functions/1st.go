package main

import (
	"fmt"
)

func main() {
	foo()
	bar("james")
	s1 := woo()

	fmt.Println(s1)
	x, y := mouse("ahmad", "moamad")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("welcome to foo")

}

//everything in go is pass by value
func bar(s string) {

	fmt.Println("hello", s)

}
func woo() string {
	return fmt.Sprint("hello from return function")

}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `,says "hello"`)
	b := false
	return a, b

}
