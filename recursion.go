package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func showRecursion() {
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Fibonacci of 8:", fibonacci(8))
}
