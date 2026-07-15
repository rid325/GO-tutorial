package main

import (
	"fmt"
	"os"
)

func writeFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func appendToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

func showFileWrite() {
	err := writeFile("output.txt", "Hello from Go!\n")
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")

	err = appendToFile("output.txt", "Appended line!\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Appended to file successfully")
}
