package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func showClosures() {
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	addFive := makeAdder(5)
	fmt.Println(addFive(3))
	fmt.Println(addFive(10))
}
