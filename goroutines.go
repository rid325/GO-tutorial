package main

import (
	"fmt"
	"sync"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

func showGoroutines() {
	var wg sync.WaitGroup

	messages := []string{"Hello from goroutine 1", "Hello from goroutine 2", "Hello from goroutine 3"}

	for _, msg := range messages {
		wg.Add(1)
		go printMessage(msg, &wg)
	}

	wg.Wait()
}
