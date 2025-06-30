package api

import (
	"net/http"
	"social-network/pkg/api/handlers"
)

// Middleware pour gérer CORS globalement
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Headers CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// for Options requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// if not OPTIONS, continue with the next handler
		next.ServeHTTP(w, r)
	}
}

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth routes
	mux.HandleFunc("POST /auth/register", corsMiddleware(handlers.RegisterHandler))
	mux.HandleFunc("POST /auth/login", corsMiddleware(handlers.LoginHandler))
	mux.HandleFunc("POST /auth/logout", corsMiddleware(handlers.LogoutHandler))
	mux.HandleFunc("OPTIONS /auth/register", corsMiddleware(handlers.RegisterHandler))
	mux.HandleFunc("OPTIONS /auth/login", corsMiddleware(handlers.LoginHandler))

	// User routes
	mux.HandleFunc("GET /user/profile", corsMiddleware(handlers.UserProfileHandler))
	mux.HandleFunc("POST /user/follow", corsMiddleware(handlers.FollowUserHandler))
	mux.HandleFunc("OPTIONS /user/profile", corsMiddleware(handlers.UserProfileHandler))
	mux.HandleFunc("OPTIONS /user/follow", corsMiddleware(handlers.FollowUserHandler))

	// Post routes
	mux.HandleFunc("POST /posts", corsMiddleware(handlers.CreatePostHandler))
	mux.HandleFunc("GET /posts", corsMiddleware(handlers.PostHandler))
	mux.HandleFunc("OPTIONS /posts", corsMiddleware(handlers.PostHandler))

	// Comment routes
	mux.HandleFunc("GET /comments", corsMiddleware(handlers.CommentsHandler))
	mux.HandleFunc("POST /comments", corsMiddleware(handlers.CommentsHandler))
	mux.HandleFunc("OPTIONS /comments", corsMiddleware(handlers.CommentsHandler))

	// Group routes
	mux.HandleFunc("GET /groups", corsMiddleware(handlers.GroupHandler))
	mux.HandleFunc("POST /groups", corsMiddleware(handlers.GroupHandler))
	mux.HandleFunc("POST /groups/create", corsMiddleware(handlers.CreateGroupHandler))
	mux.HandleFunc("POST /groups/join", corsMiddleware(handlers.RequestJoinGroupHandler))
	mux.HandleFunc("POST /groups/invite", corsMiddleware(handlers.InviteToGroupHandler))
	mux.HandleFunc("OPTIONS /groups", corsMiddleware(handlers.GroupHandler))

	// Chat routes - avec CORS ajouté
	mux.HandleFunc("GET /chat", corsMiddleware(handlers.ChatHandler))
	mux.HandleFunc("POST /chat", corsMiddleware(handlers.ChatHandler))
	mux.HandleFunc("OPTIONS /chat", corsMiddleware(handlers.ChatHandler))

	// Notification routes
	mux.HandleFunc("GET /notifications", corsMiddleware(handlers.NotificationHandler))
	mux.HandleFunc("POST /notifications", corsMiddleware(handlers.NotificationHandler))
	mux.HandleFunc("OPTIONS /notifications", corsMiddleware(handlers.NotificationHandler))

	// Route par défaut
	mux.HandleFunc("GET /{$}", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to the Social Network API"}`))
	}))

	return mux
}
