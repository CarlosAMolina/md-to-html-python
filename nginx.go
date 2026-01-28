package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

// TODO the `md-to-html` folder must only convert content, the responsability for copying web content must be outside this folder.
func copyContentToVolumeNginx() {
	// cmoli.es
	cmoliPath := filepath.Join(getPathSoftware(), "cmoli.es", "src")
	volumePath := getVolumePath("nginx-web-content")
	run("cp -r " + cmoliPath + "/* " + volumePath)
	// check-iframe
	checkIframePath := filepath.Join(getPathSoftware(), "checkIframe", "docs")
	checkIframePathInVolume := filepath.Join(volumePath, "projects", "check-iframe")
	run("mkdir " + checkIframePathInVolume)
	run("cp -r " + checkIframePath + "/* " + checkIframePathInVolume)
	// wiki
	wikiPath := filepath.Join(getPathSoftware(), "wiki", "src")
	wikiPathInVolume := filepath.Join(volumePath, "wiki")
	run("mkdir " + wikiPathInVolume)
	run("cp -r " + wikiPath + "/* " + wikiPathInVolume)
	// tools
	toolNames := [3]string{"open-urls", "job-check-lambda-name", "job-modify-issue-name"}
	toolsPathInVolume := filepath.Join(volumePath, "tools")
	for i := range len(toolNames) {
		toolRepo := toolNames[i]
		toolPath := filepath.Join(getPathSoftware(), toolRepo)
		run("cp -r " + toolPath + " " + toolsPathInVolume)
		run("rm -rf " + filepath.Join(toolsPathInVolume, toolRepo, ".git"))
	}
}

// TODO add to all dockers: if imageExist, not download
func pullDockerNginx() {
	pullDocker("nginx:latest")
}

// TODO add to all dockers: if imageExist, not build
func buildDockerImageNginx() {
	run(`docker build \
		-t nginx-cmoli \
		-f docker/Dockerfile-nginx \
		--build-arg docker_image=nginx:latest \
	.`)
}

// show logs: tail -f $(path of volume ngingx-logs)/access.log
func runDockerNginx() {
	run(`docker run \
		-it \
		--rm \
		-d \
		-p 8080:80 \
		--name nginx-cmoli-container \
		--mount type=volume,source=nginx-logs,target=/var/log/nginx \
		--mount type=volume,source=nginx-web-content,target=/usr/share/nginx/html,readonly \
		nginx-cmoli`)
	for !isNginxListening() {
		fmt.Println("Waiting for Nginx to be ready")
		sleep(1)
	}
	filePath := getVolumePath("nginx-web-content") + "/index.html"
	for !exists(filePath) {
		fmt.Println("The file " + filePath + " does not exist. Retrying again")
		sleep(2)
	}
}

func isNginxListening() bool {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		return false
	}
	resp.Body.Close()
	return true
}

func openWeb() {
	run("firefox http://localhost:8080")
}
