package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodDelete) {
		return
	}

	data := utils.JSONDecode(w, r)
	comID, ok := data["id"].(int)
	if !ok || comID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing post ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.DeletePost(map[string]any{"id": comID})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Database connection failed", nil)
		db.CloseConn()
		return
	}
	db.CloseConn()

	utils.JSONResponse(w, http.StatusOK, "Comment deleted successfully", result)
}
