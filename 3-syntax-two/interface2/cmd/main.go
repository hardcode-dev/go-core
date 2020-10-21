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
	printInfo(&p)
}

type Informer interface {
	Info() string
}

func printInfo(i Informer) {
	fmt.Println(i.Info())
}
