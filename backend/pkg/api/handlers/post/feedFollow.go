package handlers_post

import (
	"fmt"
	"math"
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func FollowFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	lastID, lastIDOk := data["last_id"].(int)
	limit, _ := data["limit"].(int)
	// Default to max int64 if LastID is not provided or invalid
	if !lastIDOk || lastID <= 0 {
		lastID = math.MaxInt64
	}
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.GetFollowFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastID,
		"limit":        limit,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve follow feed", nil)
		fmt.Println(err)
		return
	}
	db.CloseConn()

	utils.JSONResponse(w, http.StatusOK, "Follow feed retrieved successfully", result)
}
