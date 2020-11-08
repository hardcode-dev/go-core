package main

import (
	"net/http"
	_ "net/http/pprof" // импорт только ради побочного эффекта
)

func main() {
	go increment()
	http.ListenAndServe(":80", nil)
}

func increment() {
	var i int64
	for {
		i++
	}
}

// go tool pprof -web http://localhost/debug/pprof/profile
// go tool pprof -web http://localhost/debug/pprof/heap
