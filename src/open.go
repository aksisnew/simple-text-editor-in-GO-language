package main

import (
	"fmt"
	"os"
)

func openFile() {
	fmt.Print("Enter file name: ")

	var filename string
	fmt.Scanln(&filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	currentFile = filename
	content = string(data)

	fmt.Println("File opened successfully.")
	fmt.Println()
	fmt.Println(content)
}