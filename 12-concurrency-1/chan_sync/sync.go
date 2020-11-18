package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pulsar(ch chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- rand.Intn(1000)
	}
	close(ch)
}

func printer(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int) // трубка между потоками pulsar и ?
	go pulsar(ch)
	printer(ch)
}
