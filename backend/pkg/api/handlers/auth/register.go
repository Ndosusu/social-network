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

	// Handle avatar file
	avatarPath := "default-avatar.png"
	if _, fileHeader, err := r.FormFile("Avatar"); err == nil {
		avatarPath = fileHeader.Filename
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
		"nickname":   registrationData["Nickname"],
		"avatar":     avatarPath,
		"about":      registrationData["About"],
	}

	var db models.DB
	db.OpenConn()
	udb := user.New(&db)
	result, err := udb.InsertUser(userData)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "Failed to create user account", nil)
		return
	}
	db.CloseConn()

	session := result.Result.(models.Session)

	utils.JSONResponse(w, http.StatusCreated, "Registration successful", map[string]any{
		"session_uuid": session.Uuid,
	})
}
