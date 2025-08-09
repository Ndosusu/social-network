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
	userID := int(data["user_id"].(float64))
	lastID := int(data["last_id"].(float64))
	limit := int(data["limit"].(float64))

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	result, err := udb.GetFollowFrom(map[string]any{
		"client_id": clientID,
		"user_id":   userID,
		"last_id":   lastID,
		"limit":     limit,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve follow from", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Follow from retrieved successfully", result)
}
