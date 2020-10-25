package main

import (
	"testing"
)

func Test_logMsg(t *testing.T) {
	l := new(MemLogger)
	val, err := calc(2, l, "msg")
	if err != nil {
		t.Fatal(err)
	}
	got := val
	want := 4
	if got != want {
		t.Fatalf("получено %d, ожидалось %d", got, want)
	}
}
