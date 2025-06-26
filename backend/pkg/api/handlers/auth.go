package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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

	// Log des données reçues pour debug
	log.Printf("Login data received: %+v", loginData)

	// Validation basique
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

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Login successful",
		"data": map[string]interface{}{
			"mail":      mail,
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

	// Parse multipart form data (pour gérer les fichiers)
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
	log.Printf("Registration data received: %+v", registrationData)

	// Gérer le fichier avatar s'il existe
	file, fileHeader, err := r.FormFile("Avatar")
	if err == nil {
		defer file.Close()
		log.Printf("Avatar file received: %s, size: %d", fileHeader.Filename, fileHeader.Size)
		registrationData["Avatar"] = fileHeader.Filename
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

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Registration successful",
		"data": map[string]interface{}{
			"user_data": registrationData,
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
