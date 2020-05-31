package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("NumCPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	x := 0

	var wg sync.WaitGroup
	const y = 100
	wg.Add(y)

	var m sync.Mutex

	for i := 0; i < y; i++ {
		go func() {
			m.Lock()
			v := x
			// runtime.Gosched() - not needed
			v++
			x = v

			fmt.Println(x)
			m.Unlock()
			wg.Done()
			// fmt.Println("Goroutines", runtime.NumGoroutine())
		}()
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println(x)
}
