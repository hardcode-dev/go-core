package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStr string
	var a string = "a"
	var b myStr = "b"
	_, _ = a, b

	//a = b //cannot use b (type myStr) as type string in assignment

	const c = "c"
	fmt.Println("Тип с:", reflect.TypeOf(c))

	b = c

	fmt.Println("Тип b:", reflect.TypeOf(b))
	fmt.Println("Значение b:", b)
}
