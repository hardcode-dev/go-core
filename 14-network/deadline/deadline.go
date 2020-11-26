package main

// Эхо-сервер. Без цензуры.

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// обработчик подключения
func handler(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(time.Second * 10))
	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(msg))
		_, err = conn.Write(msg)
		if err != nil {
			return
		}
	}
}

func main() {
	// регистрация сетевой службы
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handler(conn)
	}
}
