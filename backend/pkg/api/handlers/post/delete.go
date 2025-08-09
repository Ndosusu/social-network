package handlers_post

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/post"
	"social-network/pkg/utils"
)

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodDelete) {
		return
	}

	data := utils.JSONDecode(w, r)

	sessionUUID := data["session_uuid"].(string)

	postID, postIDOk := data["post_id"].(float64)
	var postIDInt int
	if postIDOk {
		postIDInt = int(postID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing post ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	pdb := post.New(&db)
	result, err := pdb.IsUserPostAuthor(map[string]any{
		"post_id":      postIDInt,
		"session_uuid": sessionUUID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, "Post not found or invalid session", nil)
		return
	}

	canDelete := result.Result.(bool)
	if !canDelete {
		utils.JSONResponse(w, http.StatusForbidden, "You are not the author of this post", nil)
		return
	}

	result, err = pdb.DeletePost(map[string]any{
		"post_id": postIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to delet post", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Post deleted successfully", result)
}
