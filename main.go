package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the cmoli.es deployment CLI!")
	showHelp()
	// TODO rm below
	getVolumePath("foo")
	os.Exit(0)
	// TODO rm above
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
			os.Exit(0)
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

func testLocal() {
	fmt.Println("Testing local")
	pullGitCmoli()
	pullGitProjects()
	pullGitWiki()
	startDockerService()
	stopContainer("nginx-cmoli-container")
	removeVolume("nginx-web-content")
	removeVolume("pandoc")
	// Create Pandoc script for files
	pullDocker("python:3.8.15-alpine3.16")
	buildDockerPandoc()
	createVolume("nginx-web-content")
	createVolume("pandoc")
}

func pullGitCmoli() {
	run("git pull origin $(git branch --show-current)")
}

func pullGitProjects() {
	if exists("src/projects") {
		run("rm -rf src/projects")
	}
	run("mkdir src/projects")
	run("git clone --depth=1 --branch=main https://github.com/CarlosAMolina/checkIframe /tmp/checkIframe")
	run("mv /tmp/checkIframe/docs src/projects/check-iframe")
	run("rm -rf /tmp/checkIframe")
}

func pullGitWiki() {
	if exists("src/wiki") {
		run("rm -rf src/wiki")
	}
	run("git clone --depth=1 --branch=main https://github.com/CarlosAMolina/wiki /tmp/wiki")
	run("mv /tmp/wiki/src src/wiki")
	run("rm -rf /tmp/wiki")
}

func exists(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
