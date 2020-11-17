package main

import (
	"fmt"
)

type myMap map[int]int

func main() {
	var iface interface{} // 0
	var f func()          // 1
	var m myMap           // 2
	m = nil

	if f == nil {
		fmt.Println("f is nil")
	}
	if m == nil {
		fmt.Println("m is nil")
	}
	cycle(iface, f, m)
}

func cycle(ifaces ...interface{}) {
	for i, iface := range ifaces {
		if iface == nil {
			fmt.Println(i)
		}
	}
}
