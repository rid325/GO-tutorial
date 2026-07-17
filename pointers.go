package main

import "fmt"

func increment(n *int) {
	*n++
}

func double(n *int) {
	*n *= 2
}

func showPointers() {
	x := 10
	fmt.Println("Before:", x)
	increment(&x)
	fmt.Println("After increment:", x)
	double(&x)
	fmt.Println("After double:", x)
}
