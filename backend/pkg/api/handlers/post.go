package handlers

import "net/http"

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	setCORSHeaders(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	setCORSHeaders(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
