// +build linux
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

// the root path of the hierarchy that the memory subsystem is mounted.
const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		// container progress
		fmt.Printf("current pid %d\n", syscall.Getpid())
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	// Get the pid of the fork process in the external namespace
	fmt.Printf("pid in the external namespace: %d\n", cmd.Process.Pid)
	// Create a cgroup on the Hierarchy with the memory subsystem mounted by default
	os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)
	// Add container processes to this cgroup
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"),
		[]byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	// Limiting the use of cgroup processes
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit",
		"memory.limit_in_bytes"), []byte("100m"), 0644)
	cmd.Process.Wait()
}
