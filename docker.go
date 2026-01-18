package main

import (
	"fmt"
	"os/exec"
	"time"
)

func startDockerService() {
	if !isServiceActive() {
		run("systemctl --user start docker")
	}
	for {
		if isServiceActive() {
			break
		}
		time.Sleep(5 * time.Second)
	}
	for {
		// Although the docker daemon is active, on some systems it takes some time before docker runs.
		if runsOk("docker ps") {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func isServiceActive() bool {
	return runsOk("systemctl --user is-active --quiet docker")
}

func runsOk(command string) bool {
	err := runGetError(command)
	if err == nil {
		return true
	}
	return false
}

func runGetError(command string) error {
	fmt.Println(command)
	err := exec.Command("bash", "-c", command).Run()
	return err
}
