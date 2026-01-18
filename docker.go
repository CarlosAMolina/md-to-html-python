package main

import (
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

func sleep() {
	time.Sleep(5 * time.Second)
}
