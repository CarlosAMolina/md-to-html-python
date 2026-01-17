package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to the cmoli.es deployment CLI!")
	showHelp()
	testLocal() // TODO rm
	os.Exit(0)  // TODO rm
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
		case "3":
			testLocal()
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

func showHelp() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Full deployment. Run all steps")
	fmt.Println("2. Generate web content. Convert Markdown to HTML")
	fmt.Println("3. Testing local")
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func run(command string) {
	fmt.Println(command)
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func testLocal() {
	fmt.Println("Testing local")
	gitPull()
}

func gitPull() {
	run("git pull origin $(git branch --show-current)")
}
