package main

// TODO the `md-to-html` folder must only convert content, the responsability for copying web content must be outside this folder.
func copyContentToVolumeNginx() {
	volumePath := getVolumePath("nginx-web-content")
	run("cp -r ../src/* $volume_web_content_pathname " + volumePath)
}
