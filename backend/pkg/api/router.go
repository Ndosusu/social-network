package api

import (
	"net/http"
	"social-network/config"
	"social-network/pkg/api/handlers"
	m "social-network/pkg/api/middleware"
)

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/"+config.PicPath+"/", http.StripPrefix("/"+config.PicPath+"/", http.FileServer(http.Dir("./"+config.PicPath))))

	// Auth routes
	mux.HandleFunc("POST /auth/register", m.Cors(handlers.RegisterHandler))
	mux.HandleFunc("POST /auth/login", m.Cors(handlers.LoginHandler))
	mux.HandleFunc("POST /auth/logout", m.Cors(handlers.LogoutHandler))
	mux.HandleFunc("OPTIONS /auth/register", m.Cors(handlers.RegisterHandler))
	mux.HandleFunc("OPTIONS /auth/login", m.Cors(handlers.LoginHandler))

	// User routes
	/* mux.HandleFunc("GET /user/profile", m.Cors(handlers.UserProfileHandler))
	mux.HandleFunc("POST /user/follow", m.Cors(handlers.FollowUserHandler))
	mux.HandleFunc("OPTIONS /user/profile", m.Cors(handlers.UserProfileHandler))
	mux.HandleFunc("OPTIONS /user/follow", m.Cors(handlers.FollowUserHandler)) */

	// Post routes
	mux.HandleFunc("POST /posts", m.Cors(handlers.CreatePostHandler))
	//mux.HandleFunc("PUT /posts", m.Cors(handlers.UpdatePostHandler))
	mux.HandleFunc("DELETE /posts", m.Cors(handlers.DeletePostHandler))
	mux.HandleFunc("GET /posts/image", m.Cors(handlers.ServeImageHandler))

	mux.HandleFunc("POST /feed/global", m.Cors(handlers.GlobalFeedHandler))
	mux.HandleFunc("POST /feed/follow", m.Cors(handlers.FollowFeedHandler))
	mux.HandleFunc("POST /feed/group", m.Cors(handlers.GroupFeedHandler))

	//mux.HandleFunc("GET /posts", m.Cors(handlers.PostsHandler))

	mux.HandleFunc("OPTIONS /posts", m.Cors(handlers.CreatePostHandler))
	mux.HandleFunc("OPTIONS /posts/image", m.Cors(handlers.ServeImageHandler))
	mux.HandleFunc("OPTIONS /feed/global", m.Cors(handlers.GlobalFeedHandler))
	mux.HandleFunc("OPTIONS /feed/follow", m.Cors(handlers.FollowFeedHandler))
	mux.HandleFunc("OPTIONS /feed/group", m.Cors(handlers.GroupFeedHandler))

	// Comment routes
	mux.HandleFunc("POST /comments", m.Cors(handlers.CreateCommentHandler))
	//mux.HandleFunc("PUT /comments", m.Cors(handlers.UpdateCommentHandler))
	mux.HandleFunc("DELETE /comments", m.Cors(handlers.DeleteCommentHandler))

	mux.HandleFunc("POST /feed/detail", m.Cors(handlers.CommentsHandler))

	mux.HandleFunc("OPTIONS /comments", m.Cors(handlers.CommentsHandler))
	mux.HandleFunc("OPTIONS /feed/detail", m.Cors(handlers.CommentsHandler))

	// Like routes
	mux.HandleFunc("POST /likes", m.Cors(handlers.CreateLikeHandler))
	mux.HandleFunc("DELETE /likes", m.Cors(handlers.DeleteLikeHandler))
	mux.HandleFunc("OPTIONS /likes", m.Cors(handlers.CreateLikeHandler))

	/* 	// Group routes
	   	mux.HandleFunc("GET /groups", m.Cors(handlers.GroupHandler))
	   	mux.HandleFunc("POST /groups", m.Cors(handlers.GroupHandler))
	   	mux.HandleFunc("POST /groups/create", m.Cors(handlers.CreateGroupHandler))
	   	mux.HandleFunc("POST /groups/join", m.Cors(handlers.RequestJoinGroupHandler))
	   	mux.HandleFunc("POST /groups/invite", m.Cors(handlers.InviteToGroupHandler))
	   	mux.HandleFunc("OPTIONS /groups", m.Cors(handlers.GroupHandler))

	   	// Chat routes - avec CORS ajouté
	   	mux.HandleFunc("GET /chat", m.Cors(handlers.ChatHandler))
	   	mux.HandleFunc("POST /chat", m.Cors(handlers.ChatHandler))
	   	mux.HandleFunc("OPTIONS /chat", m.Cors(handlers.ChatHandler))

	   	// Notification routes
	   	mux.HandleFunc("GET /notifications", m.Cors(handlers.NotificationHandler))
	   	mux.HandleFunc("POST /notifications", m.Cors(handlers.NotificationHandler))
	   	mux.HandleFunc("OPTIONS /notifications", m.Cors(handlers.NotificationHandler))
	*/
	// Route par défaut
	mux.HandleFunc("GET /{$}", m.Cors(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to the Social Network API"}`))
	}))

	return mux
}
