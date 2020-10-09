package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// fibo рекурсивно вычисляет числа фибоначчи, используя идею динамического программирования
func fibo(n int, cache map[int]int) int {
	// проверка наличия уже вычисленного результата в кэше
	if val, ok := cache[n]; ok {
		return val
	}
	// базовая часть рекурсии
	if n < 2 {
		return 1
	}
	// рекурсивный шаг с записью в кэш
	num := fibo(n-1, cache) + fibo(n-2, cache)
	cache[n] = num
	return num
}

func main() {
	cache := make(map[int]int)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("N: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		text = strings.TrimSuffix(text, "\r\n")
		text = strings.TrimSuffix(text, "\n")
		n, err := strconv.Atoi(text)
		if err != nil {
			log.Println(err)
		}
		num := fibo(n, cache)
		fmt.Printf("Число Фибоначчи №%d = %d\n", n, num)
	}
}
