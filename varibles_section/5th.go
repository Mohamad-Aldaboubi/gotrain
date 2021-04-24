package main

import "fmt"

var y = 42

//declared the variable with identifier z
//is of type string
//and iam assign the value "shaken,not stirred"

var z string = "shaken, not stirred"
var a string = "james said ,'shaken, not stirred'"

func main() {
	fmt.Println(z)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
