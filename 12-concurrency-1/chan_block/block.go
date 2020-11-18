package main

import "fmt"

func main() {
	var ch = make(chan int)
	/*ch <- 10 // как исправить?
	fmt.Println(<-ch)

	val := <-ch // как исправить?
	fmt.Println(val)
	ch <- 20*/
	close(ch)
	val, ok := <-ch
	fmt.Println(val, ok, <-ch)
	ch <- 50
}
