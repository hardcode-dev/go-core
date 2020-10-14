package textproto

import (
	"bufio"
	"io"
	"net"
)

// Handler - обработчик для текстового протокола.
func Handler(conn net.Conn) error {
	defer conn.Close()         // соединение надо закрыть в конце что бы ни случилось
	r := bufio.NewReader(conn) // *Reader
	w := bufio.NewWriter(conn) // *Writer
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil // конец потока от клиента
			}
			return err
		}
		_ = str                  // обрабатываем поступившие данные
		response := []byte("OK") // ответ клиенту
		w.Write(response)        // пишем ответ в сетевое подключение
		w.Flush()                // сбрасываем буфер
	}
}
