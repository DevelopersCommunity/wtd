package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := &exec.Cmd{
		Path: "cmd",
		Args: []string{"/c", "start", "wt", "-d", "."},
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
