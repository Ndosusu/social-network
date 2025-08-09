package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func FollowFromHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)

	clientID := int(data["client_id"].(float64))
	userID, userIDOk := data["user_id"].(float64)
	var userIDInt int
	if userIDOk {
		userIDInt = int(userID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing user ID", nil)
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

	udb := user.New(&db)
	result, err := udb.GetFollowFrom(map[string]any{
		"client_id": clientID,
		"user_id":   userIDInt,
		"last_id":   lastIDInt,
		"limit":     limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve follow from", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Follow from retrieved successfully", result)
}
