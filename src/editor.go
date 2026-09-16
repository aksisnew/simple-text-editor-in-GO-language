package main

import "fmt"

var currentFile string
var content string

func newFile() {
	currentFile = ""

	fmt.Println()
	fmt.Println("New file created.")
	fmt.Print("Enter text: ")

	var text string
	fmt.Scanln(&text)

	content = text

	fmt.Println("Text added.")
}