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
	clientID := int(data["client_id"].(float64))

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()
	udb := user.New(&db)

	userID, userIDOk := data["user_id"].(float64)
	var userIDInt int
	if userIDOk {
		userIDInt = int(userID)
	} else {
		userIDInt = clientID
	}

	result, err := udb.SelectUserById(map[string]any{
		"user_id":   userIDInt,
		"client_id": clientID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve user profile", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "User profile retrieved successfully", result)
}
