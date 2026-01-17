package main

import (
	"fmt"
	"os/exec"
)

func startDockerService() {
	fmt.Println("TODO rm: ", isServiceActive())
}

func isServiceActive() bool {
	cmd := exec.Command("systemctl", "--user", "is-active", "--quiet", "docker")
	err := cmd.Run()
	if err == nil {
		return true
	}
	return false
}
