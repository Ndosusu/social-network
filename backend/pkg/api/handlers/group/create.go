package handlers_group

import (
	"net/http"
	"social-network/pkg/db/models"
	group "social-network/pkg/db/models/group"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)
	sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session UUID", nil)
		return
	}
	groupTitle, groupTitleOk := data["group_title"].(string)
	if !groupTitleOk || groupTitle == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing group title", nil)
		return
	}
	groupAbout, groupAboutOk := data["group_about"].(string)
	if !groupAboutOk || groupAbout == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing group about", nil)
		return
	}
	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	result, err := udb.GetSessionByUuid(map[string]any{"session_uuid": sessionUUID})
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid author UUID or user not found", nil)
		return
	}
	adminID := result.Result.(models.Session).User.Id

	gdb := group.New(&db)
	result, err = gdb.InsertGroup(map[string]any{
		"admin_id": adminID,
		"title":    groupTitle,
		"about":    groupAbout,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create group", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Group created successfully", result.Result)
}
