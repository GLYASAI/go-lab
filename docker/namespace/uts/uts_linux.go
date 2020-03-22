// +build linux

package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Specify the initial command in the new process that is forked.
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS, // TODO: what is uintptr?
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// Look at the relationship between processes in the system
	// pstree -pl
	// Output the current PID
	// echo $$
	// Verify that the parent and child processes are not in the same UTS Namespace
	// readlink /proc/1590/ns/uts
	// uts:[4026531838]
	// readlink /proc/1595/ns/uts
	// uts:[4026532138]
}
