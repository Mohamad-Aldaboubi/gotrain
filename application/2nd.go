package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"james","Last":"bond","Age":32},{"First":"miss","Last":"moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	//people:=[]person{} it is possible :)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("all of data", people)

	for i, v := range people {
		fmt.Println("\nperson number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
