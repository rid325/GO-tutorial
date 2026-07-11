package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string { return "Woof!" }
func (c Cat) Speak() string { return "Meow!" }

func showInterfaces() {
	animals := []Animal{Dog{}, Cat{}}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}
