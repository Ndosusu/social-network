package handlers_auth

import (
	"net/http"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
)

func FormatDate(day, month, year string) string {
	if len(day) == 1 {
		day = "0" + day
	}
	if len(month) == 1 {
		month = "0" + month
	}
	return year + "-" + month + "-" + day
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidateMethod(w, r, http.MethodPost) {
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form data", nil)
		return
	}

	// Extract form data
	registrationData := map[string]any{
		"FirstName": r.FormValue("FirstName"),
		"LastName":  r.FormValue("LastName"),
		"Mail":      r.FormValue("Mail"),
		"Password":  r.FormValue("Password"),
		"RPassword": r.FormValue("RPassword"),
		"Day":       r.FormValue("Day"),
		"Month":     r.FormValue("Month"),
		"Year":      r.FormValue("Year"),
		"Nickname":  r.FormValue("Nickname"),
		"About":     r.FormValue("About"),
	}

	// Handle image upload if present
	file, header, err := r.FormFile("Avatar")
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

		registrationData["Avatar"] = imagePath
	} else if err != http.ErrMissingFile {
		utils.JSONResponse(w, http.StatusBadRequest, "Error processing image file", nil)
		return
	}

	// Basic validation
	if registrationData["FirstName"] == "" || registrationData["LastName"] == "" ||
		registrationData["Mail"] == "" || registrationData["Password"] == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing required fields", nil)
		return
	}

	if registrationData["Password"] != registrationData["RPassword"] {
		utils.JSONResponse(w, http.StatusBadRequest, "Passwords do not match", nil)
		return
	}

	userData := map[string]any{
		"email":      registrationData["Mail"],
		"first_name": registrationData["FirstName"],
		"last_name":  registrationData["LastName"],
		"password":   registrationData["Password"],
		"date_birth": FormatDate(registrationData["Day"].(string), registrationData["Month"].(string), registrationData["Year"].(string)),
		"avatar":     registrationData["Avatar"],
		"nickname":   registrationData["Nickname"],
		"about":      registrationData["About"],
	}

	var db models.DB
	db.OpenConn()
	defer db.CloseConn()

	udb := user.New(&db)
	result, err := udb.InsertUser(userData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create user account", nil)
		return
	}

	session := result.Result.(models.Session)

	utils.JSONResponse(w, http.StatusCreated, "Registration successful", map[string]any{
		"session_uuid": session.Uuid,
	})
}
