package handlers

import "net/http"

func NotificationHandler(w http.ResponseWriter, r *http.Request) {

	setCORSHeaders(w, r)

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		// Logic to retrieve notifications
		w.Write([]byte("Retrieve notifications"))
		return
	}

	if r.Method == http.MethodPost {
		// Logic to create a new notification
		w.Write([]byte("Create notification"))
		return
	}
}
