package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	/*type divFunc = func(x, y float64) (float64, error)
	var div divFunc
	div = division
	res, err := div(10, 3)
	if err != nil {
		panic(err)
	}
	log.Printf("10 делить на 3 будет %f\n", res)
	*/
	x := 10
	f := func() {
		log.Println("X доступна из замыкания:", x)
	}
	f()

	printInts(6, 2, 3, 4)
}

func printInts(ints ...int) {
	for i, v := range ints {
		fmt.Print(i, ":", v, " ")
	}
	fmt.Println()
}

func division(x, y float64) (result float64, err error) {
	if y == 0 {
		return result, errors.New("ошибка: деление на ноль")
	}
	return x / y, nil
}

// Output: 	2020/09/27 14:26:40 10 делить на 3 будет 3.333333
//			0 1 2 3
//			2020/09/27 14:26:40 X доступна из замыкания: 10
