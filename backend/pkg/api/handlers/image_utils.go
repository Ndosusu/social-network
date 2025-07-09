package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// validateImageFile checks if the uploaded file is a valid image
func validateImageFile(file multipart.File, header *multipart.FileHeader) error {
	// Check file size (max 5MB)
	const maxSize = 5 * 1024 * 1024
	if header.Size > maxSize {
		return fmt.Errorf("file size too large, maximum 5MB allowed")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	validExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}

	isValidExt := false
	for _, validExtension := range validExtensions {
		if ext == validExtension {
			isValidExt = true
			break
		}
	}

	if !isValidExt {
		return fmt.Errorf("invalid file type, only jpg, jpeg, png, gif, webp are allowed")
	}

	return nil
}

// saveImageFile saves the uploaded image and returns the file path
func saveImageFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Create uploads directory if it doesn't exist
	uploadDir := "uploads/images"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	// Create the file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filePath, nil
}

// getImageContentType returns the appropriate content type based on file extension
func getImageContentType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}
