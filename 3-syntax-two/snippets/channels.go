package main

import (
	"fmt"
)

func main() {
	var ch chan int // переменная типа канал, принимающий int
	ch = make(chan int)
	go func(ch chan<- int) { // канал только на отправку
		ch <- 100
	}(ch)
	val := <-ch
	fmt.Println(val)

	var b = []byte{1, 2, 5, 32, 54}
	for _, char := range b {
		fmt.Printf("%v, ", char)
	}
}

// OUTPUT: 100
