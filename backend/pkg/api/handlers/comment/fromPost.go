package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	comment "social-network/pkg/db/models/comment"
	"social-network/pkg/utils"
)

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodGet) {
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
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing post ID", nil)
		return
	}
	lastID, lastIDOk := data["last_id"].(float64)
	var lastIDInt int
	if lastIDOk {
		lastIDInt = int(lastID)
	} else {
		// Default to max int if last_id is not provided or invalid
		lastIDInt = utils.DEFAULT_ID
	}
	limit, limitOk := data["limit"].(float64)
	var limitInt int
	if limitOk {
		limitInt = int(limit)
	} else {
		// Default to 10 if limit is not provided or invalid
		limitInt = utils.DEFAULT_LIMIT
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	cdb := comment.New(&db)
	result, err := cdb.SelectCommentsByPostId(map[string]any{
		"session_uuid": sessionUUID,
		"post_id":      postIDInt,
		"last_id":      lastIDInt,
		"limit":        limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve comments", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Comments retrieved successfully", result)

}
