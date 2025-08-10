package handlers_like

import (
	"net/http"
	"social-network/pkg/db/models"
	like "social-network/pkg/db/models/like"
	"social-network/pkg/utils"
)

func CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)

	clientID := int(data["client_id"].(float64))

	postID, postIDOk := data["post_id"].(float64)
	var postIDInt int
	if postIDOk {
		postIDInt = int(postID)
	}
	commentID, commentIDOk := data["comment_id"].(float64)
	var commentIDInt int
	if commentIDOk {
		commentIDInt = int(commentID)
	}

	// Validate that at only one of post_id or comment_id is provided
	if !postIDOk && !commentIDOk {
		utils.JSONResponse(w, http.StatusBadRequest, "Must specify either post_id or comment_id", nil)
		return
	}
	if postIDOk && commentIDOk {
		utils.JSONResponse(w, http.StatusBadRequest, "Cannot like both post and comment simultaneously", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	ldb := like.New(&db)
	result, err := ldb.InsertLike(map[string]any{
		"post_id":    postIDInt,
		"comment_id": commentIDInt,
		"user_id":    clientID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to insert like into database", nil)
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Like added successfully", result)

}
