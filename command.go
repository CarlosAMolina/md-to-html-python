package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func run(command string) []byte {
	fmt.Println(command)
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		panic(err)
	}
	if len(string(out)) != 0 {
		fmt.Println(strings.TrimSuffix(string(out), "\n"))
	}
	return out
}

func runsOk(command string) bool {
	fmt.Println(command)
	err := exec.Command("bash", "-c", command).Run()
	return err == nil
}
