package main

import "fmt"

func main() {
	m := map[string]int{
		"james":           32,
		"miss moneypenny": 27}
	fmt.Println(m)
	delete(m, "james")
	fmt.Println(m)
	delete(m, "aboshoqy")
	fmt.Println(m)
	if v, ok := m["miss moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "miss moneypenny")
	}
	fmt.Println(m)
}
