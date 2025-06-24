package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"social-network/config"
	"social-network/pkg/api"
	"social-network/pkg/db/sqlite"
	"strings"
	"time"
)

// This is the main entry point for the server application.
func main() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir, _ := strings.CutSuffix(filename, "cmd/server/main.go")

	err := config.LoadEnvFile(baseDir + "config/config.env") // Load environment variables from local.env
	if err != nil {
		log.Fatalf("Loading error on file .env: %v", err)
	}

	cmd := exec.Command("/bin/bash", baseDir+"script/data.sh") // Execute the data script
	fmt.Println("Executing command:", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(output))

	err = sqlite.Connect(config.DBPath)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	router := api.InitRouter() // Initialize the router from the api package

	server := &http.Server{
		Addr:              config.ServerPort,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Println("http://localhost:8080")
	if err := server.ListenAndServe(); err != nil { // open server
		log.Fatal(err)
	}
}
