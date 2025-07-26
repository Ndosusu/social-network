package handlers_group

import (
	"net/http"
	"social-network/pkg/db/models"
	group "social-network/pkg/db/models/group"
	"social-network/pkg/utils"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	data := utils.JSONDecode(w, r)
	sessionUUID, sessionUUIDOk := data["author_uuid"].(string)
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
	gdb := group.New(&db)
	result, err := gdb.InsertGroup(map[string]any{
		"admin_id": sessionUUID,
		"title":    groupTitle,
		"about":    groupAbout,
	})
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create group", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Group created successfully", result.Result)
}
