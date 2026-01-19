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
