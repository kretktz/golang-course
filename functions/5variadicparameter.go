package main

import (
	"fmt"
)

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are now adding", v, "to total", sum)
	}
	fmt.Println("the total is now ", sum)
	return sum
}

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the FINAL total is ", x)
}
