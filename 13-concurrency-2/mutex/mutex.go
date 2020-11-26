package main

import (
	"fmt"
	"sync"
)

func prn(n *int, lock *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.RLock()
	defer lock.RUnlock()
	fmt.Println(*n)
}

func inc(n *int, lock sync.Locker, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	*n++
}

func main() {
	var mux sync.RWMutex
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

// [Done] exited with code=0 in 1.689 seconds - Mutex
// [Done] exited with code=0 in 0.984 seconds - RWMutex
