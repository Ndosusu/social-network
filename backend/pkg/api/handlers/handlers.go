package handlers

import (
	auth "social-network/pkg/api/handlers/auth"
	comment "social-network/pkg/api/handlers/comment"
	like "social-network/pkg/api/handlers/like"
	post "social-network/pkg/api/handlers/post"
)

var (
	// Handlers auth
	RegisterHandler = auth.RegisterHandler
	LoginHandler    = auth.LoginHandler
	LogoutHandler   = auth.LogoutHandler

	// Handlers posts
	CreatePostHandler = post.CreatePostHandler
	//UpdatePostHandler = post.UpdatePostHandler
	DeletePostHandler = post.DeletePostHandler
	ServeImageHandler = post.ServeImageHandler

	GlobalFeedHandler = post.GlobalFeedHandler
	FollowFeedHandler = post.FollowFeedHandler
	GroupFeedHandler  = post.GroupFeedHandler

	// Handlers comments
	CreateCommentHandler = comment.CreateCommentHandler
	//UpdateCommentHandler = comment.UpdateCommentHandler
	DeleteCommentHandler = comment.DeleteCommentHandler

	CommentsHandler = comment.CommentsHandler

	// Handlers likes
	CreateLikeHandler = like.CreateLikeHandler
	DeleteLikeHandler = like.DeleteLikeHandler
)
