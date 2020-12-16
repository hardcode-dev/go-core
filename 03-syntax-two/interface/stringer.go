package main

import (
	"fmt"
)

type myInt int

func (i myInt) String() string {
	return fmt.Sprintf("Целое число: %d\n", i)
}

func main() {
	var mi myInt = 10
	fmt.Println(mi)
}

// Output: Целое число: 10
