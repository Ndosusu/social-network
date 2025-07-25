package handlers_post

import (
	"math"
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func GroupFeedHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)

	groupID, groupIDOk := data["session_uuid"].(int)
	lastID, lastIDOk := data["last_id"].(int)
	limit, _ := data["limit"].(int)
	// Default to max int64 if LastID is not provided or invalid
	if !lastIDOk || lastID <= 0 {
		lastID = math.MaxInt64
	}
	if !groupIDOk || groupID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing group ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.GetGroupFeed(map[string]any{
		"group_id": data["GroupID"],
		"last_id":  lastID,
		"limit":    limit,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve group feed", nil)
		return
	}
	db.CloseConn()

	utils.JSONResponse(w, http.StatusOK, "Group feed retrieved successfully", result)
}
