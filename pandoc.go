package main

func buildDockerPandoc() {
	run(`docker build \
	-t python-create-pandoc-script \
	-f docker/Dockerfile-create-pandoc-script-for-files \
	--build-arg docker_image=python:3.8.15-alpine3.16 \
	--build-arg volume_nginx_web_content=nginx-web-content \
	--build-arg volume_pandoc=pandoc \
	.`)
}
