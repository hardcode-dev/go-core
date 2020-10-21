package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var errExample error

type myErr int

func (me myErr) Error() string {
	return "ошибка!"
}

var ErrME myErr

func main() {
	errExample = errors.New("пример создания ошибки с помощью пакета errors") // с маленькой буквы!
	fmt.Println(errExample.Error())
	errExample = fmt.Errorf("пример создания ошибки с помощью пакета fmt")
	fmt.Println(errExample) // почему без вызова метода Error()?

	val, err := envVar("1234")
	if err == ErrME {
		fmt.Println("My Error")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

// envVar возвращает переменную окружения, заданную по имени.
// Если переменная не найдена - возвращается ошибка.
func envVar(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return val, os.ErrInvalid
	}

	return val, nil
}

// Output: 	2020/10/02 13:23:36 переменная с именем неподходящее_имя_переменной не найдена
//			exit status 1
