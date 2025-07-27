package models_like

import (
	"database/sql"
	"fmt"
	"social-network/pkg/db/models"
)

func (db *LikeDB) InsertLike(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			user_id : int
			post_id : int,
			comment_id : int,
		}
	*/
	postID, hasPost := obj["post_id"].(int)
	commentID, hasComment := obj["comment_id"].(int)
	userID := obj["user_id"].(int)

	var stmt string
	var result sql.Result
	var err error

	if hasPost && postID > 0 {
		stmt = "INSERT INTO likes (user_id, post_id) VALUES (?, ?);"
		result, err = db.Conn.Exec(stmt, userID, postID)
	} else if hasComment && commentID > 0 {
		stmt = "INSERT INTO likes (user_id, comment_id) VALUES (?, ?);"
		result, err = db.Conn.Exec(stmt, userID, commentID)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newLikeId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectLikeById(map[string]any{"like_id": newLikeId})
}

func (db *LikeDB) SelectLikeById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			like_id : int,
		}
	*/
	stmt := `SELECT
				id,
				user_id,
	  			COALESCE(post_id, 0),
	   			COALESCE(comment_id, 0)
	    	FROM likes
			WHERE id = ?;`
	result := db.Conn.QueryRow(stmt, obj["like_id"])

	var likeID, userID, postID, commentID int
	err := result.Scan(&likeID, &userID, &postID, &commentID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	like := models.Like{
		Id:     likeID,
		Author: &models.User{Id: userID},
	}
	if postID > 0 {
		like.Post = &models.Post{Id: postID}
	} else if commentID > 0 {
		like.Comment = &models.Comment{Id: commentID}
	} else {
		return nil, fmt.Errorf("like must be associated with either a post or a comment")
	}

	return &models.Response{Result: like}, nil
}

func (db *LikeDB) DeleteLike(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			like_id : int,
		}
	*/
	stmt := "DELETE FROM likes WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["like_id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}
