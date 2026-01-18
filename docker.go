package main

import (
	"strings"
	"time"
)

func removeVolume(volume string) {
	var hasBeenRemoved = false
	for {
		if !existsVolume(volume) {
			break
		}
		if !hasBeenRemoved {
			run("docker volume rm " + volume)
			hasBeenRemoved = true
		}
		sleep(1)
	}
}

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
		sleep(1)
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
		sleep(5)
	}
	for {
		// Although the docker daemon is active, on some systems it takes some time before docker runs.
		if runsOk("docker ps") {
			break
		}
		sleep(5)
	}
}

func isServiceActive() bool {
	return runsOk("systemctl --user is-active --quiet docker")
}

func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func isContainerRunning(container string) bool {
	out := run("docker ps --format '{{.Names}}'")
	return strings.Contains(string(out), container)
}

func existsVolume(volume string) bool {
	out := run("docker volume ls -q")
	return strings.Contains(string(out), volume)
}
