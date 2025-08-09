package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
	"strconv"
)

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	if r.FormValue("client_id") != r.FormValue("user_id") {
		utils.JSONResponse(w, http.StatusForbidden, "You can only update your own password", nil)
		return
	}

	if !utils.IsValidPassword(r.FormValue("password")) {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid password format", nil)
		return
	}
	if r.FormValue("password") != r.FormValue("rpassword") {
		utils.JSONResponse(w, http.StatusBadRequest, "Passwords do not match", nil)
		return
	}
	if r.FormValue("password") == r.FormValue("opassword") {
		utils.JSONResponse(w, http.StatusBadRequest, "New password cannot be the same as the old password", nil)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil || userID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	result, err := udb.UpdatePassword(
		map[string]any{
			"user_id":      userID,
			"old_password": r.FormValue("opassword"),
			"new_password": r.FormValue("password"),
		})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create user account", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Password updated successfully", result)

}
