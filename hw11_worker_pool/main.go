package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := NewCounter()
	countGoroutines := 100
	wg := sync.WaitGroup{}
	wg.Add(countGoroutines)

	for i := 0; i < countGoroutines; i++ {
		go func() {
			currentValue := counter.Inc()

			fmt.Println("Current value: ", currentValue)

			wg.Done()
		}()
	}

	wg.Wait()
}
