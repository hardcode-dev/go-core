package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch = make(chan int)
	/*go func() {
		time.Sleep(time.Second * 2)
		ch <- 10
	}() // как исправить?
	fmt.Println(<-ch)*/
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		val := <-ch // как исправить?
		fmt.Println(val)
	}(&wg)
	ch <- 20
	wg.Wait()
	/*close(ch)
	val, ok := <-ch
	fmt.Println(val, ok, <-ch)
	val, ok = <-ch
	fmt.Println(val, ok, <-ch)
	ch <- 50*/
}
