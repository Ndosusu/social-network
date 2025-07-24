package handlers_post

import (
	"io"
	"net/http"
	"os"
	"social-network/pkg/utils"
	"strings"
)

// Helper functions for post handlers

func validateRequiredFields(data map[string]any, fields []string) string {
	for _, field := range fields {
		if data[field] == nil {
			return "Missing required field: " + field
		}
	}
	return ""
}

// ServeImageHandler serves uploaded images
func ServeImageHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodGet) {
		return
	}

	// Get the image path from URL parameter
	imagePath := r.URL.Query().Get("path")
	if imagePath == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing image path parameter", nil)
		return
	}

	// Security check - ensure path is within uploads directory
	if !strings.HasPrefix(imagePath, "uploads/images/") {
		utils.JSONResponse(w, http.StatusForbidden, "Invalid image path", nil)
		return
	}

	// Check if file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		utils.JSONResponse(w, http.StatusNotFound, "Image not found", nil)
		return
	}

	// Open the file
	file, err := os.Open(imagePath)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to open image file", nil)
		return
	}
	defer file.Close()

	// Set content type based on file extension using utility function
	w.Header().Set("Content-Type", utils.GetImageContentType(imagePath))

	// Copy file content to response
	if _, err := io.Copy(w, file); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to serve image", nil)
		return
	}
}
