package handlers_search

import (
	"net/http"
	"social-network/pkg/db/models"
	group "social-network/pkg/db/models/group"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodGet) {
		return
	}

	data := utils.JSONDecode(w, r)

	query, queryOk := data["query"].(string)
	if !queryOk || query == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing parameters", nil)
		return
	}
	limit, limitOk := data["limit"].(float64)
	var limitInt int
	if limitOk {
		limitInt = int(limit)
	} else {
		// Default to 10 if limit is not provided or invalid
		limitInt = utils.DEFAULT_LIMIT
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	userResult, err := udb.SearchUser(map[string]any{
		"query": query,
		"limit": limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Search failed", nil)
		return
	}

	gdb := group.New(&db)
	groupResult, err := gdb.SearchGroup(map[string]any{
		"query": query,
		"limit": limitInt,
	})
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Search failed", nil)
		return
	}

	type concatResult struct {
		Users  []models.User
		Groups []models.Group
	}
	allResult := concatResult{
		Users:  userResult.Result.([]models.User),
		Groups: groupResult.Result.([]models.Group),
	}

	utils.JSONResponse(w, http.StatusOK, "Search results", allResult)
}
