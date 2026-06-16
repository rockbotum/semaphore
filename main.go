package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := NewSemaphore(3)

	var wg sync.WaitGroup
	taskCount := 10

	var completed int
	var mu sync.Mutex

	for i := 1; i <= taskCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem.Acquire()
			defer sem.Release()

			fmt.Printf("Горутина %d: начала работу, свободно место: %d\n",
				id, sem.Available())
			time.Sleep(100 * time.Microsecond)

			mu.Lock()
			completed++
			mu.Unlock()

			fmt.Printf("Горутина %d: завершила работу\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Printf("Все %d задач выполнено\n", completed)
}
