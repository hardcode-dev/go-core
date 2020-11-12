package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	baseURL := "http://localhost/search/"
	tokens := []string{"go", "programming", "metallica"}

	var mux sync.Mutex
	var errNum, counter, totalTime int64
	const N = 2000
	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		time.Sleep(time.Millisecond * 100)
		go func(i int) {
			defer wg.Done()
			url := baseURL + tokens[i%len(tokens)]
			t := time.Now()
			_, err := http.Get(url)
			if err != nil {
				errNum++
				return
			}
			elapsed := time.Since(t)
			mux.Lock()
			totalTime += elapsed.Milliseconds()
			counter++
			mux.Unlock()
		}(i % len(tokens))
	}

	wg.Wait()

	if counter > 0 {
		fmt.Printf("Среднее время выполнения запроса %v мс, ошибок выявлено: %d\n", totalTime/counter, errNum)
	}
}
