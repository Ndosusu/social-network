package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Success response with JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Registration validation successful",
		"data": map[string]interface{}{
			"status":    "registered",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
	})

	// registrationMap := make(map[string]any) // Initialize a map to hold registration data

	// names := [3]string{"nickname", "first_name", "last_name"}
	// for _, name := range names { // Iterate over the required fields, check if the form value is present and not empty
	// 	value := r.FormValue(name)
	// 	if value == "" {
	// 		http.Error(w, "Missing "+name, http.StatusBadRequest)
	// 		return
	// 	}
	// 	registrationMap[name] = value // Store the value in the map
	// }

	// ageStr := r.FormValue("age")
	// if ageStr == "" {
	// 	http.Error(w, "Missing age", http.StatusBadRequest)
	// 	return
	// }
	// age, err := strconv.Atoi(ageStr)
	// if err != nil || age <= 0 {
	// 	http.Error(w, "Invalid age", http.StatusBadRequest)
	// 	return
	// }
	// registrationMap["age"] = age

	// email := r.FormValue("email")
	// if match, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email); err != nil || !match {
	// 	http.Error(w, "Invalid email format", http.StatusBadRequest)
	// 	return
	// } else {
	// 	registrationMap["email"] = email
	// }

	// password := r.FormValue("password")
	// if password == "" || len(password) < 6 {
	// 	http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
	// 	return
	// }
	// registrationMap["password"] = password

	// // Create a new user
	// user := &models.User{
	// 	Email:       registrationMap["email"].(string),
	// 	Password:    registrationMap["password"].(string),
	// 	FirstName:   registrationMap["first_name"].(string),
	// 	LastName:    registrationMap["last_name"].(string),
	// 	NickName:    registrationMap["nickname"].(string),
	// 	DateBirth:   registrationMap["age"].(int),
	// 	Avatar:      "default_avatar.png",
	// 	About:       "", //TODO
	// 	PrivateMode: false,
	// }

	// // Open database connection
	// db := &models.BDD{}
	// db.OpenConn()
	// defer db.CloseConn()

	// // Set the user's UUID
	// if err := user.Save(db.Conn); err != nil {
	// 	http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// //Success response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "User created successfully",
	// 	"user_id": user.UUID,
	// })
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow POST method
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Success response with proper JSON structure
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Login validation successful",
		"data": map[string]interface{}{
			"status":    "authenticated",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
	})

	// email := r.FormValue("email")
	// if email == "" {
	// 	http.Error(w, "Missing email", http.StatusBadRequest)
	// 	return
	// }

	// password := r.FormValue("password")
	// if password == "" {
	// 	http.Error(w, "Missing password", http.StatusBadRequest)
	// 	return
	// }

	// // Open database connection
	// db := &models.BDD{}
	// db.OpenConn()
	// defer db.CloseConn()

	// //
	// user, err := models.GetUserByEmail(db.Conn, email)
	// if err != nil {
	// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	// 	return
	// }

	// // Check password
	// if !user.CheckPassword(password) {
	// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	// 	return
	// }

	// //Create a session cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "user_session",
	// 	Value:    user.UUID,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteLaxMode,
	// 	Expires:  time.Now().Add(24 * time.Hour),
	// })

	// // Success response
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"message": "Login successful",
	// 	"user": map[string]interface{}{
	// 		"id":         user.UUID,
	// 		"email":      user.Email,
	// 		"first_name": user.FirstName,
	// 		"last_name":  user.LastName,
	// 		"nick_name":  user.NickName,
	// 	},
	// })
}

// func LogoutHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Delete the session cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "user_session",
// 		Value:    "",
// 		Path:     "/",
// 		HttpOnly: true,
// 		Secure:   true,
// 		SameSite: http.SameSiteLaxMode,
// 		Expires:  time.Now().Add(-time.Hour),
// 	})

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Logout successful",
// 	})
// }
