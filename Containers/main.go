// Ref: https://www.youtube.com/watch?v=Utf-A4rODH8
// Building a container from scratch in Go - Liz Rice

// You need to run this go program under sudo
// This allows it to make the SysProcAttr calls

// syscall.CLONE_NEWUTS protects the hostname of the callee host (hostname calls)
// syscall.CLONE_NEWPID shows only PIDs in running shell



package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Oops??")
	}
}

func run() {
	// Fork/exec
	cmd := exec.Command("/proc/self/exe",append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:],os.Getpid())
	cmd := exec.Command(os.Args[2],os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("/home/dc/test/images"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc","proc","proc",0,""))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}