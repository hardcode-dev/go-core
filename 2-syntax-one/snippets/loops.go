package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}
	for index, value := range arr {
		fmt.Printf("Index: %d\tValue: %d\n", index, value)
	}
}

// Output:
// 0
// 1
// 2
// Index: 0	Value: 1
// Index: 1	Value: 2
// Index: 2	Value: 3
/*
// аналог while
for x<10 {
	fmt.Println(x)
}

// бесконечный цикл
for {
	doSmth()
}
*/
