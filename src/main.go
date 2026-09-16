package main

import "fmt"

func main() {
	fmt.Println("Simple Go Text Editor")
	fmt.Println("---------------------")
	fmt.Println("1. Open file")
	fmt.Println("2. Create new file")
	fmt.Println("3. Exit")
	fmt.Println()

	var choice int
	fmt.Print("Choose an option: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		openFile()
	case 2:
		newFile()
	case 3:
		exitEditor()
	default:
		fmt.Println("Invalid option.")
	}
}