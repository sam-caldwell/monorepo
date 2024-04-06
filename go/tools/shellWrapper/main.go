package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	shellCommand = "/bin/bash"
)

func main() {
	var commandPending bool
	// Start a /bin/bash process
	bashCmd := exec.Command(shellCommand)

	// Create pipes for stdin and stdout
	stdin, err := bashCmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		return
	}
	stdout, err := bashCmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return
	}

	// Start the bash process
	err = bashCmd.Start()
	if err != nil {
		fmt.Println("Error starting bash command:", err)
		return
	}

	// Stream the stdout of the bash process back to the console
	go func() {
		io.Copy(os.Stdout, stdout)
		if commandPending {
			fmt.Println("")
		}
		commandPending = false
	}()

	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Infinite loop to continuously accept input from stdin
	for {
		fmt.Printf("shell pid: %v ", bashCmd.Process.Pid)

		// Read the line of text from stdin
		scanner.Scan()
		text := scanner.Text()

		// Check if the input is empty
		if strings.TrimSpace(text) == "" {
			continue
		}
		commandPending = true
		// Stream the text to the stdin of the bash process
		_, err := io.WriteString(stdin, text+"\n")
		if err != nil {
			fmt.Println("Error writing to stdin:", err)
			return
		}
	}
}
