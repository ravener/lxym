package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Pager(output []byte) bool {
	// Try to find less.
	path, err := exec.LookPath("less")

	// If it wasn't found return and allow the program to continue to write in stdout.
	if err != nil {
		return false
	}

	// create the command. -r flag must be passed to less to allow interpreting raw characters.
	cmd := exec.Command(path, "-r")

	cmd.Stdin = bytes.NewReader(output)
	// XXX: Does less even run on Windows? if not it's safe to just use os.Stdout instead of the colorable wrapper.
	cmd.Stdout = stdout

	err = cmd.Run()

	if err != nil {
		// If something went wrong, print the error and continue showing the output to stdout.
		fmt.Printf("Failed to start 'less': %s\n", err)
		return false
	}

	// Everything went good, the user has quit less at this point.
	return true
}
