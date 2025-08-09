package models_comment

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
	"strings"
)

func (db *CommentDB) InsertComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			author_id : int,
			post_id : int,
			message : string,
			image : string,
		}
	*/
	stmt := "INSERT INTO comments (author_id, post_id, message, image, date) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["post_id"], obj["message"], obj["image"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newCommentId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectCommentById(map[string]any{
		"comment_id": newCommentId,
		"author_id":  obj["author_id"],
	})
}

func (db *CommentDB) SelectCommentById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			comment_id : int,
			author_id : int,
		}
	*/
	stmt := `SELECT
				c.id, 
				c.author_id, 
				c.message, 
				COALESCE(c.image, ''), 
				c.date,
				COALESCE(u.avatar, ''),
				COALESCE(u.nick_name, ''),
				u.first_name,
				u.last_name,
				CASE
					WHEN c.author_id = ? THEN 1
					ELSE 0
				END AS is_client
			FROM comments c 
			JOIN users u ON c.author_id = u.id
			WHERE c.id = ?;`
	result := db.Conn.QueryRow(stmt, obj["author_id"], obj["comment_id"])

	var commentId, authorId int
	var message, date, nickname, avatar, commentImage, firstName, lastName string
	var isClient bool

	err := result.Scan(&commentId, &authorId, &message, &commentImage, &date, &avatar, &nickname, &firstName, &lastName, &isClient)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	comment := models.Comment{
		Id:      commentId,
		Message: message,
		Date:    date,
		Author: &models.User{
			Id:        authorId,
			FirstName: firstName,
			LastName:  lastName,
			IsClient:  isClient,
		},
	}
	if nickname != "" {
		comment.Author.Nickname = nickname
	}

	if commentImage != "" {
		comment.Image = &commentImage
	}
	if avatar != "" {
		comment.Author.Avatar = &avatar
	}

	return &models.Response{Result: comment}, nil

}

func (db *CommentDB) DeleteComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			comment_id : int,
		}
	*/
	stmt := "DELETE FROM comments WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["comment_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}

func (db *CommentDB) UpdateComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			comment_id : int,
			author_id : int,
			message : string,
			image : string,
		}
	*/

	// Build dynamic SQL statement
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

	if len(setParts) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	values = append(values, obj["comment_id"])

	// Join the set parts to form the final query
	stmt := "UPDATE comments SET " + strings.Join(setParts, ", ") + " WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, values...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db.SelectCommentById(map[string]any{
		"comment_id": obj["comment_id"],
		"author_id":  obj["author_id"],
	})
}
