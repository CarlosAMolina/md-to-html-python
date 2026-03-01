package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

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
			deploy()
			fmt.Println("Deployed! :)")
			os.Exit(0)
		case "2":
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
	fmt.Println("1. Deploy")
	fmt.Println("2. Testing local")
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func deploy() {
	pullGitRepos()
	startDockerService()
	stopContainer("nginx-cmoli-container")
	removeVolume("nginx-web-content")
	removeVolume("pandoc")
	// Create Pandoc script for files
	pullDocker("python:3.8.15-alpine3.16")
	buildDockerCreatePandocScript()
	createVolume("nginx-web-content")
	createVolume("pandoc")
	copyContentToVolumeNginx()
	copyContentToVolumePandoc()
	runDockerCreatePandocScript()
	pullDockerPandoc()
	buildDockerImagePandoc()
	runDockerPandoc()
	modifyHtml()
	copyMediaToDockerVolume()
}

func testLocal() {
	deploy()
	run("firefox " + getVolumePath("nginx-web-content") + "/index.html")
}

func getPathSoftware() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, "Software")
}

func pullGitTools() {
	repoNames := [3]string{"open-urls", "job-check-lambda-name", "job-modify-issue-name"}
	for i := range len(repoNames) {
		repoName := repoNames[i]
		pullGitRepo(repoName)
	}
}

func exists(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func copyMediaToDockerVolume() {
	volumePath := getVolumePath("nginx-web-content")
	run("cp -r ~/Software/cmoli-media-content/* " + volumePath)
	videoVolumePath := filepath.Join(volumePath, "felices-fiestas/src/movie.mp4")
	run("rm " + videoVolumePath)
	run("ln -s ~/Software/cmoli-media-content/felices-fiestas/src/movie.mp4 " + videoVolumePath)
}
