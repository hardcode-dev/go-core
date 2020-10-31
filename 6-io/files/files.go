package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(f.Name(), []byte("Текст"), 0666)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Данные файла:\n%s\n", data)
	// вариант используя io.Reader
	file, err := os.Open(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	buf = buf[:n]
	fmt.Printf("Данные файла:\n%s\n", buf)
}
