package api

import (
	"net/http"
	"social-network/pkg/api/handlers"
)

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Authentication routes
	// mux.HandleFunc("POST /auth/register", handlers.RegisterHandler)
	mux.HandleFunc("POST /auth/login", handlers.LoginHandler)
	// mux.HandleFunc("POST /auth/logout", handlers.LogoutHandler)

	// User routes
	mux.HandleFunc("GET /user/profile", handlers.UserProfileHandler)
	mux.HandleFunc("POST /user/follow", handlers.FollowUserHandler)

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

	// Default route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Social Network"))
	})

	return mux
}
