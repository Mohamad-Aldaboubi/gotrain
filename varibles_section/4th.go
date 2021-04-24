package main

import "fmt"

//declare the variable y
//assign the value 43
//declare &assign =initilizationgit
var y = 43

//assign the zero value of type int to z
var z int

func main() {
	//short declaration oprator
	//declare a variable and assign a value(of certain type)
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()
	fmt.Println(z)

}

func foo() {
	fmt.Println("again", y)

}
