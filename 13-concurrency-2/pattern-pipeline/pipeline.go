package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func proc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	src := gen(1, 2, 3, 4, 5)
	res1 := proc(src)
	res2 := proc(res1)

	for val := range res2 {
		fmt.Println(val)
	}
}
