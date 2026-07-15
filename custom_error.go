package main

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistic value"}
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "cannot be empty"}
	}
	if len(name) > 50 {
		return &ValidationError{Field: "name", Message: "too long, max 50 chars"}
	}
	return nil
}
