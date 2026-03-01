package main

import (
	"os"
	"path/filepath"
)

func prepareMdContentToConvert() error {
	err := os.RemoveAll(mdPath)
	if err != nil {
		return err
	}
	err = os.Mkdir(mdPath, 0775)
	if err != nil {
		return err
	}
	// cmoli.es
	cmoliPath := filepath.Join(getPathSoftware(), "cmoli.es", "src")
	run("cp -r " + cmoliPath + "/* " + mdPath)
	// check-iframe
	checkIframePath := filepath.Join(getPathSoftware(), "checkIframe", "docs")
	checkIframePathNew := filepath.Join(mdPath, "projects", "check-iframe")
	err = os.MkdirAll(checkIframePathNew, 0775)
	run("cp -r " + checkIframePath + "/* " + checkIframePathNew)
	// wiki
	wikiPath := filepath.Join(getPathSoftware(), "wiki", "src")
	wikiPathNew := filepath.Join(mdPath, "wiki")
	err = os.Mkdir(wikiPathNew, 0775)
	run("cp -r " + wikiPath + "/* " + wikiPathNew)
	// tools
	toolNames := [3]string{"open-urls", "job-check-lambda-name", "job-modify-issue-name"}
	toolsPathNew := filepath.Join(mdPath, "tools")
	for i := range len(toolNames) {
		toolRepo := toolNames[i]
		toolPath := filepath.Join(getPathSoftware(), toolRepo)
		run("cp -r " + toolPath + " " + toolsPathNew)
		err := os.RemoveAll(filepath.Join(toolsPathNew, toolRepo, ".git"))
		if err != nil {
			return err
		}
	}
	return nil
}
