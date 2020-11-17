package main

import "fmt"

type car struct {
	model string
	price int
	year  int
}

type bike struct {
	model string
	price int
	year  int
}

func pricesSum[T any](products ...T) int {
	var sum int
	for _, p := range products {
			sum += p.price
}
	return sum
}

func main() {
	c := car{price: 100}
	b := bike{price: 50}
	sum := pricesSum(c, b)
	fmt.Println(sum)
}
