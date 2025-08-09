package handlers_auth

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)
	sessionUUID := data["session_uuid"].(string)

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	sdb := user.New(&db)
	result, err := sdb.CloseSession(map[string]any{
		"session_uuid": sessionUUID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to close session", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Logout successful", result)
}
