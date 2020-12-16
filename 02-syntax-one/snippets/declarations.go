package main

import "fmt"

var a int                                       // объявление переменной
var b string = "String var"                     // объявление переменной с присвоением значения
var c [4]float64 = [4]float64{1, 2, 3.45, 4.56} // объявление переменной (массива) с присвоением значения
var d []bool = []bool{true, false, true}        // объявление переменной (массива) с присвоением значения

func main() {
	e := 'a'             // короткая запись объявления с присвоением
	type person struct { // объявление локального типа
		name string
	}
	f := person{ // создание переменной и присвоение ей литерала типа person
		name: "Dmitriy",
	}
	g := struct { // создание переменной и присвоение ей литерала безымянного типа
		First string
		Last  string
	}{
		First: "James",
		Last:  "Hetfield",
	}
	var l map[int]string // создание переменной map
	fmt.Println(a, b, c, d, e, f, g, l)
	// l[1] = "1"
}

// Output: 0 String var [1 2 3.45 4.56] [true false true] 97 {Dmitriy} {James Hetfield} map[]
