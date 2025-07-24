package handlers_auth

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	loginData := utils.JSONDecode(w, r)

	mail, mailOk := loginData["Mail"].(string)
	password, passwordOk := loginData["Password"].(string)

	if !mailOk || !passwordOk || mail == "" || password == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing email or password", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	udb := user.New(&db)

	result, err := udb.Authenticate(map[string]any{
		"mail":     mail,
		"password": password,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Database error: "+err.Error(), nil)
		return
	}
	db.CloseConn()

	if result.Result == nil {
		utils.JSONResponse(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	session, ok := result.Result.(models.Session)
	if !ok {
		utils.JSONResponse(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Login successful", map[string]any{
		"session_uuid": session.Uuid,
	})
}
