package handlers_post

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/post"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	updateData := map[string]any{
		"post_id":      r.FormValue("post_id"),
		"message":      r.FormValue("message"),
		"privacy_mode": r.FormValue("privacy_mode"),
		"group_id":     r.FormValue("group_id"),
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

	pdb := post.New(&db)
	result, err = pdb.UpdatePost(updateData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to update post", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Post updated successfully", result)
}
