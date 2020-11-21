package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const N = 10
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(n int) {
			// идиоматичная конструкция с `defer`
			// доступ по адресу через замыкание
			// иначе нужно передавать указатель
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
}
