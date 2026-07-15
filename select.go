package main

import (
	"fmt"
	"time"
)

func showSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Millisecond)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
		}
	}
}
