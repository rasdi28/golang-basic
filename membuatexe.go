package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("cmd")

	// Connect input/output streams
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Wait for the command to complete
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command execution error:", err)
		return
	}

	fmt.Println("Command completed successfully.")
}
