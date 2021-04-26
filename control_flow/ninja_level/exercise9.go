package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "football":
		fmt.Println("go to the mountains!")
	case "basketball":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hell")
	}
}
