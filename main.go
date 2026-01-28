package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	testLocal()
	os.Exit(0) // TODO rm above
	fmt.Println("Welcome to the cmoli.es deployment CLI!")
	showHelp()
	var choice string
	for {
		fmt.Print(">> ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			fmt.Println("Starting full deployment")
			// TODO:
			//assert-required-file-updated-and-update-branch \
			//clone-projects-content \
			//clone-wiki \
			//activate-docker-if-not-active \
			//stop-containers \
			//create-web-content
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
	pullGitRepo("cmoli.es")
	pullGitRepo("cmoli.es-deploy")
	pullGitRepo("checkIframe")
	pullGitRepo("wiki")
	pullGitTools()
	startDockerService()
	// TODO add a menu option to stop the testing web container.
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
	pullDockerNginx()
	buildDockerImageNginx()
	runDockerNginx()
	openWeb()
}

func pullGitRepo(repo string) {
	repoPath := filepath.Join(getPathSoftware(), repo)
	if exists(repoPath) {
		run("cd " + repoPath + " && git pull origin $(git branch --show-current)")
	} else {
		run("git clone --depth=1 --branch=main https://github.com/CarlosAMolina/" + repo + " " + repoPath)
	}
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
	path := getVolumePath("nginx-web-content")
	run("cp -r ~/Software/cmoli-media-content/* " + path)
}
