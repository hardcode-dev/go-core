package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:13")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ответ от сервера:", string(msg))
}
