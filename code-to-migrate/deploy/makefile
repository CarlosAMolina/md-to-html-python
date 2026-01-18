ROOT_PATHNAME=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
MAIN_PROJECT_ROOT_PATHNAME=$(ROOT_PATHNAME)/..
SCRIPT_CONVERT_MD_TO_HTML_PATHNAME=$(ROOT_PATHNAME)/md-to-html/convert-md-to-html-for-files
SCRIPT_CLONE_TOOLS_PATHNAME=$(ROOT_PATHNAME)/clone-tools-to-docker-volume
SCRIPT_MODIFY_HTML_FILES=$(ROOT_PATHNAME)/modify-html
SCRIPT_OPEN_HTML_PATHNAME=$(ROOT_PATHNAME)/open-html
SCRIPT_SET_MEDIA_CONTENT=$(ROOT_PATHNAME)/set-media-content-to-docker-volume
DOCKER_NGINX_IMAGE_NAME=nginx:latest
DOCKER_PANDOC_IMAGE_NAME=pandoc/minimal:2.17-alpine
DOCKER_PYTHON_IMAGE_NAME=python:3.8.15-alpine3.16
DOCKER_VOLUME_NAME_NGINX_LOGS=nginx-logs
DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT=nginx-web-content
DOCKER_VOLUME_NAME_PANDOC=pandoc
MEDIA_CONTENT_PATHNAME=$(HOME)/Software/cmoli-media-content
URL_NGINX_LOCALHOST=http://localhost:8080

create-web-content:
	/bin/bash remove-volumes \
		$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT) \
		$(DOCKER_VOLUME_NAME_PANDOC)
	/bin/bash $(SCRIPT_CONVERT_MD_TO_HTML_PATHNAME) \
		$(DOCKER_PANDOC_IMAGE_NAME) \
		$(DOCKER_PYTHON_IMAGE_NAME) \
		$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT) \
		$(DOCKER_VOLUME_NAME_PANDOC)
	/bin/bash $(SCRIPT_MODIFY_HTML_FILES) \
		$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT)
	/bin/bash $(SCRIPT_SET_MEDIA_CONTENT) \
		$(DOCKER_VOLUME_NAME_NGINX_WEB_CONTENT) \
		$(MEDIA_CONTENT_PATHNAME)
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

stop-containers:
	make -f makefile-nginx \
	    stop-if-running

test: stop-containers \
	create-web-content \
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
