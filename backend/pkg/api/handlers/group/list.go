package handlers_group

import (
	"net/http"
	"social-network/pkg/db/models"
	group "social-network/pkg/db/models/group"
	"social-network/pkg/utils"
)

func ListGroupsHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodGet) {
		return
	}

	data := utils.JSONDecode(w, r)
	sessionUUID := data["session_uuid"].(string)

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	gdb := group.New(&db)
	result, err := gdb.GetUserGroupList(map[string]any{
		"session_uuid": sessionUUID,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to retrieve groups", nil)
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Groups retrieved successfully", result)

}
