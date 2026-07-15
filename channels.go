package main

import "fmt"

func sendData(ch chan int) {
	ch <- 42
}

func showChannels() {
	ch := make(chan int)
	go sendData(ch)
	val := <-ch
	fmt.Println("Received:", val)
}
