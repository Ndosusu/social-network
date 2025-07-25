package models_post

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"
)

func (db *PostDB) InsertComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			author_id : int,
			post_id : int,
			message : string,
			image : string,
		}
	*/
	stmt := "INSERT INTO comments (author_id, post_id, message, image, group_id, date) VALUES (?, ?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["post_id"], obj["message"], obj["image"], obj["group_id"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newCommentId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db.SelectCommentById(map[string]any{"id": newCommentId})
}

func (db *PostDB) SelectCommentById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "SELECT id, author_id, post_id, message, image, date, group_id FROM comments WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	comment := models.Comment{}
	err := result.Scan(&comment.Id, &comment.AuthorId, &comment.PostId, &comment.Message, &comment.Image, &comment.Date)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: comment}, nil
}

func (db *PostDB) DeleteComment(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/
	stmt := "DELETE FROM comments WHERE id = ?;"
	_, err := db.Conn.Exec(stmt, obj["id"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: "Ok"}, nil
}

func (db *PostDB) SelectCommentsByPostId(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			post_id : int,
			last_id : int,
		}
	*/
	stmt := `SELECT
				c.id,
				c.author_id,
				u.nick_name,
				u.avatar, 
				c.message, 
				COALESCE(c.image, ''), 
				c.date, 
				COUNT(DISTINCT l.id)
			FROM comments c
			JOIN users u ON c.author_id = u.id
			LEFT JOIN likes l ON l.comment_id = c.id
			WHERE 
				post_id = ? 
				AND c.id < ?
			GROUP BY c.id,
			ORDER BY c.id DESC
			LIMIT ?;`
	rows, err := db.Conn.Query(stmt, obj["post_id"], obj["last_id"], obj["limit"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var result []models.CommentFeed
	for rows.Next() {
		var c models.Comment
		var commentImage string
		var author models.User
		var likeCount int

		err := rows.Scan(
			&c.Id,
			&c.AuthorId,
			&author.Nickname,
			&author.Avatar,
			&c.Message,
			&commentImage,
			&c.Date,
			&likeCount,
		)
		if err != nil {
			return nil, err
		}
		if commentImage != "" {
			c.Image = &commentImage
		}
		c.Author = &author

		result = append(result, models.CommentFeed{
			Comment:   &c,
			LikeCount: likeCount,
		})
	}

	return &models.Response{Result: result}, nil
}
