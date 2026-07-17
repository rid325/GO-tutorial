package main

import "fmt"

func increment(n *int) {
	*n++
}

func showPointers() {
	x := 10
	fmt.Println("Before:", x)
	increment(&x)
	fmt.Println("After:", x)
}
