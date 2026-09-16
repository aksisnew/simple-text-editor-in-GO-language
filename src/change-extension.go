package main

import (
	"fmt"
	"path/filepath"
)

func changeExtension() {
	if currentFile == "" {
		fmt.Println("No file is currently open.")
		return
	}

	fmt.Print("Enter new extension: ")

	var extension string
	fmt.Scanln(&extension)

	if extension[0] != '.' {
		extension = "." + extension
	}

	newFileName := filepath.Base(currentFile)
	newFileName = newFileName[:len(newFileName)-len(filepath.Ext(newFileName))]
	newFileName += extension

	currentFile = newFileName

	fmt.Println("File extension changed to:", currentFile)
}