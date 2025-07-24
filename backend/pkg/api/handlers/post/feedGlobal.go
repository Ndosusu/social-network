package handlers_post

import (
	"math"
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func GlobalFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["SessionUUID"].(string)
	lastID, lastIDOk := data["LastID"].(int)
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
	result, err := pdb.GetGlobalFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve global feed", nil)
		return
	}
	db.CloseConn()

	utils.JSONResponse(w, http.StatusOK, "Global feed retrieved successfully", result)
}
