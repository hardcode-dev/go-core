package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	var b int64 = 20
	// c := a + b // invalid operation: a + b (mismatched types int and int64)
	c := int64(a) + b
	fmt.Println("Тип с:", reflect.TypeOf(c))

	d := 10
	e := 10.0
	f := 'a'
	g := "ABC"
	h := "E"
	fmt.Println("d, e, f, g, h", reflect.TypeOf(d), reflect.TypeOf(e), reflect.TypeOf(f), reflect.TypeOf(g), reflect.TypeOf(h))
}
