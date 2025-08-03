package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"social-network/config"
	"strings"
	"time"
)

func ImageProcess(r *http.Request, fieldName string) (string, error) {
	file, header, err := r.FormFile(fieldName)
	if err == http.ErrMissingFile {
		return "", nil // No image, no errors
	}
	if err != nil {
		return "", fmt.Errorf("error processing image file")
	}
	defer file.Close()

	if err := ValidateImageFile(file, header); err != nil {
		return "", err
	}

	return SaveImageFile(file, header)
}

// validateImageFile checks if the uploaded file is a valid image
func ValidateImageFile(file multipart.File, header *multipart.FileHeader) error {
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

// SaveImageFile saves the uploaded image and returns the file name
func SaveImageFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Create uploads directory if it doesn't exist
	uploadDir := config.PicPath
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, fileName)

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

	return fileName, nil
}

// GetImageContentType returns the appropriate content type based on file extension
func GetImageContentType(filePath string) string {
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
