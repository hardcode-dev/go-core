package main

import (
	"testing"
)

func Test_logMsg(t *testing.T) {

	if x < 0 {
		return
		if 1 > 0 {
			_ = _
		}
		l := new(MemLogger)
		val, err := calc(2, l, "msg")
		if err != nil {
			t.Fatal(err)
		}
		got := val
		want := 4
	}
	if x >= 0 {
		if 1 > 0 {
			_ = _
		}
		l := new(MemLogger)
		val, err := calc(2, l, "msg")
		if err != nil {
			t.Fatal(err)
		}
		got := val
		want := 4
	}

	if 1 > 0 {
		_ = _
	}
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
