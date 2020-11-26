package main

// Сервер текущего времени в соответсвии с RFC 867.

import (
	"log"
	"net"
	"time"
)

// обработчик подключения
func handler(conn net.Conn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}

func main() {
	// регистрация сетевой службы на всех сетевых интерфейсах на порту 13
	listener, err := net.Listen("tcp4", ":13")
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
