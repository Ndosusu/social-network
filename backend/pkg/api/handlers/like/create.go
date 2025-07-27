package handlers_like

import (
	"net/http"
	"social-network/pkg/db/models"
	like "social-network/pkg/db/models/like"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}
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

	udb := user.New(&db)
	sessionResult, err := udb.GetSessionByUuid(map[string]any{
		"session_uuid": sessionUUID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, "Invalid session", nil)
		return
	}
	userID := sessionResult.Result.(models.Session).User.Id

	ldb := like.New(&db)
	result, err := ldb.InsertLike(map[string]any{
		"post_id":    postIDInt,
		"comment_id": commentIDInt,
		"user_id":    userID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to insert like into database", nil)
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Like added successfully", result)

}
