package main

import (
	"fmt"
)

type Pnt struct {
	X, Y int
}

type Pl struct {
	Name string
	Pnt
}

/*func (p Pnt) String() string {
	return fmt.Sprintf("X: %d, Y: %d\n", p.X, p.Y)
}*/

type A struct {
	B struct {
		C int
	}
}

func main() {
	va := A{}
	va.B.C = 10
	fmt.Println(va)
	print("OK")
	return
}
