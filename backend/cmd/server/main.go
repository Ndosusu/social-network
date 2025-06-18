package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "social-network-backend/internal/handlers"
)

func main() {
    r := mux.NewRouter()

    // Add root route
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Social Network API",
            "status": "running",
            "endpoints": "/api/users, /api/posts",
        })
    })

    // Set up routes
    r.HandleFunc("/api/users", handlers.HandleUser).Methods("GET", "POST")
    r.HandleFunc("/api/posts", handlers.HandlePost).Methods("GET", "POST")

    // Start the server
    log.Println("Starting server on http://localhost:3281")
    if err := http.ListenAndServe(":3281", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}