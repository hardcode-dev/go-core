package main

import (
	"fmt"
	"sync"
)

func process(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
}

func main() {
	src := make(chan int)
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			process(src, res)
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	go func() {
		for i := 0; i < 100; i++ {
			src <- i
		}
		close(src)
	}()

	for val := range res {
		fmt.Println(val)
	}
}
