package main

import "fmt"

func main() {
	for {
		fmt.Println()
		fmt.Println("Simple Go Text Editor")
		fmt.Println("---------------------")
		fmt.Println("1. Open file")
		fmt.Println("2. Create new file")
		fmt.Println("3. Save file")
		fmt.Println("4. Change extension")
		fmt.Println("5. Exit")
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
			saveFile()
		case 4:
			changeExtension()
		case 5:
			exitEditor()
		default:
			fmt.Println("Invalid option.")
		}
	}
}