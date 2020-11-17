package main

import (
	"fmt"
	"log"
)

type cassete struct {
	label string
}

type modern interface {
	isModern() bool
}

type old interface {
	isOld() bool
	isObsolete() bool
}

func (c cassete) isModern() bool {
	return false
}
func (c cassete) isOld() bool {
	return true
}
func (c cassete) isObsolete() bool {
	return false
}

func main() {
	var m modern
	var c cassete
	m = c
	// m.label = "Ace Of Base" - ошибка
	m.isModern()
	if cassete, ok := m.(cassete); ok {
		cassete.label = "Ace Of Base"
		log.Printf("%+v", cassete)
	}
	if cassete, ok := m.(old); ok {
		fmt.Println("Old: ", cassete.isOld(), "\tObsolete: ", cassete.isObsolete())
	}
}
