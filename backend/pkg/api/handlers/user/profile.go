package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()
	udb := user.New(&db)

	result, err := udb.GetSessionByUuid(map[string]any{"session_uuid": sessionUUID})
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid session UUID or user not found", nil)
		return
	}
	data["client_id"] = result.Result.(models.Session).User.Id

	userID, userIDOk := data["user_id"].(float64)
	var userIDInt int
	if userIDOk {
		userIDInt = int(userID)
	} else {
		userIDInt = data["client_id"].(int)
	}

	result, err = udb.SelectUserById(map[string]any{
		"user_id":   userIDInt,
		"client_id": data["client_id"],
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve user profile", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "User profile retrieved successfully", result)
}
