package main

import (
	"testing"
)

func Test_logMsg(t *testing.T) {
	l := new(MemLogger)
	err := logMsg(l, "msg")
	if err != nil {
		t.Fatal(err)
	}
}
