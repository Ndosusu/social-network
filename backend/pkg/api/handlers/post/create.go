package handlers_post

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/post"
	rel "social-network/pkg/db/models/relation"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	// Parse multipart form data to handle file uploads
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form data", nil)
		return
	}

	// Extract form data
	postData := map[string]any{
		"message":      r.FormValue("message"),
		"privacy_mode": r.FormValue("privacy_mode"),
	}

	// Logic to get Author ID
	sessionUUID := r.FormValue("session_uuid")
	if sessionUUID == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing required field: session_uuid", nil)
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
	postData["author_id"] = result.Result.(models.Session).User.Id

	// Group ID is optional
	if groupID := r.FormValue("group_id"); groupID != "" {
		postData["group_id"] = groupID
	}

	// Handle image upload if present
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		// Validate image file
		if err := utils.ValidateImageFile(file, header); err != nil {
			utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		// Save the image
		imagePath, err := utils.SaveImageFile(file, header)
		if err != nil {
			utils.JSONResponse(w, http.StatusInternalServerError, "Failed to save image: "+err.Error(), nil)
			return
		}

		postData["image"] = imagePath
	} else if err != http.ErrMissingFile {
		utils.JSONResponse(w, http.StatusBadRequest, "Error processing image file", nil)
		return
	}

	/* // Validate required fields
	if errMsg := ValidateRequiredFields(postData, []string{"author_id", "message", "privacy_mode"}); errMsg != "" {
		utils.JSONResponse(w, http.StatusBadRequest, errMsg, nil)
		return
	} */

	pdb := post.New(&db)
	result, err = pdb.InsertPost(postData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create post", nil)
		return
	}

	// If privacy mode is whitelist, handle followers
	if postData["privacy_mode"] == utils.PRIVACY_MODE_WHITELIST {
		postRels := map[string]any{
			"post_id": result.Result.(models.Post).Id,
		}
		followersIDs := r.PostForm["followers_id"]
		// Add the author ID to the list, he needs to see his own post
		followersIDs = append(followersIDs, postData["author_id"].(string))

		rdb := rel.New(&db)
		for _, followerID := range followersIDs {
			postRels["user_id"] = followerID
			_, err = rdb.InsertPrivacyPostRel(postRels)
			if err != nil {
				utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create post privacy relation: "+err.Error(), nil)
				db.CloseConn()
				return
			}
		}
	}

	utils.JSONResponse(w, http.StatusCreated, "Post created successfully", result)
}
