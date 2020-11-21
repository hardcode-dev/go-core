// Шаблон Fan-In - Fan-Out
// Шаблон демонстрирует технику распараллеливания вычислений с помощью разделения (Fan-Out)
// обработки значений, получаемых от генератора с помощью отдельных горутин.
// Далее происходит объединение вывода результатов обработки (Fan-In) в общий выходной поток.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// обработчик - этап конвейера, который обрабатывет данные, полученные от генератора
func processor(input <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range input {
			// симуляция продолжительных вычислений
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- i * i // возвращаем квадрат числа
		}
	}()
	return ch
}

// объединяем данные из массива каналов в один выходной канал (поток)
func fanIn(channels []<-chan int) <-chan int {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		go func(in <-chan int) {
			defer wg.Done()
			for i := range in {
				ch <- i
			}
		}(c)
	}
	// необходимо закрыть выходной канал, иначе deadlock
	// клинч возникнет между циклом range выше и точкой присоединения fanIn() к main()
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	// инициализируем исходные данные
	var in []int
	for i := 0; i < 10; i++ {
		in = append(in, i)
	}
	gen := make(chan int) // генератор
	go func() {
		defer close(gen)
		for i := 0; i < len(in); i++ {
			gen <- i
		}
	}()

	// собственно секция Fan-Out
	channels := make([]<-chan int, len(in))
	for i := 0; i < len(in); i++ {
		channels[i] = processor(gen)
	}

	ch := fanIn(channels)
	for i := range ch {
		fmt.Println("Recieved: ", i)
	}
}
