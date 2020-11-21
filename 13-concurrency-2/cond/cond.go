// Пример использования sync.Cond и sync.Once.
//
// sync.Cond представляет собой блок и событие, которое можно использовать как управляющий сигнал для горутин.
// Содержит мюьтекс, который блокирует при вызове метода Wait до наступления события.
// При этом блок необходимо предварительно установить явно, метод Wait блок не устанавливает.
// sync.Once используется для выполнения функции-аргумента строго один раз.
// Считается только количество вызовов once.Do. Изменение аргумента результата не даёт.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var once sync.Once
	var condition = sync.NewCond(&sync.Mutex{})
	var counter int
	var wg sync.WaitGroup
	wg.Add(5)
	incr := func() { counter++ }
	decr := func() { counter-- }

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			condition.L.Lock()
			defer condition.L.Unlock()
			once.Do(incr)
			fmt.Printf("Increment done. Goroutine #%d waits for event.\n", i)
			condition.Wait()
			fmt.Printf("Recieved condition event. Goroutine #%d exits.\n", i)
		}(i)
	}

	// Ожидаем выполнение всех горутин. Конечно, лучше использовать каналы
	time.Sleep(3 * time.Second)

	condition.Broadcast()

	wg.Wait()

	fmt.Printf("Counter is: %d\n", counter)

	once.Do(decr)

	fmt.Printf("After 'decr' counter is: %d\n", counter)
}
