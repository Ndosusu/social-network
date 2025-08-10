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
		"firstName": r.FormValue("firstName"),
		"lastName":  r.FormValue("lastName"),
		"mail":      r.FormValue("mail"),
		"password":  r.FormValue("password"),
		"rpassword": r.FormValue("rpassword"),
		"day":       r.FormValue("day"),
		"month":     r.FormValue("month"),
		"year":      r.FormValue("year"),
		"nickname":  r.FormValue("nickname"),
		"about":     r.FormValue("about"),
	}
	// Process avatar image
	imageName, err := utils.ImageProcess(r, "avatar")
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if imageName != "" {
		registrationData["avatar"] = imageName
	}

	// Basic validation
	if registrationData["firstName"] == "" || registrationData["lastName"] == "" ||
		registrationData["mail"] == "" || registrationData["password"] == "" {
		utils.JSONResponse(w, http.StatusBadRequest, "Missing required fields", nil)
		return
	}

	if !utils.IsValidEmail(registrationData["mail"].(string)) {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid email format", nil)
		return
	}
	if !utils.IsValidPassword(registrationData["password"].(string)) {
		utils.JSONResponse(w, http.StatusBadRequest, "Password must be at least 8 characters long and contain uppercase, lowercase, digit, and special character", nil)
		return
	}

	if registrationData["password"] != registrationData["rpassword"] {
		utils.JSONResponse(w, http.StatusBadRequest, "Passwords do not match", nil)
		return
	}

	userData := map[string]any{
		"email":      registrationData["mail"],
		"first_name": registrationData["firstName"],
		"last_name":  registrationData["lastName"],
		"password":   registrationData["password"],
		"date_birth": FormatDate(registrationData["day"].(string), registrationData["month"].(string), registrationData["year"].(string)),
		"avatar":     registrationData["avatar"],
		"nickname":   registrationData["nickname"],
		"about":      registrationData["about"],
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
