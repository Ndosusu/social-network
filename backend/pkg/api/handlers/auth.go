// Package handlers provides HTTP handlers for the social network API
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"social-network/config"
	"social-network/pkg/db/models"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Helper functions to reduce code duplication
func writeJSONResponse(w http.ResponseWriter, statusCode int, data map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	writeJSONResponse(w, statusCode, map[string]any{
		"success": false,
		"error":   message,
	})
}

func writeSuccessResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	response := map[string]any{
		"success": true,
		"message": message,
	}
	if data != nil {
		response["data"] = data
	}
	writeJSONResponse(w, statusCode, response)
}

func validateMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return false
	}
	return true
}

func formatDate(day, month, year string) string {
	if len(day) == 1 {
		day = "0" + day
	}
	if len(month) == 1 {
		month = "0" + month
	}
	return year + "-" + month + "-" + day
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodPost) {
		return
	}

	var loginData map[string]any
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON data")
		return
	}

	mail, mailOk := loginData["Mail"].(string)
	password, passwordOk := loginData["Password"].(string)

	if !mailOk || !passwordOk || mail == "" || password == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Missing email or password")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	userModel := &models.DB{Conn: db}
	result := userModel.Authenticate(map[string]any{
		"mail":     mail,
		"password": password,
	})

	if result.Result == nil {
		writeErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	user, ok := result.Result.(models.User)
	if !ok || user.Id == 0 {
		writeErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	writeSuccessResponse(w, http.StatusOK, "Login successful", map[string]any{
		"user_id":   user.Id,
		"uuid":      user.Uuid,
		"email":     user.Email,
		"nickname":  user.Nickname,
		"status":    "authenticated",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodPost) {
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Failed to parse form data")
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
		writeErrorResponse(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	if registrationData["Password"] != registrationData["RPassword"] {
		writeErrorResponse(w, http.StatusBadRequest, "Passwords do not match")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	userModel := &models.DB{Conn: db}
	userData := map[string]any{
		"email":      registrationData["Mail"],
		"first_name": registrationData["FirstName"],
		"last_name":  registrationData["LastName"],
		"password":   registrationData["Password"],
		"date_birth": formatDate(registrationData["Day"].(string), registrationData["Month"].(string), registrationData["Year"].(string)),
		"nickname":   registrationData["Nickname"],
		"avatar":     avatarPath,
		"about":      registrationData["About"],
	}

	result := userModel.InsertUser(userData)
	if result.Result == 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to create user account")
		return
	}

	user, ok := result.Result.(models.User)
	if !ok {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve user data")
		return
	}

	writeSuccessResponse(w, http.StatusCreated, "Registration successful", map[string]any{
		"user_id":   user.Id,
		"uuid":      user.Uuid,
		"email":     user.Email,
		"nickname":  user.Nickname,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if !validateMethod(w, r, http.MethodPost) {
		return
	}

	writeSuccessResponse(w, http.StatusOK, "Logout successful", nil)
}

func getDBConnection() (*sql.DB, error) {
	_, filename, _, _ := runtime.Caller(1)
	var baseDir string

	if strings.Contains(filename, "auth.go") {
		baseDir, _ = strings.CutSuffix(filename, "pkg/api/handlers/auth.go")
	} else if strings.Contains(filename, "post.go") {
		baseDir, _ = strings.CutSuffix(filename, "pkg/api/handlers/post.go")
	} else {

		baseDir, _ = strings.CutSuffix(filename, "pkg/api/handlers/auth.go")
	}

	dbPath := baseDir + config.DBPath + "/" + config.DBName
	log.Printf("Attempting to connect to database at: %s", dbPath)

	return sql.Open("sqlite3", dbPath)
}
