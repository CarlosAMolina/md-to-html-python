ROOT_PATHNAME=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
MAIN_PROJECT_ROOT_PATHNAME=$(ROOT_PATHNAME)/..
SCRIPT_CLONE_TOOLS_PATHNAME=$(ROOT_PATHNAME)/clone-tools-to-docker-volume
SCRIPT_OPEN_HTML_PATHNAME=$(ROOT_PATHNAME)/open-html
SCRIPT_SET_MEDIA_CONTENT=$(ROOT_PATHNAME)/set-media-content-to-docker-volume
DOCKER_NGINX_IMAGE_NAME=nginx:latest
DOCKER_VOLUME_NAME_NGINX_LOGS=nginx-logs
DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT=nginx-web-content
MEDIA_CONTENT_PATHNAME=$(HOME)/Software/cmoli-media-content
URL_NGINX_LOCALHOST=http://localhost:8080

create-web-content:
	# TODO continue here
	/bin/bash $(SCRIPT_CLONE_TOOLS_PATHNAME) \
		$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT)

create-web-server:
	make -f makefile-nginx \
		run \
		DOCKER_IMAGE_NAME=$(DOCKER_NGINX_IMAGE_NAME) \
		DOCKER_VOLUME_NGINX_WEB_CONTENT=$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT) \
		DOCKER_VOLUME_NGINX_LOGS=$(DOCKER_VOLUME_NAME_NGINX_LOGS) \
		URL_NGINX_LOCALHOST=$(URL_NGINX_LOCALHOST)

show-web:
	/bin/bash $(SCRIPT_OPEN_HTML_PATHNAME) \
		$(URL_NGINX_LOCALHOST) \
		$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT)

show-logs:
	make -f makefile-nginx \
		logs \
		DOCKER_VOLUME_NGINX_LOGS=$(DOCKER_VOLUME_NAME_NGINX_LOGS)

test: create-web-content \
	create-web-server \
	show-web

deploy: assert-required-file-updated-and-update-branch \
	clone-projects-content \
	clone-wiki \
	activate-docker-if-not-active \
	stop-containers \
	create-web-content

export-pandoc-template-html:
	# https://pandoc.org/MANUAL.html#option--print-default-template
	pandoc -D html > /tmp/template-default.html
