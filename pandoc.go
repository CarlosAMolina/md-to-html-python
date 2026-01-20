package main

// Connect: docker exec -it python-create-pandoc-script-container /bin/sh
func buildDockerPandoc() {
	run(`docker build \
	-t python-create-pandoc-script \
	-f docker/Dockerfile-create-pandoc-script-for-files \
	--build-arg docker_image=python:3.8.15-alpine3.16 \
	--build-arg volume_nginx_web_content=nginx-web-content \
	--build-arg volume_pandoc=pandoc \
	.`)
}

func copyContentToVolumePandoc() {
	volumePath := getVolumePath("pandoc")
	run("cp -r md-to-html/create-pandoc-script-for-files " + volumePath)
	run("cp -r md-to-html/pandoc-config " + volumePath)
	run("cp md-to-html/convert-md-to-html " + volumePath)
	run("cp md-to-html/run-create-pandoc-script-for-files " + volumePath)
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
