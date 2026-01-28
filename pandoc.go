package main

import "path/filepath"

// To export pandoc template html: pandoc -D html > /tmp/template-default.html
// https://pandoc.org/MANUAL.html#option--print-default-template

// Connect: docker exec -it python-create-pandoc-script-container /bin/sh
func buildDockerCreatePandocScript() {
	run(`docker build \
	-t python-create-pandoc-script \
	-f docker/Dockerfile-create-pandoc-script-for-files \
	--build-arg docker_image=python:3.8.15-alpine3.16 \
	--build-arg volume_nginx_web_content=nginx-web-content \
	--build-arg volume_pandoc=pandoc \
	.`)
}

func buildDockerImagePandoc() {
	run(`docker build \
	-t pandoc-convert-md-to-html \
	--build-arg docker_image=pandoc/minimal:2.17-alpine \
	--build-arg volume_pandoc=pandoc \
	-f md-to-html/Dockerfile-convert-md-to-html-for-files \
	.`)
}

func copyContentToVolumePandoc() {
	mdPath := filepath.Join(getPathSoftware(), "cmoli.es", "deploy", "md-to-html")
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
	for isContainerRunning("pandoc-convert-md-to-html-container") {
		sleep(4)
	}
}
