package handlers_user

import (
	"net/http"
	"social-network/pkg/db/models"
	rel "social-network/pkg/db/models/relation"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func FollowDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing id", nil)
		return
	}

	if userIDInt == data["client_id"].(int) {
		utils.JSONResponse(w, http.StatusBadRequest, "You cannot follow yourself", nil)
		return
	}

	rdb := rel.New(&db)
	result, err = rdb.DeleteFollowRel(map[string]any{
		"user_from": data["client_id"],
		"user_to":   userIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create follow relationship", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Follow relationship created successfully", result)
}
