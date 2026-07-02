package main

import "fmt"

func showMaps() {
	person := map[string]string{
		"name":  "Rydam",
		"city":  "Toronto",
		"hobby": "coding",
	}

	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}
}
