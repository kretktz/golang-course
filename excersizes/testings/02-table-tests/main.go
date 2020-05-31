package main

import "fmt"

func main() {
	x := mySum(2, 3)
	fmt.Println(x)
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
