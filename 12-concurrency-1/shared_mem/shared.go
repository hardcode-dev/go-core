package main

import (
	"fmt"
	"sync"
)

// add прибавляет x к sum
func add(sum *int, x int, wg *sync.WaitGroup, mux *sync.Mutex) {
	defer wg.Done()
	mux.Lock()
	*sum += x
	mux.Unlock()
}

func main() {
	var mux sync.Mutex
	sum := 0

	const N = 10_000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go add(&sum, 1, &wg, &mux)
	}
	wg.Wait()
	fmt.Println(sum)
}
