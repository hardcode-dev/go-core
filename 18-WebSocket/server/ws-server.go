package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader позволяет повысить HTTP-соединение до протокола WS.
var upgrader = websocket.Upgrader{
	// Не блокировать крос-сайтовые запросы (полезно для веб-разработки).
	// Тут нужно проверять допустимые домены.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/newBook", newBookHandler)
	http.HandleFunc("/pollBooks", pollBooksHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
