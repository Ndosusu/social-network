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
	mux.HandleFunc("POST /auth/logout", m.Cors(m.CheckSession(handlers.LogoutHandler)))

	mux.HandleFunc("OPTIONS /auth/register", m.Cors(handlers.RegisterHandler))
	mux.HandleFunc("OPTIONS /auth/login", m.Cors(handlers.LoginHandler))
	mux.HandleFunc("OPTIONS /auth/logout", m.Cors(handlers.LogoutHandler))

	// User routes
	mux.HandleFunc("POST /profile", m.Cors(m.CheckSession(handlers.UserProfileHandler)))
	mux.HandleFunc("POST /follow", m.Cors(m.CheckSession(handlers.FollowCreateHandler)))
	mux.HandleFunc("DELETE /follow", m.Cors(m.CheckSession(handlers.FollowDeleteHandler)))
	mux.HandleFunc("POST /follow/by", m.Cors(m.CheckSession(handlers.FollowByHandler)))
	mux.HandleFunc("POST /follow/from", m.Cors(m.CheckSession(handlers.FollowFromHandler)))

	mux.HandleFunc("OPTIONS /profile", m.Cors(handlers.UserProfileHandler))
	mux.HandleFunc("OPTIONS /follow", m.Cors(handlers.FollowCreateHandler))
	mux.HandleFunc("OPTIONS /follow/by", m.Cors(handlers.FollowByHandler))
	mux.HandleFunc("OPTIONS /follow/from", m.Cors(handlers.FollowFromHandler))

	// Post routes
	mux.HandleFunc("POST /posts", m.Cors(m.CheckSession(handlers.CreatePostHandler)))
	mux.HandleFunc("PUT /posts", m.Cors(m.CheckSession(handlers.UpdatePostHandler)))
	mux.HandleFunc("DELETE /posts", m.Cors(m.CheckSession(handlers.DeletePostHandler)))
	mux.HandleFunc("GET /posts/image", m.Cors(handlers.ServeImageHandler))

	mux.HandleFunc("POST /feed/global", m.Cors(m.CheckSession(handlers.GlobalFeedHandler)))
	mux.HandleFunc("POST /feed/follow", m.Cors(m.CheckSession(handlers.FollowFeedHandler)))
	mux.HandleFunc("POST /feed/group", m.Cors(m.CheckSession(handlers.GroupFeedHandler)))
	mux.HandleFunc("POST /feed/profile", m.Cors(m.CheckSession(handlers.ProfileFeedHandler)))

	mux.HandleFunc("OPTIONS /posts", m.Cors(handlers.CreatePostHandler))
	mux.HandleFunc("OPTIONS /posts/image", m.Cors(handlers.ServeImageHandler)) // Useful ?
	mux.HandleFunc("OPTIONS /feed/global", m.Cors(handlers.GlobalFeedHandler))
	mux.HandleFunc("OPTIONS /feed/follow", m.Cors(handlers.FollowFeedHandler))
	mux.HandleFunc("OPTIONS /feed/group", m.Cors(handlers.GroupFeedHandler))
	mux.HandleFunc("OPTIONS /feed/profile", m.Cors(handlers.ProfileFeedHandler))

	// Comment routes
	mux.HandleFunc("POST /comments", m.Cors(m.CheckSession(handlers.CreateCommentHandler)))
	mux.HandleFunc("PUT /comments", m.Cors(m.CheckSession(handlers.UpdateCommentHandler)))
	mux.HandleFunc("DELETE /comments", m.Cors(m.CheckSession(handlers.DeleteCommentHandler)))

	mux.HandleFunc("POST /feed/detail", m.Cors(m.CheckSession(handlers.CommentsHandler)))

	mux.HandleFunc("OPTIONS /comments", m.Cors(handlers.CommentsHandler))
	mux.HandleFunc("OPTIONS /feed/detail", m.Cors(handlers.CommentsHandler))

	// Like routes
	mux.HandleFunc("POST /likes", m.Cors(m.CheckSession(handlers.CreateLikeHandler)))
	mux.HandleFunc("DELETE /likes", m.Cors(m.CheckSession(handlers.DeleteLikeHandler)))

	mux.HandleFunc("OPTIONS /likes", m.Cors(handlers.CreateLikeHandler))

	// Group routes
	mux.HandleFunc("POST /groups", m.Cors(m.CheckSession(handlers.CreateGroupHandler)))
	mux.HandleFunc("DELETE /groups", m.Cors(m.CheckSession(handlers.DeleteGroupHandler)))
	mux.HandleFunc("POST /groups/list", m.Cors(m.CheckSession(handlers.ListGroupsHandler)))

	mux.HandleFunc("OPTIONS /groups", m.Cors(handlers.CreateGroupHandler))
	mux.HandleFunc("OPTIONS /groups/list", m.Cors(handlers.ListGroupsHandler))

	// Search routes
	mux.HandleFunc("POST /search", m.Cors(m.CheckSession(handlers.SearchHandler)))

	mux.HandleFunc("OPTIONS /search", m.Cors(handlers.SearchHandler))
	/*
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
