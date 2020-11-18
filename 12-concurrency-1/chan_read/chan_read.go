package main

import (
	"fmt"
	"time"
)

func generator(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	var ch = make(chan int, 20)
	go generator(ch)
	for {
		time.Sleep(time.Millisecond * 500)
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
