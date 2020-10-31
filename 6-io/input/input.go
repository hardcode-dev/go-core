package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // буфер для os.Stdin
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')         // чтение строки (до символа перевода)
		text = strings.Replace(text, "\n", "", -1) // удаление перевода строки
		fmt.Println("echo:", text)
	}
}
