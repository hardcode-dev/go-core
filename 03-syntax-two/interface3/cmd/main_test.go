package main

import (
	"go-core/3-syntax-two/interface3/pkg/storage"
	"testing"
)

func Test_usersNum(t *testing.T) {
	// в тестах используем заглушку в памяти
	s := new(storage.MemDB)
	got := usersNum(s)
	want := 1
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
