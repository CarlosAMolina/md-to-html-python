package main

import (
	"fmt"
	"os"
)

func showHelp() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Full deployment. Run all steps")
	fmt.Println("2. Generate web content. Convert Markdown to HTML")
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func main() {
	fmt.Println("Welcome to the cmoli.es deployment CLI!")
	showHelp()

	var choice string

	for {
		fmt.Print(">> ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			fmt.Println("Starting full deployment")
			os.Exit(0)
		case "2":
			fmt.Println("Generating web content")
			os.Exit(0)
		case "e":
			fmt.Println("Bye!")
		case "h":
			showHelp()
		default:
			fmt.Println("Invalid input")
		}
	}
}
