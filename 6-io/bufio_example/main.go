package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b []byte
	buf := bytes.NewBuffer(b) // буфер, выполняет контракты io.Reader и io.Writer
	reader := bufio.NewReader(buf)
	writer := bufio.NewWriter(buf)
	writer.Write([]byte("Hello, World!")) // запись строки в буфер
	writer.Flush()
	businessLogic(*reader, os.Stdout)
}

func businessLogic(r bufio.Reader, w io.Writer) error {
	bytes, _, err := r.ReadLine()
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
