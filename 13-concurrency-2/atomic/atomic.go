package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func handler(counter *uint64) {
	atomic.AddUint64(counter, 1)
}

func main() {
	var c uint64
	N := 100_000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			handler(&c)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Счётчик:", c)
}
