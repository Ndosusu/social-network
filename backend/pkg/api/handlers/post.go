package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/pkg/db/models"
	"strconv"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Parse JSON body
	var postData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON format",
		})
		return
	}

	// Validate required fields
	requiredFields := []string{"author_id", "message", "privacy_mode"}
	for _, field := range requiredFields {
		if postData[field] == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Missing required field: " + field,
			})
			return
		}
	}

	// Connect to database
	db, err := getDBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Create DB instance
	dbInstance := &models.DB{Conn: db}

	// Set default values for optional fields
	if postData["image"] == nil {
		postData["image"] = ""
	}
	if postData["group_id"] == nil {
		postData["group_id"] = 0
	}

	// Insert post
	result := dbInstance.InsertPost(postData)
	if result.Result == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to create post",
		})
		return
	}

	// Return success response with created post
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    result.Result,
		"message": "Post created successfully",
	})
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Get post ID from URL query parameters
	postIdStr := r.URL.Query().Get("id")
	if postIdStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Post ID is required",
		})
		return
	}

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid post ID format",
		})
		return
	}

	// Connect to database
	db, err := getDBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Create DB instance
	dbInstance := &models.DB{Conn: db}

	// Get post by ID
	result := dbInstance.SelectPostById(map[string]any{"id": postId})
	post, ok := result.Result.(models.Post)
	if !ok || post.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Post not found",
		})
		return
	}

	// Return post data
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    post,
	})
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Get query parameters
	userIdStr := r.URL.Query().Get("user_id")
	groupIdStr := r.URL.Query().Get("group_id")
	privacyModeStr := r.URL.Query().Get("privacy_mode")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// Parse parameters
	queryParams := make(map[string]any)

	if limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			queryParams["limit"] = float64(limit)
		}
	}

	if offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil {
			queryParams["offset"] = float64(offset)
		}
	}

	// Connect to database
	db, err := getDBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Create DB instance
	dbInstance := &models.DB{Conn: db}

	var result models.Response

	// Determine which query to use based on parameters
	if userIdStr != "" {
		if userId, err := strconv.Atoi(userIdStr); err == nil {
			queryParams["user_id"] = userId
			result = dbInstance.SelectPostsByUserId(queryParams)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid user ID format",
			})
			return
		}
	} else if groupIdStr != "" {
		if groupId, err := strconv.Atoi(groupIdStr); err == nil {
			queryParams["group_id"] = groupId
			result = dbInstance.SelectPostsByGroupId(queryParams)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid group ID format",
			})
			return
		}
	} else if privacyModeStr != "" {
		if privacyMode, err := strconv.Atoi(privacyModeStr); err == nil {
			queryParams["privacy_mode"] = privacyMode
			result = dbInstance.SelectPostsByPrivacyMode(queryParams)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid privacy mode format",
			})
			return
		}
	} else {
		// Get all posts
		result = dbInstance.SelectAllPosts(queryParams)
	}

	// Return posts
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    result.Result,
	})
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Get post ID from URL query parameters
	postIdStr := r.URL.Query().Get("id")
	if postIdStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Post ID is required",
		})
		return
	}

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid post ID format",
		})
		return
	}

	// Connect to database
	db, err := getDBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Create DB instance
	dbInstance := &models.DB{Conn: db}

	// Check if post exists first
	existingPost := dbInstance.SelectPostById(map[string]any{"id": postId})
	if post, ok := existingPost.Result.(models.Post); !ok || post.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Post not found",
		})
		return
	}

	// Delete the post
	result := dbInstance.DeletePost(map[string]any{"id": postId})
	if result.Result == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to delete post",
		})
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Post deleted successfully",
	})
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Get post ID from URL query parameters
	postIdStr := r.URL.Query().Get("id")
	if postIdStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Post ID is required",
		})
		return
	}

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid post ID format",
		})
		return
	}

	// Parse JSON body
	var updateData map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON format",
		})
		return
	}

	// Connect to database
	db, err := getDBConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Create DB instance
	dbInstance := &models.DB{Conn: db}

	// Check if post exists first
	existingPost := dbInstance.SelectPostById(map[string]any{"id": postId})
	if post, ok := existingPost.Result.(models.Post); !ok || post.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Post not found",
		})
		return
	}

	// Add the post ID to the update data
	updateData["id"] = postId

	// Update the post using UpdatePost method (we'll need to add this to the model)
	result := dbInstance.UpdatePost(updateData)
	if result.Result == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to update post",
		})
		return
	}

	// Return updated post
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    result.Result,
		"message": "Post updated successfully",
	})
}
