package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

type book struct {
	Title  string
	Author string
}

var books = []book{
	{
		Title:  "The Lord Of The Rings",
		Author: "J.R.R. Tolkien",
	},
}

// возвращает список книг
func booksHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	jsonData, err := json.Marshal(books)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	conn.WriteMessage(websocket.TextMessage, jsonData)
}

// поддердживает соединение и возвращает список книг по требованию
func pollBooksHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			return
		}

		switch string(message) {
		case "books":
			jsonData, err := json.Marshal(books)
			if err != nil {
				conn.WriteMessage(mt, []byte(err.Error()))
				return
			}
			conn.WriteMessage(websocket.TextMessage, jsonData)
		default:
			conn.WriteMessage(websocket.TextMessage, []byte("return"))
			return
		}
	}
}

// добавляет книгу в список
func newBookHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	mt, message, err := conn.ReadMessage()
	if err != nil {
		conn.WriteMessage(mt, []byte(err.Error()))
		return
	}

	var b book
	err = json.NewDecoder(bytes.NewBuffer(message)).Decode(&b)
	if err != nil {
		conn.WriteMessage(mt, []byte(err.Error()))
		return
	}
	books = append(books, b)
}
