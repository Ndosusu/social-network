// Package models contains the database models and their associated methods.
// This file (post.go) contains the core Post functionality including:
// - Post and PostWithAuthor struct definitions
// - Basic CRUD operations (Create, Read, Update, Delete)
// - Essential post management functions
package models

import (
	"fmt"
	"social-network/pkg/utils"
)

func (db *DB) InsertPost(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			author_id : int,
			message : string,
			image : string (optional),
			privacy_mode : int,
			group_id : int (optional),
		}
	*/
	var imageValue any
	if obj["image"] != nil && obj["image"] != "" {
		imageValue = obj["image"]
	}

	var groupIdValue any
	if obj["group_id"] != nil && obj["group_id"] != 0 {
		groupIdValue = obj["group_id"]
	}

	stmt := "INSERT INTO posts (author_id, message, image, privacy_mode, group_id, date) VALUES (?, ?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["message"], imageValue, obj["privacy_mode"], groupIdValue, utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newPostId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectPostWithAuthorById(map[string]any{"id": newPostId})
}

func (db *DB) SelectPostById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	post := Post{}
	err := result.Scan(&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId)
	if err != nil {
		fmt.Println(err)
		return Response{Post{}}
	}

	return Response{post}
}
func (db *DB) DeletePost(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM posts WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}

// SelectAllPosts moved to post_queries.go

// SelectPostsByUserId moved to post_queries.go

// SelectPostsByGroupId moved to post_queries.go

// SelectPostsByPrivacyMode moved to post_queries.go

func (db *DB) UpdatePost(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
			message : string (optional),
			image : string (optional),
			privacy_mode : int (optional),
			group_id : int (optional),
		}
	*/
	if obj["id"] == nil {
		fmt.Println("Post ID is required for update")
		return Response{0}
	}

	// Build dynamic update query
	setParts := []string{}
	values := []any{}

	if obj["message"] != nil {
		setParts = append(setParts, "message = ?")
		values = append(values, obj["message"])
	}

	if obj["image"] != nil {
		setParts = append(setParts, "image = ?")
		if obj["image"] == "" {
			values = append(values, nil)
		} else {
			values = append(values, obj["image"])
		}
	}

	if obj["privacy_mode"] != nil {
		setParts = append(setParts, "privacy_mode = ?")
		values = append(values, obj["privacy_mode"])
	}

	if obj["group_id"] != nil {
		setParts = append(setParts, "group_id = ?")
		if obj["group_id"] == 0 {
			values = append(values, nil)
		} else {
			values = append(values, obj["group_id"])
		}
	}

	if len(setParts) == 0 {
		fmt.Println("No fields to update")
		return Response{0}
	}

	// Add ID to the end of values for WHERE clause
	values = append(values, obj["id"])

	stmt := "UPDATE posts SET " + setParts[0]
	for i := 1; i < len(setParts); i++ {
		stmt += ", " + setParts[i]
	}
	stmt += " WHERE id = ?;"

	_, err := db.Conn.Exec(stmt, values...)
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	// Return the updated post
	return db.SelectPostById(map[string]any{"id": obj["id"]})
}

// SelectPostsWithoutGroup moved to post_queries.go

// SelectPostWithAuthorById moved to post_with_author.go

// SelectAllPostsWithAuthors moved to post_with_author.go

// SelectPostsByUserIdWithAuthors moved to post_with_author.go

// SelectPostsByGroupIdWithAuthors moved to post_with_author.go
