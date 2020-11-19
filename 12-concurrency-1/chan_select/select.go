package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generator(ch chan<- string, num int) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		ch <- "Сообщение из канала №" + strconv.Itoa(num)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go generator(ch1, 1)
	go generator(ch2, 2)

	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}
