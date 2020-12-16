package main

import "testing"

type memInfo int

func (mi *memInfo) Info() string {
	return "42"
}

func Test_printInfo(t *testing.T) {
	var mi memInfo
	printInfo(&mi)
}
