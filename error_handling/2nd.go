package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string
	fmt.Print("name:")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}
	fmt.Print("fav food:")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}

	fmt.Print("fav sport:")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans1, ans2, ans3)
}
