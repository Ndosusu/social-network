package handlers_group

import (
	"net/http"
	"social-network/pkg/db/models"
	group "social-network/pkg/db/models/group"
	"social-network/pkg/utils"
)

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodDelete) {
		return
	}
	data := utils.JSONDecode(w, r)
	sessionUUID, ok := data["session_uuid"].(string)
	if !ok || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session UUID", nil)
		return
	}
	groupID, ok := data["group_id"].(float64)
	var groupIDInt int
	if ok {
		groupIDInt = int(groupID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing group ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	gdb := group.New(&db)
	result, err := gdb.IsUserGroupAdmin(map[string]any{
		"group_id":     groupIDInt,
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

	result, err = gdb.DeleteGroup(map[string]any{
		"group_id": groupIDInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to delete group", nil)
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Group deleted successfully", result)
}
