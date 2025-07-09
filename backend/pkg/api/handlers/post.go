// Package handlers provides HTTP handlers for the social network API
package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"social-network/pkg/db/models"
	"strconv"
	"strings"
)

// Helper functions for post handlers
func parsePostID(r *http.Request) (int, error) {
	postIDStr := r.URL.Query().Get("id")
	if postIDStr == "" {
		return 0, nil
	}
	return strconv.Atoi(postIDStr)
}

func getAuthorIDFromUUID(uuid string) (int, error) {
	db, err := getDBConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}
	result := dbInstance.SelectUserByUuid(map[string]any{"uuid": uuid})

	user, ok := result.Result.(models.User)
	if !ok || user.Id == 0 {
		return 0, errors.New("user not found")
	}

	return user.Id, nil
}

func validateRequiredFields(data map[string]interface{}, fields []string) string {
	for _, field := range fields {
		if data[field] == nil {
			return "Missing required field: " + field
		}
	}
	return ""
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodPost) {
		return
	}

	// Parse multipart form data to handle file uploads
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		writeErrorResponse(w, http.StatusBadRequest, "Failed to parse form data")
		return
	}

	// Extract form data
	postData := make(map[string]interface{})
	authorUuid := r.FormValue("author_uuid")
	if authorUuid == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Missing required field: author_uuid")
		return
	}

	// Get author_id from UUID
	authorID, err := getAuthorIDFromUUID(authorUuid)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid author UUID or user not found")
		return
	}

	postData["author_id"] = authorID
	postData["message"] = r.FormValue("message")
	postData["privacy_mode"] = r.FormValue("privacy_mode")

	// Group ID is optional
	if groupID := r.FormValue("group_id"); groupID != "" {
		postData["group_id"] = groupID
	}

	// Convert string values to appropriate types
	if privacyMode, err := strconv.Atoi(postData["privacy_mode"].(string)); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid privacy_mode format")
		return
	} else {
		postData["privacy_mode"] = privacyMode
	}

	if groupIDStr, ok := postData["group_id"].(string); ok && groupIDStr != "" {
		if groupID, err := strconv.Atoi(groupIDStr); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid group_id format")
			return
		} else {
			postData["group_id"] = groupID
		}
	}

	// Handle image upload if present
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		// Validate image file
		if err := validateImageFile(file, header); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Save the image
		imagePath, err := saveImageFile(file, header)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "Failed to save image: "+err.Error())
			return
		}

		postData["image"] = imagePath
	} else if err != http.ErrMissingFile {
		writeErrorResponse(w, http.StatusBadRequest, "Error processing image file")
		return
	}

	// Validate required fields
	if errMsg := validateRequiredFields(postData, []string{"author_id", "message", "privacy_mode"}); errMsg != "" {
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}
	result := dbInstance.InsertPost(postData)
	if result.Result == 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	writeSuccessResponse(w, http.StatusCreated, "Post created successfully", result.Result)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodGet) {
		return
	}

	postID, err := parsePostID(r)
	if err != nil || postID == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid or missing post ID")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}
	result := dbInstance.SelectPostWithAuthorById(map[string]any{"id": postID})
	post, ok := result.Result.(models.PostWithAuthor)
	if !ok || post.Id == 0 {
		writeErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	writeSuccessResponse(w, http.StatusOK, "", post)
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodGet) {
		return
	}

	// Parse query parameters
	queryParams := make(map[string]any)

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			queryParams["limit"] = float64(limit)
		}
	}

	// Support for cursor-based pagination
	if lastIDStr := r.URL.Query().Get("last_id"); lastIDStr != "" {
		if lastID, err := strconv.Atoi(lastIDStr); err == nil {
			queryParams["last_id"] = lastID
		}
	}

	// Support for timestamp-based pagination
	if beforeStr := r.URL.Query().Get("before"); beforeStr != "" {
		queryParams["before"] = beforeStr
	}

	// Keep offset for backward compatibility, but warn about potential issues
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			queryParams["offset"] = float64(offset)
		}
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}
	var result models.Response

	// Route based on query parameters
	switch {
	case r.URL.Query().Get("user_id") != "":
		if userID, err := strconv.Atoi(r.URL.Query().Get("user_id")); err == nil {
			queryParams["user_id"] = userID
			result = dbInstance.SelectPostsByUserIdWithAuthors(queryParams)
		} else {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID format")
			return
		}
	case r.URL.Query().Get("group_id") != "":
		if groupID, err := strconv.Atoi(r.URL.Query().Get("group_id")); err == nil {
			queryParams["group_id"] = groupID
			result = dbInstance.SelectPostsByGroupIdWithAuthors(queryParams)
		} else {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid group ID format")
			return
		}
	case r.URL.Query().Get("privacy_mode") != "":
		if privacyMode, err := strconv.Atoi(r.URL.Query().Get("privacy_mode")); err == nil {
			queryParams["privacy_mode"] = privacyMode
			result = dbInstance.SelectPostsByPrivacyMode(queryParams)
		} else {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid privacy mode format")
			return
		}
	default:
		result = dbInstance.SelectAllPostsWithAuthors(queryParams)
	}

	writeSuccessResponse(w, http.StatusOK, "", result.Result)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodDelete) {
		return
	}

	postID, err := parsePostID(r)
	if err != nil || postID == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid or missing post ID")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}

	// Check if post exists
	if existingPost := dbInstance.SelectPostById(map[string]any{"id": postID}); existingPost.Result == nil {
		writeErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	result := dbInstance.DeletePost(map[string]any{"id": postID})
	if result.Result == 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to delete post")
		return
	}

	writeSuccessResponse(w, http.StatusOK, "Post deleted successfully", nil)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodPut) {
		return
	}

	postID, err := parsePostID(r)
	if err != nil || postID == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid or missing post ID")
		return
	}

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}

	// Check if post exists
	if existingPost := dbInstance.SelectPostById(map[string]any{"id": postID}); existingPost.Result == nil {
		writeErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	updateData["id"] = postID
	result := dbInstance.UpdatePost(updateData)
	if result.Result == 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to update post")
		return
	}

	writeSuccessResponse(w, http.StatusOK, "Post updated successfully", result.Result)
}

// ServeImageHandler serves uploaded images
func ServeImageHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodGet) {
		return
	}

	// Get the image path from URL parameter
	imagePath := r.URL.Query().Get("path")
	if imagePath == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Missing image path parameter")
		return
	}

	// Security check - ensure path is within uploads directory
	if !strings.HasPrefix(imagePath, "uploads/images/") {
		writeErrorResponse(w, http.StatusForbidden, "Invalid image path")
		return
	}

	// Check if file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		writeErrorResponse(w, http.StatusNotFound, "Image not found")
		return
	}

	// Open the file
	file, err := os.Open(imagePath)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to open image file")
		return
	}
	defer file.Close()

	// Set content type based on file extension using utility function
	w.Header().Set("Content-Type", getImageContentType(imagePath))

	// Copy file content to response
	if _, err := io.Copy(w, file); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to serve image")
		return
	}
}
