package main

import (
	"fmt"
)

func main() {
	defer foo()
	boo()
}

func foo() {

	fmt.Println("used defer")

}

func boo() {
	fmt.Println("not used deffer")

}
