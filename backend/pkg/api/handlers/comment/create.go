package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	comment "social-network/pkg/db/models/comment"
	"social-network/pkg/utils"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	// Extract form data
	comData := map[string]any{
		"post_id":   r.FormValue("post_id"),
		"message":   r.FormValue("message"),
		"author_id": r.FormValue("client_id"),
	}

	// Process image upload
	imageName, err := utils.ImageProcess(r, "image")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if imageName != "" {
		comData["image"] = imageName
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	cdb := comment.New(&db)
	result, err := cdb.InsertComment(comData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create comment", nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "Comment created successfully", result)
}
