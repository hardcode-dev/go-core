package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func worker(in, out chan int, err chan error) {
	for val := range in {
		if val%10 == 0 {
			err <- errors.New("ошибка:" + strconv.Itoa(val))
			continue
		}
		out <- val * val
	}
	close(out)
	close(err)
}

func main() {
	data := make(chan int)
	res := make(chan int)
	err := make(chan error)
	go worker(data, res, err)

	// генератор
	go func() {
		for i := 0; i < 12; i++ {
			data <- i
		}
		close(data)
	}()
	var wg sync.WaitGroup
	wg.Add(2)
	// обработка результатов и ошибок
	go func() {
		defer wg.Done()
		for e := range err {
			log.Println("ОШИБКА:", e.Error())
		}
	}()
	// обработка результатов и ошибок
	go func() {
		defer wg.Done()
		for val := range res {
			fmt.Println("Результат:", val)
		}
	}()
	wg.Wait()
	/*for {
		select {
		case val, ok := <-res:
			if !ok {
				break
			}
			fmt.Println("Результат:", val)
		case e, ok := <-err:
			if !ok {
				break
			}
			log.Println("ОШИБКА:", e.Error())
		}
	}*/
}
