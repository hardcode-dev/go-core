package main

import (
	"fmt"
	"sync"
)

func printN(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)
}

func generate(ch chan<- string) {
	ch <- "message from generator"
}
func process(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	const N = 10
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go printN(i, &wg)
	}
	wg.Wait()

	var ch = make(chan string)
	go generate(ch)
	process(ch)
}
