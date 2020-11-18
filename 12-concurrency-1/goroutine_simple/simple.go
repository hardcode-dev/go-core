package main

import "fmt"

func printN(n int) {
	fmt.Println(n)
}

func main() {
	for i := 0; i < 10; i++ {
		go printN(i) // что будет выведено на экран?
	}
}
