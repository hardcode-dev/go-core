package main

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
	// регистрация сетевой службы
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
