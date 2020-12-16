package textproto

import (
	"io"
	"net"
	"time"
)

// Handler - обработчик для текстового протокола.
func Handler(conn net.Conn) error {

	defer conn.Close()

	err := conn.SetDeadline(time.Now().Add(time.Second * 120))
	if err != nil {
		return err
	}

	for {
		// читаем сообщение от клиента
		buf := make([]byte, 70000)
		size, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		message := buf[:size]

		// делаем что-то с message (массив байт пакета)
		_ = message

		response := "response"

		_, err = conn.Write([]byte(response))
		if err != nil {
			return err
		}
	}
}
