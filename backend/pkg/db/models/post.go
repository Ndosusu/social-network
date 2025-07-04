package models

import (
	"fmt"
	"social-network/pkg/utils"
)

type Post struct {
	Id          int
	AuthorId    int
	Message     string
	Image       *string
	Date        string
	PrivacyMode int
	GroupId     int
}

func (db *DB) InsertPost(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			author_id : int,
			message : string,
			image : string (optional),
			privacy_mode : int,
			group_id : int,
		}
	*/
	var imageValue interface{}
	if obj["image"] != nil && obj["image"] != "" {
		imageValue = obj["image"]
	} else {
		imageValue = nil
	}

	stmt := "INSERT INTO posts (author_id, message, image, privacy_mode, group_id, date) VALUES (?, ?, ?, ?, ?, ?);"
	result, err := db.Conn.Exec(stmt, obj["author_id"], obj["message"], imageValue, obj["privacy_mode"], obj["group_id"], utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newPostId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return db.SelectPostById(map[string]any{"id": newPostId})
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

func (db *DB) SelectAllPosts(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			limit : int (optional),
			offset : int (optional), // kept for backward compatibility
			last_id : int (optional), // cursor-based pagination
			before : string (optional), // timestamp-based pagination
		}
	*/
	limit := 50 // default limit

	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []interface{}

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE id < ? ORDER BY id DESC LIMIT ?;"
		args = []interface{}{obj["last_id"], limit}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE date < ? ORDER BY date DESC LIMIT ?;"
		args = []interface{}{obj["before"], limit}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts ORDER BY date DESC LIMIT ? OFFSET ?;"
		args = []interface{}{limit, offset}
	}

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]Post{}}
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}

func (db *DB) SelectPostsByUserId(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			user_id : int,
			limit : int (optional),
			offset : int (optional), // kept for backward compatibility
			last_id : int (optional), // cursor-based pagination
			before : string (optional), // timestamp-based pagination
		}
	*/
	limit := 50 // default limit

	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []interface{}

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE author_id = ? AND id < ? ORDER BY id DESC LIMIT ?;"
		args = []interface{}{obj["user_id"], obj["last_id"], limit}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE author_id = ? AND date < ? ORDER BY date DESC LIMIT ?;"
		args = []interface{}{obj["user_id"], obj["before"], limit}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE author_id = ? ORDER BY date DESC LIMIT ? OFFSET ?;"
		args = []interface{}{obj["user_id"], limit, offset}
	}

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]Post{}}
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}

func (db *DB) SelectPostsByGroupId(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			group_id : int,
			limit : int (optional),
			offset : int (optional), // kept for backward compatibility
			last_id : int (optional), // cursor-based pagination
			before : string (optional), // timestamp-based pagination
		}
	*/
	limit := 50 // default limit

	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []interface{}

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE group_id = ? AND id < ? ORDER BY id DESC LIMIT ?;"
		args = []interface{}{obj["group_id"], obj["last_id"], limit}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE group_id = ? AND date < ? ORDER BY date DESC LIMIT ?;"
		args = []interface{}{obj["group_id"], obj["before"], limit}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE group_id = ? ORDER BY date DESC LIMIT ? OFFSET ?;"
		args = []interface{}{obj["group_id"], limit, offset}
	}

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]Post{}}
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}

func (db *DB) SelectPostsByPrivacyMode(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			privacy_mode : int,
			user_id : int (optional, for filtering accessible posts),
			limit : int (optional),
			offset : int (optional), // kept for backward compatibility
			last_id : int (optional), // cursor-based pagination
			before : string (optional), // timestamp-based pagination
		}
	*/
	limit := 50 // default limit

	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []interface{}

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE privacy_mode = ? AND id < ? ORDER BY id DESC LIMIT ?;"
		args = []interface{}{obj["privacy_mode"], obj["last_id"], limit}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE privacy_mode = ? AND date < ? ORDER BY date DESC LIMIT ?;"
		args = []interface{}{obj["privacy_mode"], obj["before"], limit}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		stmt = "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts WHERE privacy_mode = ? ORDER BY date DESC LIMIT ? OFFSET ?;"
		args = []interface{}{obj["privacy_mode"], limit, offset}
	}

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]Post{}}
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}

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
	values := []interface{}{}

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
		values = append(values, obj["group_id"])
	}

	if len(setParts) == 0 {
		fmt.Println("No fields to update")
		return Response{0}
	}

	// Add ID to the end of values for WHERE clause
	values = append(values, obj["id"])

	stmt := "UPDATE posts SET " + fmt.Sprintf("%s", setParts[0])
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
