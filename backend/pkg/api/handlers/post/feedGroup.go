package handlers_post

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func GroupFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}
	groupID, groupIDOk := data["group_id"].(float64)
	var groupIDInt int
	if groupIDOk {
		groupIDInt = int(groupID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing group ID", nil)
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
	pdb := post.New(&db)
	result, err := pdb.GetGroupFeed(map[string]any{
		"group_id": groupIDInt,
		"last_id":  lastIDInt,
		"limit":    limitInt,
	})
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve group feed", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Group feed retrieved successfully", result)
}
