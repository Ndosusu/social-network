// Package models contains the database models and their associated methods.
// This file (post.go) contains the core Post functionality including:
// - Post and PostWithAuthor struct definitions
// - Basic CRUD operations (Create, Read, Update, Delete)
// - Essential post management functions
package models_post

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *PostDB) InsertPost(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	newPostId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectPostById(map[string]any{
		"post_id":   newPostId,
		"author_id": obj["author_id"],
	})
}

func (db *PostDB) SelectPostById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			post_id : int,
			author_id : int,
		}
	*/
	stmt := `SELECT 
                p.id,
				p.author_id,
				p.message,
				COALESCE(p.image, ''),
				p.date,
				p.privacy_mode, 
				COALESCE(p.group_id, 0),
                COALESCE(u.avatar, ''),
				COALESCE(u.nick_name, ''),
				u.first_name,
				u.last_name,
				COALESCE(g.title, ''),
				CASE
					WHEN p.author_id = ? THEN 1
					ELSE 0
				END as is_client
	        FROM posts p
            JOIN users u ON p.author_id = u.id
            LEFT JOIN groups g ON p.group_id = g.id
            WHERE p.id = ?;`
	result := db.Conn.QueryRow(stmt, obj["author_id"], obj["post_id"])

	var postId, authorId, privacyMode, groupId int
	var message, date, nickname, groupTitle, avatar, postImage, firstName, lastName string
	var isClient bool

	err := result.Scan(&postId, &authorId, &message, &postImage, &date, &privacyMode, &groupId, &avatar, &nickname, &firstName, &lastName, &groupTitle, &isClient)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	post := models.Post{
		Id:          postId,
		Message:     message,
		Date:        date,
		PrivacyMode: privacyMode,
		Author: &models.User{
			Id:        authorId,
			FirstName: firstName,
			LastName:  lastName,
			IsClient:  isClient,
		},
	}
	if nickname != "" {
		post.Author.Nickname = nickname
	}

	if postImage != "" {
		post.Image = &postImage
	}
	if avatar != "" {
		post.Author.Avatar = &avatar
	}
	if groupId > 0 {
		post.Group = &models.Group{
			Id:    groupId,
			Title: groupTitle,
		}
	}

	return &models.Response{Result: post}, nil
}

func (db *PostDB) DeletePost(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			post_id : int,
		}
	*/
	stmt := "DELETE FROM posts WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["post_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}

/*func (db *PostDB) UpdatePost(obj map[string]any) (*models.Response, error) {

		expected input (as json object) :
		{
			id : int,
			message : string (optional),
			image : string (optional),
			privacy_mode : int (optional),
			group_id : int (optional),
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
		err := errors.New("no fields to update")
		fmt.Println("No fields to update")
		return nil, err
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
		return nil, err
	}

	// Return the updated post
	return db.SelectPostById(map[string]any{"id": obj["id"]})
}
*/
