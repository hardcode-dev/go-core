package main

import (
	"fmt"
	"sync"
)

// add прибавляет x к sum
func add(sum *int, x int, wg *sync.WaitGroup) {
	defer wg.Done()
	*sum += x
}

func main() {
	sum := 0
	const N = 10_000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go add(&sum, 1, &wg)
	}
	wg.Wait()
	fmt.Println(sum)
}
