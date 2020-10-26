package main

import "fmt"

type Product struct {
	Name string
	Attribute
	Color
}
type Attribute struct {
	ID   int
	Name string
}
type Color struct {
	ID int
}

func main() {
	a := Attribute{
		ID:   100,
		Name: "Color",
	}
	c := Color{
		ID: 300,
	}
	p := Product{
		Name:      "Computer",
		Attribute: a,
		Color:     c,
	}

	fmt.Println(p.Color.ID, p.Attribute.ID)
}
