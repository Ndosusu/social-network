package handlers

import (
    "encoding/json"
    "net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(map[string]string{"message": "Get users"})
    case "POST":
        json.NewEncoder(w).Encode(map[string]string{"message": "Create user"})
    }
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(map[string]string{"message": "Get posts"})
    case "POST":
        json.NewEncoder(w).Encode(map[string]string{"message": "Create post"})
    }
}