package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	default:
		fmt.Println("this is default")

	}

}
