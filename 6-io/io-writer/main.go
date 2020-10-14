package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
	age  int
}

func main() {
	log.Fatal("er")
	p := person{"Dmitriy", 37}
	var b bytes.Buffer        // буфер, выполняет контракт io.Writer
	b.Write([]byte("Hello ")) // запись строки в буфер
	fmt.Fprintf(&b, "World!") // Fprintf() принимает io.Writer
	// Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	// NewFile(fd uintptr, name string) *File
	// os.File выполняет контракт io.Writer
	b.WriteTo(os.Stdout)
}

func businessLogic(r io.Reader, w io.Writer) error {
	var data []byte      // все данные
	b := make([]byte, 8) // буфер
	for {
		n, err := r.Read(b) // считываем данные в буфер
		_ = n
		data = append(data, b...) // дополняем данные
		if err == io.EOF {
			break
		}
	}

	// какие-то действия с data

	_, err := w.Write(data) // записываем результат вычислений
	return err
}
