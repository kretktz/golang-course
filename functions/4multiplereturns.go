package main

import (
	"fmt"
)

func foo(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says Hello`)
	b := true
	return a, b
}

func main() {
	x, y := foo("Greg", "Krycinski")
	fmt.Println(x)
	fmt.Println(y)

}
