package handlers_post

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func FollowFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["SessionUUID"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}
	lastID, lastIDOk := data["LastID"].(float64)
	var lastIDInt int
	if lastIDOk {
		lastIDInt = int(lastID)
	} else {
		// Default to max int64 if LastID is not provided or invalid
		lastIDInt = utils.DEFAULT_ID
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.GetFollowFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve global feed", nil)
		return
	}
	db.CloseConn()

	utils.JSONResponse(w, http.StatusOK, "Global feed retrieved successfully", result)
}
