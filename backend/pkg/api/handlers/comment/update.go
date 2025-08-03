package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	comment "social-network/pkg/db/models/comment"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	// Parse multipart form data to handle file uploads
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form data", nil)
		return
	}

	updateData := map[string]any{
		"comment_id": r.FormValue("comment_id"),
		"message":    r.FormValue("message"),
	}

	// Process image upload
	imageName, err := utils.ImageProcess(r, "image")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if imageName != "" {
		updateData["image"] = imageName
	}

	// Logic to get Author ID
	sessionUUID := r.FormValue("session_uuid")
	if sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing required field: session_uuid", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	result, err := udb.GetSessionByUuid(map[string]any{"session_uuid": sessionUUID})
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid author UUID or user not found", nil)
		return
	}
	updateData["author_id"] = result.Result.(models.Session).User.Id

	cdb := comment.New(&db)
	result, err = cdb.UpdateComment(updateData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to update comment", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Comment updated successfully", result)
}
