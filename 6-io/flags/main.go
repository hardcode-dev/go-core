package main

import (
	"flag"
	"fmt"
)

func main() {
	//func String(name string, value string, usage string) *string
	s := flag.String("filename", "file.txt", "имя файла")
	var n int
	flag.IntVar(&n, "n", 10, "количество")
	flag.Parse()
	fmt.Println(*s, n)
}
