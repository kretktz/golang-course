package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("NumCPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var x int64

	var wg sync.WaitGroup
	const y = 100
	wg.Add(y)

	for i := 0; i < y; i++ {
		go func() {
			atomic.AddInt64(&x, 1)
			fmt.Println("count:", atomic.LoadInt64(&x))
			runtime.Gosched()
			wg.Done()
		}()
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("final count", x)
}
