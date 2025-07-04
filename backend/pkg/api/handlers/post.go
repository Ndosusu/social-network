// Package handlers provides HTTP handlers for the social network API
package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/pkg/db/models"
	"strconv"
)

// Helper functions for post handlers
func parsePostID(r *http.Request) (int, error) {
	postIDStr := r.URL.Query().Get("id")
	if postIDStr == "" {
		return 0, nil
	}
	return strconv.Atoi(postIDStr)
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

	var postData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

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

	// Set defaults - group_id will be NULL if not provided

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
	result := dbInstance.SelectPostById(map[string]any{"id": postID})
	post, ok := result.Result.(models.Post)
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
			result = dbInstance.SelectPostsByUserId(queryParams)
		} else {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid user ID format")
			return
		}
	case r.URL.Query().Get("group_id") != "":
		if groupID, err := strconv.Atoi(r.URL.Query().Get("group_id")); err == nil {
			queryParams["group_id"] = groupID
			result = dbInstance.SelectPostsByGroupId(queryParams)
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
		result = dbInstance.SelectAllPosts(queryParams)
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
