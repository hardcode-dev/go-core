package main

import (
	"fmt"
	"sync"
)

func prn(n *int, lock sync.Locker, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	fmt.Println(*n)
}

func inc(n *int, lock sync.Locker, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	*n++
}

func main() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	var n int
	const N = 100_000
	wg.Add(N)
	for i := 0; i < N; i++ {
		if i%10 == 0 {
			go inc(&n, &mux, &wg)
			continue
		}
		go prn(&n, &mux, &wg)
	}
	wg.Wait()
	fmt.Println("N:", n)
}
