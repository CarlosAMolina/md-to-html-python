ROOT_PATHNAME=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# TODO contine here
create-web-server:
	make -f makefile-nginx \
		run \
		DOCKER_IMAGE_NAME=nginx:latest \
		DOCKER_VOLUME_NGINX_WEB_CONTENT=nginx-web-content \
		DOCKER_VOLUME_NGINX_LOGS=nginx-logs \
		URL_NGINX_LOCALHOST=http://localhost:8080

show-web:
	/bin/bash $(ROOT_PATHNAME)/open-html \
		http://localhost:8080 \
		nginx-web-content

show-logs:
	make -f makefile-nginx \
		logs \
		DOCKER_VOLUME_NGINX_LOGS=nginx-logs

test: create-web-server \
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
