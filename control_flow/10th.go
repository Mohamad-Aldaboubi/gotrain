package main

import "fmt"

func main() {
	fmt.Printf("true && true\t %v", true && true)
	fmt.Printf("true && false\t %v", true && false)
	fmt.Printf("true || true\t%v", true || true)
	fmt.Printf("true || false \t %v", true || false)
	fmt.Printf("!true\t%v", !true)
}
