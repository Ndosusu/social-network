package handlers

import "net/http"

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w, r)
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {

		//...

		w.Write([]byte("Retrieve comments"))
		return
	}

	if r.Method == http.MethodPost {

		//...
	}
}
