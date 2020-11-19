package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Старт процесса №%d\n", id)
	for job := range jobs {
		time.Sleep(time.Second) // выполнение работы
		fmt.Printf("Процесс №%d выполнил задание №%d\n", id, job)
		results <- 0
	}
}

func main() {
	const W = 5
	const N = 10
	jobs := make(chan int) // буферизованный канал
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(W)
	for i := 0; i < W; i++ {
		go worker(i, jobs, res, &wg)
	}
	go func(ch chan int) {
		for val := range ch {
			_ = val
		}
	}(res)
	for i := 0; i < N; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
