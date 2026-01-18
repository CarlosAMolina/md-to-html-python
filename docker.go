package main

import (
	"strings"
	"time"
)

func stopContainer(container string) {
	var hasBeenStopped = false
	for {
		if !isContainerRunning(container) {
			break
		}
		if !hasBeenStopped {
			run("docker stop " + container)
			hasBeenStopped = true
		}
		sleep()
	}
}

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

func isContainerRunning(container string) bool {
	out := run("docker ps --format '{{.Names}}'")
	return strings.Contains(string(out), container)
}
