package main

import (
	"fmt"
	"os"
)

func writeFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func showFileWrite() {
	err := writeFile("output.txt", "Hello from Go!\n")
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")
}
