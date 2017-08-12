package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	bin, _ := exec.LookPath("ls") // find abs path of ls
	syscall.Exec(
		bin,
		[]string{"ls", "-l", "-h"}, // note the first argument is program name since exec will not know what to exec
		os.Environ())               // new env == current env
}
