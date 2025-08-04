package handlers

import (
	auth "social-network/pkg/api/handlers/auth"
	comment "social-network/pkg/api/handlers/comment"
	group "social-network/pkg/api/handlers/group"
	like "social-network/pkg/api/handlers/like"
	post "social-network/pkg/api/handlers/post"
	search "social-network/pkg/api/handlers/search"
	user "social-network/pkg/api/handlers/user"
)

var (
	// Handlers auth
	RegisterHandler = auth.RegisterHandler
	LoginHandler    = auth.LoginHandler
	LogoutHandler   = auth.LogoutHandler

	// Handlers posts
	CreatePostHandler = post.CreatePostHandler
	UpdatePostHandler = post.UpdatePostHandler
	DeletePostHandler = post.DeletePostHandler
	ServeImageHandler = post.ServeImageHandler // Useful ?
	GlobalFeedHandler = post.GlobalFeedHandler
	FollowFeedHandler = post.FollowFeedHandler
	GroupFeedHandler  = post.GroupFeedHandler

	// Handlers comments
	CreateCommentHandler = comment.CreateCommentHandler
	UpdateCommentHandler = comment.UpdateCommentHandler
	DeleteCommentHandler = comment.DeleteCommentHandler
	CommentsHandler      = comment.CommentsHandler

	// Handlers likes
	CreateLikeHandler = like.CreateLikeHandler
	DeleteLikeHandler = like.DeleteLikeHandler

	// Handlers groups
	CreateGroupHandler = group.CreateGroupHandler
	//UpdateGroupHandler = group.UpdateGroupHandler
	DeleteGroupHandler = group.DeleteGroupHandler
	ListGroupsHandler  = group.ListGroupsHandler

	// Handlers search
	SearchHandler = search.SearchHandler

	// Handlers profile
	UserProfileHandler = user.UserProfileHandler
	//UpdateProfileHandler = user.UpdateProfileHandler
	FollowCreateHandler = user.FollowCreateHandler
	FollowDeleteHandler = user.FollowDeleteHandler
)
