package handlers_post

// ...
/* import (
	"encoding/json"
	"net/http"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	postID, err := parsePostID(r)
	if err != nil || postID == 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing post ID")
		return
	}

	var updateData map[string]any
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}

	// Check if post exists
	if existingPost := dbInstance.SelectPostById(map[string]any{"id": postID}); existingPost.Result == nil {
		utils.JSONResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	updateData["id"] = postID
	result := dbInstance.UpdatePost(updateData)
	if result.Result == 0 {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to update post", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Post updated successfully", result.Result)
}
*/
