package handlers_post

import (
	"fmt"
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
	fmt.Println(data)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	lastID, lastIDOk := data["last_id"].(int)
	limit, limitOk := data["limit"].(int)
	fmt.Println(limit)
	// Default to max int64 if LastID is not provided or invalid
	if !lastIDOk || lastID <= 0 {
		lastID = math.MaxInt64
	}
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}
	if !limitOk || limit <= 0 {
		limit = 20
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.GetGlobalFeed(map[string]any{
		"session_uuid": sessionUUID,
		"last_id":      lastID,
		"limit":        limit,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve global feed", nil)
		return
	}
	db.CloseConn()

	fmt.Println(result)

	utils.JSONResponse(w, http.StatusOK, "Global feed retrieved successfully", result)
}
