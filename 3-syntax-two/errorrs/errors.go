package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var errExample error

func main() {
	errExample = errors.New("пример создания ошибки с помощью пакета errors") // с маленькой буквы!
	fmt.Println(errExample.Error())
	errExample = fmt.Errorf("пример создания ошибки с помощью пакета fmt")
	fmt.Println(errExample) // почему без вызова метода Error()?

	val, err := envVar("неподходящее_имя_переменной")
	if err != nil {
		log.Fatal(err)
	}
	_ = val
}

// envVar возвращает переменную окружения, заданную по имени.
// Если переменная не найдена - возвращается ошибка.
func envVar(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return val, fmt.Errorf("переменная с именем %s не найдена", name)
	}

	return val, nil
}

// Output: 	2020/10/02 13:23:36 переменная с именем неподходящее_имя_переменной не найдена
//			exit status 1
