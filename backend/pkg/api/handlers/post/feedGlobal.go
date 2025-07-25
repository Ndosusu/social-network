package handlers_post

import (
	"fmt"
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func GlobalFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)
	fmt.Println(data)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
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
	result, err := pdb.GetGlobalFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastIDInt,
		"limit":        limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve global feed", nil)
		return
	}
	db.CloseConn()

	utils.JSONResponse(w, http.StatusOK, "Global feed retrieved successfully", result)
}
