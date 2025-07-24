package api

import (
	"net/http"
	"social-network/pkg/api/handlers"
	"social-network/pkg/api/middleware"
)

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth routes
	mux.HandleFunc("POST /auth/register", middleware.Cors(handlers.RegisterHandler))
	mux.HandleFunc("POST /auth/login", middleware.Cors(handlers.LoginHandler))
	mux.HandleFunc("POST /auth/logout", middleware.Cors(handlers.LogoutHandler))
	mux.HandleFunc("OPTIONS /auth/register", middleware.Cors(handlers.RegisterHandler))
	mux.HandleFunc("OPTIONS /auth/login", middleware.Cors(handlers.LoginHandler))

	// User routes
	/* mux.HandleFunc("GET /user/profile", middleware.Cors(handlers.UserProfileHandler))
	mux.HandleFunc("POST /user/follow", middleware.Cors(handlers.FollowUserHandler))
	mux.HandleFunc("OPTIONS /user/profile", middleware.Cors(handlers.UserProfileHandler))
	mux.HandleFunc("OPTIONS /user/follow", middleware.Cors(handlers.FollowUserHandler)) */

	// Post routes
	mux.HandleFunc("POST /posts", middleware.Cors(handlers.CreatePostHandler))
	//mux.HandleFunc("PUT /posts", middleware.Cors(handlers.UpdatePostHandler))
	mux.HandleFunc("DELETE /posts", middleware.Cors(handlers.DeletePostHandler))
	mux.HandleFunc("GET /posts/image", middleware.Cors(handlers.ServeImageHandler))

	mux.HandleFunc("GET /feed/global", middleware.Cors(handlers.GlobalFeedHandler))
	mux.HandleFunc("GET /feed/follow", middleware.Cors(handlers.FollowFeedHandler))
	mux.HandleFunc("GET /feed/group", middleware.Cors(handlers.GroupFeedHandler))

	/* 	mux.HandleFunc("GET /posts", middleware.Cors(handlers.PostsHandler))

	   	mux.HandleFunc("OPTIONS /posts", middleware.Cors(handlers.PostsHandler)) */
	mux.HandleFunc("OPTIONS /posts/image", middleware.Cors(handlers.ServeImageHandler))
	mux.HandleFunc("OPTIONS /feed/global", middleware.Cors(handlers.GlobalFeedHandler))
	mux.HandleFunc("OPTIONS /feed/follow", middleware.Cors(handlers.FollowFeedHandler))
	mux.HandleFunc("OPTIONS /feed/group", middleware.Cors(handlers.GroupFeedHandler))

	/* // Comment routes
	mux.HandleFunc("GET /comments", middleware.Cors(handlers.CommentsHandler))
	mux.HandleFunc("GET /comments/single", middleware.Cors(handlers.CommentHandler))
	mux.HandleFunc("POST /comments", middleware.Cors(handlers.CommentsHandler))
	mux.HandleFunc("DELETE /comments", middleware.Cors(handlers.DeleteCommentHandler))
	mux.HandleFunc("OPTIONS /comments", middleware.Cors(handlers.CommentsHandler))
	mux.HandleFunc("OPTIONS /comments/single", middleware.Cors(handlers.CommentHandler))

	// Group routes
	mux.HandleFunc("GET /groups", middleware.Cors(handlers.GroupHandler))
	mux.HandleFunc("POST /groups", middleware.Cors(handlers.GroupHandler))
	mux.HandleFunc("POST /groups/create", middleware.Cors(handlers.CreateGroupHandler))
	mux.HandleFunc("POST /groups/join", middleware.Cors(handlers.RequestJoinGroupHandler))
	mux.HandleFunc("POST /groups/invite", middleware.Cors(handlers.InviteToGroupHandler))
	mux.HandleFunc("OPTIONS /groups", middleware.Cors(handlers.GroupHandler))

	// Chat routes - avec CORS ajouté
	mux.HandleFunc("GET /chat", middleware.Cors(handlers.ChatHandler))
	mux.HandleFunc("POST /chat", middleware.Cors(handlers.ChatHandler))
	mux.HandleFunc("OPTIONS /chat", middleware.Cors(handlers.ChatHandler))

	// Notification routes
	mux.HandleFunc("GET /notifications", middleware.Cors(handlers.NotificationHandler))
	mux.HandleFunc("POST /notifications", middleware.Cors(handlers.NotificationHandler))
	mux.HandleFunc("OPTIONS /notifications", middleware.Cors(handlers.NotificationHandler))
	*/
	// Route par défaut
	mux.HandleFunc("GET /{$}", middleware.Cors(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to the Social Network API"}`))
	}))

	return mux
}
