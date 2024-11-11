package main

import (
	"fmt"
	"sync"
)

func worker(countGorutines int) int {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	number := 2
	wg.Add(countGorutines)
	for i := 0; i < countGorutines; i++ {
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			number++
			fmt.Printf("Выполнилась %d горутина, результат %d\n", id, number)
			defer mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("Итоговое значение:", number)
	return number
}

func main() {
	countGorutines := 100
	worker(countGorutines)
}
