package handlers

import (
	handlers_auth "social-network/pkg/api/handlers/auth"
	handlers_post "social-network/pkg/api/handlers/post"
)

var (
	// Handlers auth
	RegisterHandler = handlers_auth.RegisterHandler
	LoginHandler    = handlers_auth.LoginHandler
	LogoutHandler   = handlers_auth.LogoutHandler

	// Handlers posts
	CreatePostHandler = handlers_post.CreatePostHandler
	//UpdatePostHandler = handlers_post.UpdatePostHandler
	DeletePostHandler = handlers_post.DeletePostHandler
	ServeImageHandler = handlers_post.ServeImageHandler

	GlobalFeedHandler = handlers_post.GlobalFeedHandler
	FollowFeedHandler = handlers_post.FollowFeedHandler
	GroupFeedHandler  = handlers_post.GroupFeedHandler
)
