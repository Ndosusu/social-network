package handlers_like

import (
	"net/http"
	"social-network/pkg/db/models"
	like "social-network/pkg/db/models/like"
	"social-network/pkg/utils"
)

func DeleteLikeHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodDelete) {
		return
	}

	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}

	likeID, likeIDOk := data["like_id"].(float64)
	var likeIDInt int
	if likeIDOk {
		likeIDInt = int(likeID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing like ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	ldb := like.New(&db)

	canDeleteResult, err := ldb.IsUserLikeAuthor(map[string]any{
		"session_uuid": sessionUUID,
		"like_id":      likeIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Error verifying author", nil)
		return
	}

	if !canDeleteResult.Result.(bool) {
		utils.JSONResponse(w, http.StatusForbidden, "You are not the author of this like", nil)
		return
	}

	result, err := ldb.DeleteLike(map[string]any{"like_id": likeIDInt})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Error deleting like", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Like deleted successfully", result)
}
