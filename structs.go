package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func showStructs() {
	p := Person{Name: "Rydam", Age: 20}
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
