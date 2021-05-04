package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}
type Byname []person

func (a Byname) Len() int           { return len(a) }
func (a Byname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] } //swapping
func (a Byname) Less(i, j int) bool { return a[i].Name < a[j].Name }
func main() {

	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}
	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(Byname(people))
	fmt.Println(people)

}
