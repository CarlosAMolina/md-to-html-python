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
		run("git clone --depth=1 --branch=main https://github.com/CarlosAMolina/" + repo + " " + repoPath)
	}
}
