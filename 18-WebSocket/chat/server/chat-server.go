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
	msgQueue = make(chan string)

	mux     sync.Mutex
	clients = make([]chan string, 0)
)

func main() {
	http.HandleFunc("/send", send)
	http.HandleFunc("/messages", messages)
	go publishMessages()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

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
	msgQueue <- string(message)
}

func messages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	mux.Lock()
	client := make(chan string)
	defer close(client)
	clients = append(clients, client)
	mux.Unlock()

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

	for msg := range client {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			return
		}
	}
}

func publishMessages() {
	for msg := range msgQueue {
		for _, c := range clients {
			c <- msg
		}
	}
}
