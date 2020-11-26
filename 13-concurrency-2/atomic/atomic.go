package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var c uint64
	N := 100_000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			atomic.AddUint64(&c, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Счётчик:", c)
}
