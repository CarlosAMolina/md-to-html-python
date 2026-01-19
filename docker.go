package main

import (
	"strings"
	"time"
)

func createVolume(volume string) {
	run("docker volume create " + volume)
}

func pullDocker(image string) {
	run("docker pull " + image)
}

func removeVolume(volume string) {
	var hasBeenRemoved = false
	for existsVolume(volume) {
		if !hasBeenRemoved {
			run("docker volume rm " + volume)
			hasBeenRemoved = true
		}
		sleep(1)
	}
}

// To test the function, you can run a container with:
// docker run -it --rm --name nginx-cmoli-container nginx-cmoli
func stopContainer(container string) {
	var hasBeenStopped = false
	for isContainerRunning(container) {
		if !hasBeenStopped {
			run("docker stop " + container)
			hasBeenStopped = true
		}
		sleep(1)
	}
}

func startDockerService() {
	var hasBeenActivated = false
	for !isServiceActive() {
		if !hasBeenActivated {
			run("systemctl --user start docker")
			hasBeenActivated = true
			continue
		}
		sleep(5)
	}
	for !runsOk("docker ps") {
		// Although the docker daemon is active, on some systems it takes some time before docker runs.
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
