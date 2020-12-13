package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// получение сообщений в фоне
	go messages()
	// интерактивная отправка сообщений в основном потоке
	send()
}

func send() {
	reader := bufio.NewReader(os.Stdin) // буфер для os.Stdin
	for {
		fmt.Print("-> ")
		msg, _ := reader.ReadString('\n')        // чтение строки (до символа перевода)
		msg = strings.Replace(msg, "\n", "", -1) // удаление перевода строки

		ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/send", nil)
		if err != nil {
			ws.Close()
			log.Fatalf("не удалось подключиться к серверу: %v", err)
		}

		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			ws.Close()
			log.Fatalf("не удалось отправить сообщение: %v", err)
		}
		ws.Close()
	}
}

func messages() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/messages", nil)
	if err != nil {
		log.Fatalf("не удалось подключиться к серверу: %v", err)
	}
	defer ws.Close()

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Fatalf("не удалось прочитать сообщение: %v", err)
		}
		log.Printf("Сообщение от сервера: %s.\n", string(p))
		time.Sleep(time.Second)
	}
}
