package handlers_post

import (
	"fmt"
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	"social-network/pkg/utils"
)

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodDelete) {
		return
	}

	data := utils.JSONDecode(w, r)
	postID, ok := data["id"].(int)
	if !ok || postID <= 0 {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing post ID", nil)
		return
	}

	var db models.DB
	db.OpenConn()
	pdb := post.New(&db)
	result, err := pdb.DeletePost(map[string]any{"id": postID})
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Database connection failed", nil)
		return
	}
	fmt.Println(result)

	utils.JSONResponse(w, http.StatusOK, "Post deleted successfully", result)
}
