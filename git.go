package main

import (
	"path/filepath"
)

func pullGitRepos() {
	pullGitRepo("cmoli.es")
	pullGitRepo("cmoli.es-deploy")
	pullGitRepo("checkIframe")
	pullGitRepo("wiki")
	pullGitTools()
}

func pullGitRepo(repo string) {
	repoPath := filepath.Join(getPathSoftware(), repo)
	if exists(repoPath) {
		run("cd " + repoPath + " && git pull origin $(git branch --show-current)")
	} else {
		run("git clone git@github.com:CarlosAMolina/" + repo + " " + repoPath)
	}
}
