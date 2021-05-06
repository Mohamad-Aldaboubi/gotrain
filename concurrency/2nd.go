package main

//this is incorrect because we want it to be like that -_-
import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func info(s shape) {
	fmt.Println("area", s.area())
}
func main() {
	c := circle{5}
	fmt.Println(c.area())
}
