package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("value =40")
	} else if x == 41 {
		fmt.Println("value =41")
	} else {
		fmt.Println("value was not 40,41")
	}
}
