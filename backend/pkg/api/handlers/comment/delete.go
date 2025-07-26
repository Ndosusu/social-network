package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	comment "social-network/pkg/db/models/comment"
	"social-network/pkg/utils"
)

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodDelete) {
		return
	}

	data := utils.JSONDecode(w, r)
	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}

	comID, comIDOk := data["comment_id"].(float64)
	var comIDInt int
	if comIDOk {
		comIDInt = int(comID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing comment ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	cdb := comment.New(&db)
	result, err := cdb.IsUserCommentAuthor(map[string]any{
		"comment_id":   comIDInt,
		"session_uuid": sessionUUID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, "Comment not found or invalid session", nil)
		return
	}

	canDelete := result.Result.(bool)
	if !canDelete {
		utils.JSONResponse(w, http.StatusForbidden, "You are not the author of this comment", nil)
		return
	}

	result, err = cdb.DeleteComment(map[string]any{
		"comment_id": comIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to delete comment", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Comment deleted successfully", result)
}
