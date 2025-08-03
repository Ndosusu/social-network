package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	comment "social-network/pkg/db/models/comment"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	// Parse multipart form data to handle file uploads
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form data", nil)
		return
	}

	// Extract form data
	comData := map[string]any{
		"post_id": r.FormValue("post_id"),
		"message": r.FormValue("message"),
	}

	// Logic to get Author ID
	sessionUUID := r.FormValue("session_uuid")
	if sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing required field: author_uuid", nil)
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
	comData["author_id"] = result.Result.(models.Session).User.Id

	// Process image upload
	imageName, err := utils.ImageProcess(r, "image")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if imageName != "" {
		comData["image"] = imageName
	}

	cdb := comment.New(&db)
	result, err = cdb.InsertComment(comData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create comment", nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "Comment created successfully", result)
}
