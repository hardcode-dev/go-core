package main

import (
	"fmt"
	"log"

	"go-core/1-intro/crawler/pkg/spider"
)

func main() {
	const url = "https://habr.com"
	data, err := spider.Scan(url, 2)
	if err != nil {
		log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
	}

	for k, v := range data {
		fmt.Printf("Страница %s имеет адрес: %s\n", v, k)
	}
}
