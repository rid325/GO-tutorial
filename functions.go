package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func showFunctions() {
	result := add(3, 7)
	fmt.Printf("3 + 7 = %d\n", result)
}
