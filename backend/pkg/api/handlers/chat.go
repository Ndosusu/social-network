package handlers

import "net/http"

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w, r)

}
