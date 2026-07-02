package main

import "fmt"

func showArrays() {
	fruits := []string{"apple", "banana", "cherry"}

	for i, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", i, fruit)
	}
}
