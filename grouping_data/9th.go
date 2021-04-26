package main

import "fmt"

func main() {
	m := map[string]int{
		"james":           32,
		"miss moneypenny": 27}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["barnabas"])
	v, ok := m["barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["barnabas"]; ok {
		fmt.Println("this is the if print", v)
	}
	if v, ok := m["miss moneypenny"]; ok {
		fmt.Println("this is the if print", v)
	}
}
