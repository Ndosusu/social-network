package handlers_user

import (
	"net/http"
	"social-network/pkg/utils"
)

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPut) {
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form data", nil)
		return
	}

	updateData := map[string]any{}
}

// 3 form : avatar/infos/password
