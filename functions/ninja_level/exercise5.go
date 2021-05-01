package main

import (
	"fmt"
	"math"
)

type square struct {
	height float64
	weight float64
}
type circle struct {
	raduis float64
}

func (s square) area() float64 {
	area := s.height * s.weight
	return area
}
func (s circle) area() float64 {
	area := s.raduis * s.raduis * math.Pi
	return area
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())

}

func main() {
	s := square{
		height: 4,
		weight: 4,
	}
	c := circle{
		raduis: 2,
	}

	info(s)
	info(c)
}
