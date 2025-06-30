package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run path/to/main.go <script_name> <options>")
		fmt.Println("Available script: migrate, pair")
	}

	_, filename, _, _ := runtime.Caller(0)

	scriptName := os.Args[1]
	scriptArgs := os.Args[2:]
	baseDir, _ := strings.CutSuffix(filename, "cmd/cli/main.go")
	scriptPath := baseDir + "script/" + scriptName + ".sh"
	fmt.Println("Script path:", scriptPath)

	cmd := exec.Command("/bin/bash", append([]string{scriptPath}, scriptArgs...)...)
	fmt.Println("Executing command:", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(output))
}
