package main

import (
	"fmt"
	"go-core/3-syntax-two/interface3/pkg/storage"
)

func main() {
	// в рабочем коде используем конкретную реализацию в БД
	s := storage.NewPgDB()
	num := usersNum(s)
	fmt.Printf("В системе зарегистрировано %d пользователей.\n", num)
}

func usersNum(s storage.Interface) int {
	return len(s.Users())
}
