package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p *int
	fmt.Println(unsafe.Sizeof(p))
}
