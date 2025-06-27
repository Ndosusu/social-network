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

type LoginRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	FirstName string `form:"FirstName"`
	LastName  string `form:"LastName"`
	Mail      string `form:"Mail"`
	Password  string `form:"Password"`
	RPassword string `form:"RPassword"`
	Day       string `form:"Day"`
	Month     string `form:"Month"`
	Year      string `form:"Year"`
	Nickname  string `form:"Nickname"`
	About     string `form:"About"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Parse JSON body for login data
	var loginData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid JSON data",
		})
		return
	}

	// Log the received data for debugging, excluding sensitive information
	mail, mailOk := loginData["Mail"].(string)
	password, passwordOk := loginData["Password"].(string)

	if !mailOk || !passwordOk || mail == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Missing email or password",
		})
		return
	}

	// DB connection
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Verify user credentials
	userModel := &models.DB{Conn: db}
	log.Printf("Attempting authentication for email: %s", mail)

	result := userModel.Authenticate(map[string]any{
		"mail":     mail,
		"password": password,
	})

	// Check if authentication failed
	if result.Result == nil {
		log.Printf("Authentication failed for email: %s - invalid credentials", mail)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid email or password",
		})
		return
	}

	user, ok := result.Result.(models.User)
	if !ok || user.Id == 0 {
		log.Printf("Authentication failed for email: %s - failed to retrieve user data", mail)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid email or password",
		})
		return
	}

	log.Printf("Authentication successful for user: %s (ID: %d, UUID: %s)", user.Email, user.Id, user.Uuid)

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Login successful",
		"data": map[string]interface{}{
			"user_id":   user.Id,
			"uuid":      user.Uuid,
			"email":     user.Email,
			"nickname":  user.Nickname,
			"status":    "authenticated",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Parse multipart form data
	err := r.ParseMultipartForm(32 << 20) // 32MB max
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to parse form data",
		})
		return
	}

	// Récupérer les données du formulaire
	registrationData := map[string]interface{}{
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

	// Log des données reçues pour debug
	logData := make(map[string]interface{})
	for k, v := range registrationData {
		if k != "Password" && k != "RPassword" {
			logData[k] = v
		} else {
			logData[k] = "[HIDDEN]"
		}
	}
	log.Printf("Registration data received: %+v", logData)

	// Gérer le fichier avatar s'il existe
	avatarPath := "default-avatar.png" // valeur par défaut
	file, fileHeader, err := r.FormFile("Avatar")
	if err == nil {
		defer file.Close()
		log.Printf("Avatar file received: %s, size: %d", fileHeader.Filename, fileHeader.Size)
		avatarPath = fileHeader.Filename
	}

	// Validation basique
	if registrationData["FirstName"] == "" || registrationData["LastName"] == "" ||
		registrationData["Mail"] == "" || registrationData["Password"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Missing required fields",
		})
		return
	}

	// Vérifier que les mots de passe correspondent
	if registrationData["Password"] != registrationData["RPassword"] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Passwords do not match",
		})
		return
	}

	// Connexion à la base de données
	db, err := getDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Database connection failed",
		})
		return
	}
	defer db.Close()

	// Préparer les données pour l'insertion en base
	userModel := &models.DB{Conn: db}

	// Formatter la date de naissance (assurer le format YYYY-MM-DD)
	day := registrationData["Day"].(string)
	month := registrationData["Month"].(string)
	year := registrationData["Year"].(string)

	// Ajouter des zéros si nécessaire
	if len(day) == 1 {
		day = "0" + day
	}
	if len(month) == 1 {
		month = "0" + month
	}

	dateBirth := year + "-" + month + "-" + day

	userData := map[string]any{
		"email":      registrationData["Mail"],
		"first_name": registrationData["FirstName"],
		"last_name":  registrationData["LastName"],
		"password":   registrationData["Password"],
		"date_birth": dateBirth,
		"nickname":   registrationData["Nickname"],
		"avatar":     avatarPath,
		"about":      registrationData["About"],
	}

	// Insérer l'utilisateur en base de données
	result := userModel.InsertUser(userData)

	if result.Result == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to create user account",
		})
		return
	}

	// Extraire les données utilisateur du résultat
	user, ok := result.Result.(models.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to retrieve user data",
		})
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Registration successful",
		"data": map[string]interface{}{
			"user_id":   user.Id,
			"uuid":      user.Uuid,
			"email":     user.Email,
			"nickname":  user.Nickname,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Logout logic here...
}

func getDBConnection() (*sql.DB, error) {
	_, filename, _, _ := runtime.Caller(1)
	baseDir, _ := strings.CutSuffix(filename, "pkg/api/handlers/auth.go")
	dbPath := baseDir + config.DBPath + "/" + config.DBName
	return sql.Open("sqlite3", dbPath)
}
