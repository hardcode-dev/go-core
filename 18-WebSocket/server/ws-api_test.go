// +build !arm

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func Test_booksHandler(t *testing.T) {
	mux := http.DefaultServeMux
	http.HandleFunc("/books", booksHandler)
	server := httptest.NewServer(mux)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/books"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("не удалось подключиться к серверу ws на %s: %v", wsURL, err)
	}
	defer ws.Close()

	mt, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("не удалось прочитать сообщение: %v", err)
	}
	if mt != websocket.TextMessage {
		t.Fatalf("получен тип сообщения %d, ожидался %d", mt, websocket.TextMessage)
	}

	t.Logf("Сообщение от сервера: %s. Тип сообщения: %d.\n", string(p), mt)
}

func Test_pollBooksHandler(t *testing.T) {
	mux := http.DefaultServeMux
	http.HandleFunc("/pollBooks", pollBooksHandler)

	server := httptest.NewServer(mux)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/pollBooks"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("не удалось подключиться к серверу ws на %s: %v", wsURL, err)
	}
	defer ws.Close()

	for i := 0; i < 4; i++ {
		msg := []byte("books")
		if i > 2 {
			msg = []byte("quit")
		}
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			t.Fatalf("не удалось отправить сообщение: %v", err)
		}

		mt, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("не удалось прочитать сообщение: %v", err)
		}
		if mt != websocket.TextMessage {
			t.Fatalf("получен тип сообщения %d, ожидался %d", mt, websocket.TextMessage)
		}
		t.Logf("Сообщение от сервера: %s. Тип сообщения: %d.\n", string(p), mt)
		time.Sleep(time.Second)
	}
}

func Test_newBookHandler(t *testing.T) {
	// тест создания новой книжки
	mux := http.DefaultServeMux
	http.HandleFunc("/newBook", newBookHandler)
	server := httptest.NewServer(mux)
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/newBook"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("не удалось подключиться к серверу ws на %s: %v", wsURL, err)
	}
	defer ws.Close()

	item := book{
		Author: "George Orwell",
		Title:  "1984",
	}
	b, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = ws.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		t.Fatalf("не удалось отправить сообщение: %v", err)
	}
	ws.Close()

	// проверка успешности создания - чтение списка книг
	http.HandleFunc("/books", booksHandler)
	wsURL = "ws" + strings.TrimPrefix(server.URL, "http") + "/books"

	ws, _, err = websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("не удалось подключиться к серверу ws на %s: %v", wsURL, err)
	}
	defer ws.Close()

	mt, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("не удалось прочитать сообщение: %v", err)
	}
	if mt != websocket.TextMessage {
		t.Fatalf("получен тип сообщения %d, ожидался %d", mt, websocket.TextMessage)
	}

	t.Logf("Сообщение от сервера: %s. Тип сообщения: %d.\n", string(p), mt)
}
