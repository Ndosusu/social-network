package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
	"strconv"
)

func UpdateAvatarHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	if r.FormValue("client_id") != r.FormValue("user_id") {
		utils.JSONResponse(w, http.StatusForbidden, "You can only update your own avatar", nil)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil || userID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	// Process avatar image
	imageName, err := utils.ImageProcess(r, "avatar")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if imageName == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "No image provided", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	result, err := udb.UpdateAvatar(
		map[string]any{
			"user_id": userID,
			"avatar":  imageName,
		})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create user account", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Avatar updated successfully", result)

}

// 3 form : avatar/infos/password
