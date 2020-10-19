package main

import (
	"fmt"
	"go-core/3-syntax-two/interface2/pkg"
)

func main() {
	p := pkg.Person{
		Name: "Дмитрий",
		Age:  37,
	}
	printInfo(p)
}

func printInfo(p pkg.Person) {
	fmt.Println(p.Info())
}
