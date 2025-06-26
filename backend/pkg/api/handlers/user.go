package handlers

import "net/http"

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	setCORSHeaders(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func FollowUserHandler(w http.ResponseWriter, r *http.Request) {

	setCORSHeaders(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
