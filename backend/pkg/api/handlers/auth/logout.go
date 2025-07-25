package handlers_auth

import (
	"net/http"
	"social-network/pkg/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Logout successful", nil)
}
