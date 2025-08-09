package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
	"strconv"
)

func UpdateDataHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	if r.FormValue("client_id") != r.FormValue("user_id") {
		utils.JSONResponse(w, http.StatusForbidden, "You can only update your own data", nil)
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
	result, err := udb.UpdateData(
		map[string]any{
			"user_id":    userID,
			"first_name": r.FormValue("firstName"),
			"last_name":  r.FormValue("lastName"),
			"mail":       r.FormValue("mail"),
			"nickname":   r.FormValue("nickname"),
			"about":      r.FormValue("about"),
		})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to update user data", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "User data updated successfully", result)
}
