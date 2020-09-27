package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a := 10 // int
	b := int64(20)
	//a == b // invalid operation: a == b (mismatched types int and int64)
	a = int(b)
	b = int64(a)
	var c float64
	// c = a // cannot use a (type int) as type float64 in assignment
	c = float64(a) + 0.75

	s1 := strconv.Itoa(a) // s = "10"
	s2 := strconv.FormatFloat(c, 'f', 2, 64)
	num := "10" // string
	d, _ := strconv.ParseInt(num, 10, 64)
	fmt.Println(a, b, c, d, s1, s2)
	num += "asd"
	d, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(85))
}

// Output: 	20 20 20.75 10 20 20.75
//			2020/09/27 21:58:07 strconv.ParseInt: parsing "10asd": invalid syntax
//			U
