// Package models contains the database models and their associated methods.
// This file (post_queries.go) contains the post query functions including:
// - Pagination utility functions for consistent query building
// - Post selection methods with various filters (user, group, privacy mode)
// - Optimized query building to reduce code duplication
// - Support for cursor-based, timestamp-based, and offset-based pagination
package models

import "fmt"

// buildPaginationQuery builds SQL query with pagination support
func buildPaginationQuery(baseQuery string, whereClause string, obj map[string]any) (string, []interface{}) {
	limit := 50 // default limit
	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []interface{}

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		if whereClause != "" {
			stmt = baseQuery + " WHERE " + whereClause + " AND id < ? ORDER BY id DESC LIMIT ?;"
			args = []interface{}{obj["last_id"], limit}
		} else {
			stmt = baseQuery + " WHERE id < ? ORDER BY id DESC LIMIT ?;"
			args = []interface{}{obj["last_id"], limit}
		}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		if whereClause != "" {
			stmt = baseQuery + " WHERE " + whereClause + " AND date < ? ORDER BY date DESC LIMIT ?;"
			args = []interface{}{obj["before"], limit}
		} else {
			stmt = baseQuery + " WHERE date < ? ORDER BY date DESC LIMIT ?;"
			args = []interface{}{obj["before"], limit}
		}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		if whereClause != "" {
			stmt = baseQuery + " WHERE " + whereClause + " ORDER BY date DESC LIMIT ? OFFSET ?;"
			args = []interface{}{limit, offset}
		} else {
			stmt = baseQuery + " ORDER BY date DESC LIMIT ? OFFSET ?;"
			args = []interface{}{limit, offset}
		}
	}

	return stmt, args
}

// buildPaginationQueryWithParam builds SQL query with pagination support and a parameter
func buildPaginationQueryWithParam(baseQuery string, whereClause string, param interface{}, obj map[string]any) (string, []interface{}) {
	limit := 50 // default limit
	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []interface{}

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		stmt = baseQuery + " WHERE " + whereClause + " AND id < ? ORDER BY id DESC LIMIT ?;"
		args = []interface{}{param, obj["last_id"], limit}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		stmt = baseQuery + " WHERE " + whereClause + " AND date < ? ORDER BY date DESC LIMIT ?;"
		args = []interface{}{param, obj["before"], limit}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		stmt = baseQuery + " WHERE " + whereClause + " ORDER BY date DESC LIMIT ? OFFSET ?;"
		args = []interface{}{param, limit, offset}
	}

	return stmt, args
}

func (db *DB) SelectAllPosts(obj map[string]any) Response {
	baseQuery := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts"
	stmt, args := buildPaginationQuery(baseQuery, "", obj)

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
	baseQuery := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts"
	stmt, args := buildPaginationQueryWithParam(baseQuery, "author_id = ?", obj["user_id"], obj)

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
	baseQuery := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts"
	stmt, args := buildPaginationQueryWithParam(baseQuery, "group_id = ?", obj["group_id"], obj)

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
	baseQuery := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts"
	stmt, args := buildPaginationQueryWithParam(baseQuery, "privacy_mode = ?", obj["privacy_mode"], obj)

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

func (db *DB) SelectPostsWithoutGroup(obj map[string]any) Response {
	baseQuery := "SELECT id, author_id, message, image, date, privacy_mode, group_id FROM posts"
	stmt, args := buildPaginationQuery(baseQuery, "group_id IS NULL", obj)

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
