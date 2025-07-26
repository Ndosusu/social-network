package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodGet) {
		return
	}

	data := utils.JSONDecode(w, r)

	postID, postIDOk := data["id"].(int)
	lastID, lastIDOk := data["last_id"].(int)
	if !postIDOk || postID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing post ID", nil)
		return
	}
	if !lastIDOk || lastID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing comment ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.SelectCommentsByPostId(map[string]any{
		"post_id": postID,
		"last_id": lastID,
	})
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve comments", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Comments retrieved successfully", result)

}
