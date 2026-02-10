package main

import (
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
