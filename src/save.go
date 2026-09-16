package main

import (
	"fmt"
	"os"
)

func saveFile() {
	if currentFile == "" {
		fmt.Print("Enter file name to save: ")
		fmt.Scanln(&currentFile)
	}

	err := os.WriteFile(currentFile, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("File saved successfully.")
}