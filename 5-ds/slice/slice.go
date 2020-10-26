package main

import (
	"fmt"
)

func main() {
	var s []int
	oldCap := -1
	for i := 0; i < 10_000_000; i++ {
		s = append(s, i) // какие проблемы с производительностью это может вызвать?
		if cap(s) != oldCap {
			var ratio float64
			if oldCap > 0 {
				ratio = float64(cap(s)) / float64(oldCap)
			}
			fmt.Printf("len() %d\tcap() %d\tratio %f\n", len(s), cap(s), ratio)
			oldCap = cap(s)
		}
	}
}
