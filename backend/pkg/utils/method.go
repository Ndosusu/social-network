package utils

import "net/http"

func ValidateMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		JSONResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return false
	}
	return true
}
