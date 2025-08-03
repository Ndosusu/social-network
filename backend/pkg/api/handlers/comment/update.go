package handlers_comment

import (
	"net/http"
	"social-network/pkg/utils"
)

func UpdateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Comment updated successfully")
}
