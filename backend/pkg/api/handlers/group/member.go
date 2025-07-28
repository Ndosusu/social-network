package handlers_group

import (
	"net/http"
	"social-network/pkg/utils"
)

func RequestJoinGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	// Logic to request to join a group
}

func InviteToGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	// Logic to invite a user to a group
}
...