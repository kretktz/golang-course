package main

import (
	"fmt"
)

func foo(s string) string {
	return fmt.Sprint("I am printing a return, which is: ", s)
}

func main() {
	x := foo("BollockS")
	fmt.Println(x)
}
