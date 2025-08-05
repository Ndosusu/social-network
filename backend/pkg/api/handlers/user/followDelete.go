package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	rel "social-network/pkg/db/models/relation"
	"social-network/pkg/utils"
)

func FollowDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)
	clientID := int(data["client_id"].(float64))

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	userID, userIDOk := data["user_id"].(float64)
	var userIDInt int
	if userIDOk {
		userIDInt = int(userID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing id", nil)
		return
	}

	if userIDInt == data["client_id"].(int) {
		utils.JSONResponse(w, http.StatusBadRequest, "You cannot follow yourself", nil)
		return
	}

	rdb := rel.New(&db)
	result, err := rdb.DeleteFollowRel(map[string]any{
		"user_from": clientID,
		"user_to":   userIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create follow relationship", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Follow relationship created successfully", result)
}
