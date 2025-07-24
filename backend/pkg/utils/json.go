package utils

import (
	"encoding/json"
	"net/http"
)

// Helper functions to reduce code duplication

// JSONResponse sends a JSON response with the specified status code, message, and data.
func JSONResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]any{
		"message": message,
		"data":    data,
	}
	_ = json.NewEncoder(w).Encode(response)
}

// JSONDecode decodes JSON data from the request body into a map.
func JSONDecode(w http.ResponseWriter, r *http.Request) map[string]any {
	var data map[string]any
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		JSONResponse(w, http.StatusBadRequest, "Invalid JSON data", nil)
		return nil
	}
	return data
}
