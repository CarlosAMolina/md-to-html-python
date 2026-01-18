package main

import (
	"fmt"
	"os/exec"
	"time"
)

func startDockerService() {
	var hasBeenActivated = false
	for {
		if isServiceActive() {
			break
		}
		if !hasBeenActivated {
			run("systemctl --user start docker")
			hasBeenActivated = true
			continue
		}
		sleep()
	}
	for {
		// Although the docker daemon is active, on some systems it takes some time before docker runs.
		if runsOk("docker ps") {
			break
		}
		sleep()
	}
}

func isServiceActive() bool {
	return runsOk("systemctl --user is-active --quiet docker")
}

func runsOk(command string) bool {
	fmt.Println(command)
	err := exec.Command("bash", "-c", command).Run()
	if err == nil {
		return true
	}
	return false
}

func sleep() {
	time.Sleep(5 * time.Second)
}
