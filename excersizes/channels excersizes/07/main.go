package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	cn := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	go grrrr(c, cn)

	for v := range cn {
		fmt.Println(v)
	}

	fmt.Println("gonna go")
}

func grrrr(c, cn chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c {
				func(v2 int) {
					cn <- v2
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(cn)
}
