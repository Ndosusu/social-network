package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"social-network/pkg/db/models"
	user "social-network/pkg/db/models/user"
	"social-network/pkg/utils"
	"strconv"
	"strings"
)

func CheckSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Verify general request method
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "DELETE" {
			contentType := r.Header.Get("Content-Type")

			var db models.DB
			db.OpenConn()
			defer db.CloseConn()
			udb := user.New(&db)

			if strings.HasPrefix(contentType, "application/json") {

				// Decode the JSON body
				data := utils.JSONDecode(w, r)

				// Check if session_uuid is present in the data
				sessionUUID, sessionUUIDOk := data["session_uuid"].(string)
				if !sessionUUIDOk {
					utils.JSONResponse(w, http.StatusBadRequest, "Invalid session", nil)
					return
				}
				if sessionUUID == "" {
					utils.JSONResponse(w, http.StatusUnauthorized, "Missing session", nil)
					return
				}

				// Verify if the session exists and returns it
				sessionResult, err := udb.GetSessionByUuid(map[string]any{"session_uuid": sessionUUID})
				if err != nil {
					utils.JSONResponse(w, http.StatusUnauthorized, "Invalid session", nil)
					return
				}

				// Inject the returned session into the request body
				session := sessionResult.Result.(models.Session)
				data["client_id"] = session.User.Id
				data["client_session"] = session

				// Marshal the data back to JSON
				dataBytes, err := json.Marshal(data)
				if err != nil {
					utils.JSONResponse(w, http.StatusInternalServerError, "Failed to process request", nil)
					return
				}

				fmt.Println(string(dataBytes))

				// Recreate the request body with the updated data
				r.Body = io.NopCloser(bytes.NewBuffer(dataBytes))

				// Serve the next handler with the updated request
				next.ServeHTTP(w, r)

			} else if strings.HasPrefix(contentType, "multipart/form-data") {

				// Parse multipart form data to handle file uploads
				if err := r.ParseMultipartForm(10 << 20); err != nil {
					utils.JSONResponse(w, http.StatusBadRequest, "Failed to parse form", nil)
					return
				}

				// Check if session_uuid is present in the form data
				sessionUUID := r.FormValue("session_uuid")
				if sessionUUID == "" {
					utils.JSONResponse(w, http.StatusUnauthorized, "Missing session", nil)
					return
				}

				// Verify if the session exists and returns it
				sessionResult, err := udb.GetSessionByUuid(map[string]any{"session_uuid": sessionUUID})
				if err != nil {
					utils.JSONResponse(w, http.StatusUnauthorized, "Invalid session", nil)
					return
				}

				// Inject the returned session into the form data
				session := sessionResult.Result.(models.Session)

				// Initialize PostForm if it doesn't exist
				if r.Form == nil {
					r.Form = make(url.Values)
				}

				// Add client data to form (check for conflicts first)
				if r.Form.Get("client_id") != "" {
					utils.JSONResponse(w, http.StatusBadRequest, "client_id should not be provided by client", nil)
					return
				}

				// Inject authenticated user data into form fields
				r.Form.Set("client_id", strconv.Itoa(session.User.Id))
				r.Form.Set("client_session_uuid", sessionUUID)

				// Serve the next handler with the updated request
				next.ServeHTTP(w, r)

			} else {
				// Unsupported content type
				utils.JSONResponse(w, http.StatusBadRequest, "Unsupported content type", nil)
				return
			}
		}
	}
}
