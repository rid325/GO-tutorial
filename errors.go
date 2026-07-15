package main

import (
	"errors"
	"fmt"
)

// ErrDivByZero is a sentinel error for division by zero
var ErrDivByZero = errors.New("cannot divide by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func showErrors() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// test divide by zero
	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Caught error:", err)
	}

	// use errors.Is to check sentinel
	if errors.Is(err, ErrDivByZero) {
		fmt.Println("Confirmed: divide by zero error")
	}
}
