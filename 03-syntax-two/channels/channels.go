package main

import (
	"fmt"
)

func main() {

	const bufSize = 10

	var bufchan chan int = make(chan int, bufSize) // буферизированный канал
	bufchan <- 10                                  // неблокирующая операция
	val := <-bufchan                               // получение значения в новую переменную
	fmt.Println(val)

	for i := 1; i < bufSize; i++ {
		bufchan <- i * 2
	}
	close(bufchan)

	for val := range bufchan {
		fmt.Println(val)
	}

	var ch chan int // переменная типа канал, принимающий int
	ch = make(chan int)
	go func(ch chan<- int) { // канал только на отправку
		ch <- 100 // блок
	}(ch)
	val = <-ch // блок
	fmt.Println(val)
}
