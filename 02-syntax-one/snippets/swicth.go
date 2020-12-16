package main

import (
	"fmt"
)

func main() {
	x := 1
	switch x {
	case 1:
		fmt.Println("X = 1")
	case 2:
		{
			fmt.Println("X = 2")
		}
	default:
		fmt.Println("X unknown")

	}
	//***************************************//
	if x > max {
		x = max
	} else {
		x = 0
	}
	// Output: X = 1
}
