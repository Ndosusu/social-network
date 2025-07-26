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
	sessionUUID, ok := data["author_uuid"].(string)
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
	gdb := group.New(&db)
	result, err := gdb.DeleteGroup(map[string]any{
		"admin_id": sessionUUID,
		"group_id": groupIDInt,
	})
}
