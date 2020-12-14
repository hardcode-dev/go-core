package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	msgQueue = make(chan string) // очередь поступающих сообщений от клиентов

	// каждому подключенному клиенту сопоставляется канал
	// для получения сообщений от остальных клиентов
	mux     sync.Mutex
	clients = make([]chan string, 0)
)

func main() {
	// регистрация точек API
	http.HandleFunc("/send", send)
	http.HandleFunc("/messages", messages)

	// отправка сообщений всем подключенным клиентам
	go publishMessages()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// приём сообщений от клиентов
func send(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("получено сообщение:", string(message))

	// все входящие сообщения пишутся в очередь,
	// дальше они обрабатываются в потоке publishMessages
	msgQueue <- string(message)
}

// получение сообщений от всех клиентов
func messages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// при подключении для клиента создаётся канал и добавляется в массив
	mux.Lock()
	client := make(chan string)
	clients = append(clients, client)
	mux.Unlock()

	// при отключении канал удаляется из массива, чтобы избежать паники.
	defer func() {
		mux.Lock()
		for i := range clients {
			if clients[i] == client {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
		mux.Unlock()
	}()

	// чтение сообщений из канала данного клиента
	for msg := range client {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return
		}
	}
}

// Отправка сообщений всем подключенным клиентам.
//
// Поступающие сообщения записываются в канал msgQueue,
// откуда они перенаправляются в каналы-клиенты.
// Шаблон Fan-Out.
func publishMessages() {
	for msg := range msgQueue {
		for _, c := range clients {
			c <- msg
		}
	}
}
