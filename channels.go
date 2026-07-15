package main

import "fmt"

func sendData(ch chan int, val int) {
	ch <- val
}

func showChannels() {
	ch := make(chan int, 3)
	go sendData(ch, 10)
	go sendData(ch, 20)
	go sendData(ch, 30)

	for i := 0; i < 3; i++ {
		fmt.Println("Received:", <-ch)
	}
}
