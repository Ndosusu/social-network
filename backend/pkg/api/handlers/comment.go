package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/pkg/db/models"
	"strconv"
)

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		// Récupérer les commentaires d'un post
		postIDStr := r.URL.Query().Get("post_id")
		if postIDStr == "" {
			writeErrorResponse(w, http.StatusBadRequest, "Missing post_id parameter")
			return
		}

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid post_id parameter")
			return
		}

		// Vérifier si on veut les détails de l'auteur
		withAuthor := r.URL.Query().Get("with_author")

		db, err := getDBConnection()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
			return
		}
		defer db.Close()

		dbInstance := &models.DB{Conn: db}

		if withAuthor == "true" {
			result := dbInstance.SelectCommentsByPostIdWithAuthor(map[string]any{"post_id": postID})
			comments, ok := result.Result.([]models.CommentWithAuthor)
			if !ok {
				writeErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve comments")
				return
			}
			writeSuccessResponse(w, http.StatusOK, "", comments)
		} else {
			result := dbInstance.SelectCommentsByPostId(map[string]any{"post_id": postID})
			comments, ok := result.Result.([]models.Comment)
			if !ok {
				writeErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve comments")
				return
			}
			writeSuccessResponse(w, http.StatusOK, "", comments)
		}
		return
	}

	if r.Method == http.MethodPost {
		// Créer un nouveau commentaire
		var commentData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&commentData); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
			return
		}

		// Valider les champs obligatoires
		requiredFields := []string{"author_id", "post_id", "message"}
		for _, field := range requiredFields {
			if commentData[field] == nil {
				writeErrorResponse(w, http.StatusBadRequest, "Missing required field: "+field)
				return
			}
		}

		// Définir les valeurs par défaut
		if commentData["image"] == nil {
			commentData["image"] = ""
		}
		if commentData["group_id"] == nil {
			commentData["group_id"] = 0
		}

		db, err := getDBConnection()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
			return
		}
		defer db.Close()

		dbInstance := &models.DB{Conn: db}
		result := dbInstance.InsertComment(commentData)

		comment, ok := result.Result.(models.Comment)
		if !ok || comment.Id == 0 {
			writeErrorResponse(w, http.StatusInternalServerError, "Failed to create comment")
			return
		}

		writeSuccessResponse(w, http.StatusCreated, "Comment created successfully", comment)
		return
	}
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodDelete) {
		return
	}

	commentIDStr := r.URL.Query().Get("id")
	if commentIDStr == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Missing comment ID")
		return
	}

	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	// Vérifier si le commentaire existe avant de le supprimer
	dbInstance := &models.DB{Conn: db}
	result := dbInstance.SelectCommentById(map[string]any{"id": commentID})
	comment, ok := result.Result.(models.Comment)
	if !ok || comment.Id == 0 {
		writeErrorResponse(w, http.StatusNotFound, "Comment not found")
		return
	}

	// Supprimer le commentaire
	deleteResult := dbInstance.DeleteComment(map[string]any{"id": commentID})
	if deleteResult.Result == 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Failed to delete comment")
		return
	}

	writeSuccessResponse(w, http.StatusOK, "Comment deleted successfully", nil)
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r, http.MethodGet) {
		return
	}

	commentIDStr := r.URL.Query().Get("id")
	if commentIDStr == "" {
		writeErrorResponse(w, http.StatusBadRequest, "Missing comment ID")
		return
	}

	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	db, err := getDBConnection()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Database connection failed")
		return
	}
	defer db.Close()

	dbInstance := &models.DB{Conn: db}
	result := dbInstance.SelectCommentById(map[string]any{"id": commentID})
	comment, ok := result.Result.(models.Comment)
	if !ok || comment.Id == 0 {
		writeErrorResponse(w, http.StatusNotFound, "Comment not found")
		return
	}

	writeSuccessResponse(w, http.StatusOK, "", comment)
}
