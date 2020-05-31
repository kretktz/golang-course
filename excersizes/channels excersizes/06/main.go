package main

import "fmt"

func main() {

	q := make(chan int)

	c := zrob(q)

	wez(c, q)

	fmt.Println("finito")

}

func zrob(q chan int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1 // need this, otherwise it runs infinitely
		close(c)
	}()
	return c
}

func wez(c, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Received from c: ", v)
		case <-q:
			return
		}
	}
}
