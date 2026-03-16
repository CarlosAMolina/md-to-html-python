package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	startDockerService()
	stopContainer("nginx-cmoli-container")
	removeVolume("nginx-web-content")
	removeVolume("pandoc")
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
	compareResults()
}

func getPathSoftware() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, "Software")
}

func exists(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func compareResults() {
	volumePath := getVolumePath("nginx-web-content")
	run("meld " + volumePath + " testdata/dir-converted/")
}
