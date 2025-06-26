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

	mux.HandleFunc("POST /auth/register", corsMiddleware(handlers.RegisterHandler))
	mux.HandleFunc("POST /auth/login", corsMiddleware(handlers.LoginHandler))
	mux.HandleFunc("POST /auth/logout", corsMiddleware(handlers.LogoutHandler))
	mux.HandleFunc("OPTIONS /auth/register", corsMiddleware(handlers.RegisterHandler))
	mux.HandleFunc("OPTIONS /auth/login", corsMiddleware(handlers.LoginHandler))

	mux.HandleFunc("GET /user/profile", corsMiddleware(handlers.UserProfileHandler))
	mux.HandleFunc("POST /user/follow", corsMiddleware(handlers.FollowUserHandler))
	mux.HandleFunc("OPTIONS /user/profile", corsMiddleware(handlers.UserProfileHandler))
	mux.HandleFunc("OPTIONS /user/follow", corsMiddleware(handlers.FollowUserHandler))

	// Post routes
	mux.HandleFunc("POST /posts", handlers.CreatePostHandler)
	mux.HandleFunc("GET /posts", handlers.PostHandler)

	// Comment routes
	mux.HandleFunc("GET /comments", handlers.CommentsHandler)
	mux.HandleFunc("POST /comments", handlers.CommentsHandler)

	// Group routes
	mux.HandleFunc("GET /groups", handlers.GroupHandler)
	mux.HandleFunc("POST /groups", handlers.GroupHandler)
	mux.HandleFunc("POST /groups/create", handlers.CreateGroupHandler)
	mux.HandleFunc("POST /groups/join", handlers.RequestJoinGroupHandler)
	mux.HandleFunc("POST /groups/invite", handlers.InviteToGroupHandler)

	// Chat routes
	mux.HandleFunc("GET /chat", handlers.ChatHandler)
	mux.HandleFunc("POST /chat", handlers.ChatHandler)

	// Notification routes
	mux.HandleFunc("GET /notifications", handlers.NotificationHandler)
	mux.HandleFunc("POST /notifications", handlers.NotificationHandler)

	// Route par défaut - avec pattern exact
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to the Social Network API"}`))
	})

	return mux
}
