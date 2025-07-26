package handlers_comment

import (
	"net/http"
	"social-network/pkg/db/models"
	post "social-network/pkg/db/models/postco"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	// Parse multipart form data to handle file uploads
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form data", nil)
		return
	}

	// Extract form data
	comData := map[string]any{
		"post_id": r.FormValue("post_id"),
		"message": r.FormValue("message"),
	}

	// Logic to get Author ID
	authorUuid := r.FormValue("author_uuid")
	if authorUuid == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing required field: author_uuid", nil)
		return
	}
	var db models.DB
	db.OpenConn()
	udb := user.New(&db)
	result, err := udb.GetSessionByUuid(map[string]any{"uuid": authorUuid})
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid author UUID or user not found", nil)
		return
	}
	comData["author_id"] = result.Result.(models.Session).UserId

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

		comData["image"] = imagePath
	} else if err != http.ErrMissingFile {
		utils.JSONResponse(w, http.StatusBadRequest, "Error processing image file", nil)
		return
	}

	/* // Validate required fields
	if errMsg := ValidateRequiredFields(postData, []string{"author_id", "message", "privacy_mode"}); errMsg != "" {
		utils.JSONResponse(w, http.StatusBadRequest, errMsg, nil)
		return
	} */

	db.OpenConn()
	pdb := post.New(&db)
	result, err = pdb.InsertComment(comData)
	db.CloseConn()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create comment", nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "Comment created successfully", result)
}
