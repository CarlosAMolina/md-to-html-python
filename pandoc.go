package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// To export pandoc template html: pandoc -D html > /tmp/template-default.html
// https://pandoc.org/MANUAL.html#option--print-default-template

// Connect: docker exec -it python-create-pandoc-script-container /bin/sh
func buildDockerCreatePandocScript() {
	image := "python-create-pandoc-script"
	if existsImage(image) {
		fmt.Println("No build is required: " + image)
		return
	}
	command := `docker build \
	-t {image} \
	-f {dockerfile} \
	--build-arg docker_image=python:3.8.15-alpine3.16 \
	--build-arg volume_nginx_web_content=nginx-web-content \
	--build-arg volume_pandoc=pandoc \
	{buildDir}`
	command = strings.ReplaceAll(command, "{image}", image)
	buildDir := filepath.Join(getPathSoftware(), "cmoli.es-deploy")
	// Fixed builDir to not use random value if the executable runs in a differente directory.
	command = strings.ReplaceAll(command, "{buildDir}", buildDir)
	dockerfile := filepath.Join(buildDir, "docker/Dockerfile-create-pandoc-script-for-files")
	command = strings.ReplaceAll(command, "{dockerfile}", dockerfile)
	run(command)
}

func buildDockerImagePandoc() {
	image := "pandoc-convert-md-to-html"
	if existsImage(image) {
		fmt.Println("No build is required: " + image)
		return
	}
	command := `docker build \
	-t {image} \
	-f {dockerfile} \
	--build-arg docker_image=pandoc/minimal:2.17-alpine \
	--build-arg volume_pandoc=pandoc \
	{buildDir}`
	command = strings.ReplaceAll(command, "{image}", image)
	buildDir := filepath.Join(getPathSoftware(), "cmoli.es-deploy")
	// Fixed builDir to not use random value if the executable runs in a differente directory.
	command = strings.ReplaceAll(command, "{buildDir}", buildDir)
	dockerfile := filepath.Join(buildDir, "md-to-html/Dockerfile-convert-md-to-html-for-files")
	command = strings.ReplaceAll(command, "{dockerfile}", dockerfile)
	run(command)
}

func copyContentToVolumePandoc() {
	mdPath := filepath.Join(getPathSoftware(), "cmoli.es-deploy", "md-to-html")
	volumePath := getVolumePath("pandoc")
	run("cd " + mdPath + " && cp -r create-pandoc-script-for-files " + volumePath)
	run("cd " + mdPath + " && cp -r pandoc-config " + volumePath)
	run("cd " + mdPath + " && cp convert-md-to-html " + volumePath)
	run("cd " + mdPath + " && cp run-create-pandoc-script-for-files " + volumePath)
}

func pullDockerPandoc() {
	pullDocker("pandoc/minimal:2.17-alpine")
}

func runDockerCreatePandocScript() {
	run(`docker run \
		-it \
		--rm \
		-d \
		--name python-create-pandoc-script-container \
		--mount type=volume,source=pandoc,target=/pandoc \
		--mount type=volume,source=nginx-web-content,target=/nginx-web-content \
		python-create-pandoc-script`)
	for isContainerRunning("python-create-pandoc-script-container") {
		sleep(3)
	}
	scriptPath := getVolumePath("pandoc") + "/run-on-files-convert-md-to-html"
	if !exists(scriptPath) {
		panic("The pandoc script does not exist: " + scriptPath)
	}
	run("chmod +x " + scriptPath)
}

func runDockerPandoc() {
	run(`docker run \
		-it \
		--rm \
		-d \
		--name pandoc-convert-md-to-html-container \
		--mount type=volume,source=pandoc,target=/pandoc \
		--mount type=volume,source=nginx-web-content,target=/nginx-web-content \
		pandoc-convert-md-to-html`)
	fmt.Println("Waiting Pandoc to finish")
	for isContainerRunning("pandoc-convert-md-to-html-container") {
		sleep(5)
	}
}
