package handlers_post

import (
	"fmt"
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/post"
	"social-network/pkg/utils"
)

func FollowFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
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
	result, err := pdb.GetFollowFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastIDInt,
		"limit":        limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve follow feed", nil)
		fmt.Println(err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Follow feed retrieved successfully", result)
}
