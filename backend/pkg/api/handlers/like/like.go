package handlers_like

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	data := utils.JSONDecode(w, r)

	sessionUUID, sessionUUIDOk := data["author_uuid"].(string)
	if !sessionUUIDOk || sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing session", nil)
		return
	}
	var db models.DB
	db.OpenConn()
	udb := user.New(&db)
	result, err := udb.GetSessionByUuid(map[string]any{"uuid": sessionUUID})
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid author UUID or user not found", nil)
		return
	}
	userID := result.Result.(models.Session).UserId

	postID, postIDOk := data["post_id"].(float64)
	var postIDInt int
	if postIDOk {
		postIDInt = int(postID)
	}
	commentID, commentIDOk := data["comment_id"].(float64)
	var commentIDInt int
	if commentIDOk {
		commentIDInt = int(commentID)
	}

	alreadyLiked, alreadyLikedOk := data["already_liked"].(bool)
	if !alreadyLikedOk {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing already_liked flag", nil)
		return
	}

	likeID, likeIDOk := data["like_id"].(float64)
	var likeIDInt int
	if likeIDOk {
		likeIDInt = int(likeID)
	} else {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid or missing like ID", nil)
		return
	}

	pdb := post.New(&db)
	if !alreadyLiked {
		result, err = pdb.InsertLike(map[string]any{
			"post_id":    postIDInt,
			"comment_id": commentIDInt,
			"user_id":    userID,
		})
	} else {
		result, err = pdb.DeleteLike(map[string]any{
			"id": likeIDInt,
		})
	}
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to update like status in database", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Like status changed successfully", result)
}
