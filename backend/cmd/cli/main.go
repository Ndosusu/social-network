package main

import (
	"fmt"
	"log"
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

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error getting source file path")
	}

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
