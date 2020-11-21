// Пример использования sync.Pool
//
// В примере создаётся пул ресурсов - объектов в памяти размером 1МБ. Пул инициализируется (прогревается) четырьмя элементами.
// Далее в примере горутины конкурентно требуют ресурсы из пула 10 000 раз. Теоретически это может потребовать ~10GB памяти.
// Однако за счёт возврата ресурсов в пул и переиспользования требуется примерно 5 - 20 экземпляров.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mux sync.Mutex
	var counter int
	var pool sync.Pool
	// конструктор нового элемента пула
	pool.New = func() interface{} {
		mux.Lock()
		counter++ // подсчитаем общее (максимальное) количество созданных ресурсов
		mux.Unlock()
		// создам элемент набора - выдедяем 1024Kb памяти
		var mem = make([]byte, 1024*1024)
		return &mem
	}

	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())

	const N = 10000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			memory := pool.Get().(*[]byte)
			defer pool.Put(memory)
		}()

	}
	wg.Wait()

	fmt.Printf("Total number of pool resources created: %d\n", counter)
}
