package handlers_post

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/post"
	"social-network/pkg/utils"
)

func GlobalFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	sessionUUID := data["session_uuid"].(string)

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

	pdb := post.New(&db)
	result, err := pdb.GetGlobalFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastIDInt,
		"limit":        limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve global feed", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Global feed retrieved successfully", result)
}
