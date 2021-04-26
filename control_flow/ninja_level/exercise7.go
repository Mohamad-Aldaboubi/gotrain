package main

import "fmt"

func main() {
	x := 105
	if x >= 50 && x <= 100 {
		fmt.Println("pass")
	} else if x < 50 && x >= 0 {
		fmt.Println("fail")
	} else {
		fmt.Println("error")
	}
}
