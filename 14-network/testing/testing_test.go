package main

import (
	"io"
	"io/ioutil"
	"net"
	"strings"
	"sync"
	"testing"
)

func Test_handler(t *testing.T) {
	srv, cl := net.Pipe()
	var wg sync.WaitGroup
	wg.Add(1)

	var msg string
	go func() {
		b, err := ioutil.ReadAll(cl)
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}
		msg = string(b)
		cl.Close()
		wg.Done()
	}()

	handler(srv)
	wg.Wait()
	if !strings.Contains(msg, "2020") {
		t.Fatalf("неверное сообщение от сервера: %s", msg)
	}
}
