package main

import "fmt"

func main() {
	c := make(chan int)

	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan int) <-chan int {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
