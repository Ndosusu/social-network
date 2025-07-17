// Package models contains the database models and their associated methods.
// This file (post_with_author.go) contains the complex post queries with author information:
// - PostWithAuthor queries that include user information via JOIN operations
// - Utility functions for scanning and building complex queries
// - Pagination support for queries with author details
// - Functions that return posts enriched with author data for better performance
package models

import "fmt"

// scanPostWithAuthor scans a row into PostWithAuthor struct
func scanPostWithAuthor(rows any, post *Post) error {
	switch r := rows.(type) {
	case interface {
		Scan(dest ...any) error
	}:
		return r.Scan(
			&post.Id, &post.AuthorId, &post.Message, &post.Image, &post.Date, &post.PrivacyMode, &post.GroupId,
			&post.Author.Id, &post.Author.Uuid, &post.Author.Email, &post.Author.First_name, &post.Author.Last_name,
			&post.Author.Birth_date, &post.Author.Avatar, &post.Author.Nickname, &post.Author.About,
			&post.Author.Created_date, &post.Author.Private_mode)
	default:
		return fmt.Errorf("invalid row type")
	}
}

// buildPostWithAuthorQuery builds the base query for PostWithAuthor
func buildPostWithAuthorQuery() string {
	return `SELECT 
		p.id, p.author_id, p.message, p.image, p.date, p.privacy_mode, p.group_id,
		u.id, u.uuid, u.email, u.first_name, u.last_name, u.date_birth, u.avatar, u.nick_name, u.about, u.date_creation, u.private_mode
	FROM posts p 
	JOIN users u ON p.author_id = u.id`
}

// buildPostWithAuthorPaginationQuery builds SQL query with pagination for PostWithAuthor
func buildPostWithAuthorPaginationQuery(baseQuery string, whereClause string, obj map[string]any) (string, []any) {
	limit := 50 // default limit
	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []any

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		if whereClause != "" {
			stmt = baseQuery + " WHERE " + whereClause + " AND p.id < ? ORDER BY p.id DESC LIMIT ?;"
			args = []any{obj["last_id"], limit}
		} else {
			stmt = baseQuery + " WHERE p.id < ? ORDER BY p.id DESC LIMIT ?;"
			args = []any{obj["last_id"], limit}
		}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		if whereClause != "" {
			stmt = baseQuery + " WHERE " + whereClause + " AND p.date < ? ORDER BY p.date DESC LIMIT ?;"
			args = []any{obj["before"], limit}
		} else {
			stmt = baseQuery + " WHERE p.date < ? ORDER BY p.date DESC LIMIT ?;"
			args = []any{obj["before"], limit}
		}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		if whereClause != "" {
			stmt = baseQuery + " WHERE " + whereClause + " ORDER BY p.date DESC LIMIT ? OFFSET ?;"
			args = []any{limit, offset}
		} else {
			stmt = baseQuery + " ORDER BY p.date DESC LIMIT ? OFFSET ?;"
			args = []any{limit, offset}
		}
	}

	return stmt, args
}

// buildPostWithAuthorPaginationQueryWithParam builds SQL query with pagination and a parameter for PostWithAuthor
func buildPostWithAuthorPaginationQueryWithParam(baseQuery string, whereClause string, param any, obj map[string]any) (string, []any) {
	limit := 50 // default limit
	if obj["limit"] != nil {
		limit = int(obj["limit"].(float64))
	}

	var stmt string
	var args []any

	// Priority: cursor-based > timestamp-based > offset-based
	if obj["last_id"] != nil {
		// Cursor-based pagination using last_id
		stmt = baseQuery + " WHERE " + whereClause + " AND p.id < ? ORDER BY p.id DESC LIMIT ?;"
		args = []any{param, obj["last_id"], limit}
	} else if obj["before"] != nil {
		// Timestamp-based pagination
		stmt = baseQuery + " WHERE " + whereClause + " AND p.date < ? ORDER BY p.date DESC LIMIT ?;"
		args = []any{param, obj["before"], limit}
	} else {
		// Fallback to offset-based pagination for backward compatibility
		offset := 0
		if obj["offset"] != nil {
			offset = int(obj["offset"].(float64))
		}
		stmt = baseQuery + " WHERE " + whereClause + " ORDER BY p.date DESC LIMIT ? OFFSET ?;"
		args = []any{param, limit, offset}
	}

	return stmt, args
}

func (db *DB) SelectPostWithAuthorById(obj map[string]any) Response {
	stmt := buildPostWithAuthorQuery() + " WHERE p.id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	post := PostWithAuthor{}
	err := scanPostWithAuthor(result, &post)
	if err != nil {
		fmt.Println(err)
		return Response{PostWithAuthor{}}
	}

	return Response{post}
}

func (db *DB) SelectAllPostsWithAuthors(obj map[string]any) Response {
	baseQuery := buildPostWithAuthorQuery()
	stmt, args := buildPostWithAuthorPaginationQuery(baseQuery, "", obj)

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]PostWithAuthor{}}
	}
	defer rows.Close()

	var posts []PostWithAuthor
	for rows.Next() {
		post := PostWithAuthor{}
		err := scanPostWithAuthor(rows, &post)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}

func (db *DB) SelectPostsByUserIdWithAuthors(obj map[string]any) Response {
	baseQuery := buildPostWithAuthorQuery()
	stmt, args := buildPostWithAuthorPaginationQueryWithParam(baseQuery, "p.author_id = ?", obj["user_id"], obj)

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]PostWithAuthor{}}
	}
	defer rows.Close()

	var posts []PostWithAuthor
	for rows.Next() {
		post := PostWithAuthor{}
		err := scanPostWithAuthor(rows, &post)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}

func (db *DB) SelectPostsByGroupIdWithAuthors(obj map[string]any) Response {
	baseQuery := buildPostWithAuthorQuery()
	stmt, args := buildPostWithAuthorPaginationQueryWithParam(baseQuery, "p.group_id = ?", obj["group_id"], obj)

	rows, err := db.Conn.Query(stmt, args...)
	if err != nil {
		fmt.Println(err)
		return Response{[]PostWithAuthor{}}
	}
	defer rows.Close()

	var posts []PostWithAuthor
	for rows.Next() {
		post := PostWithAuthor{}
		err := scanPostWithAuthor(rows, &post)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return Response{posts}
}
