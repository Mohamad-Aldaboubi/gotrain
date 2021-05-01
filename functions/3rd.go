package main

import (
	"fmt"
)

func main() {
	y := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(y...)
	fmt.Println(x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in idex position, ", i, "we are now adding,", v, "to the total which is now", sum)
		fmt.Println("the total is ", sum)
	}
	return sum
}
