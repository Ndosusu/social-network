package models

import "fmt"

// Queries about follow relations
func (db *DB) InsertFollowRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			user_from : int,
			user_to : int,
		}
	*/
	stmt := "INSERT INTO follow_rel (user_to, user_from) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["user_to"], obj["user_from"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
func (db *DB) DeleteFollowRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			user_from : int,
			user_to : int,
		}
	*/
	stmt := "DELETE FROM follow_rel WHERE user_to = ? AND user_from = ?;"
	_, err := db.Conn.Exec(stmt, obj["user_to"], obj["user_from"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}

// Queries about post privacy relations
func (db *DB) InsertPrivacyPostRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			post_id : int,
			follower_id : int,
		}
	*/
	stmt := "INSERT INTO privacy_post_rel (post_id, follower_id) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["post_id"], obj["follower_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
func (db *DB) DeletePrivacyPostRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			post_id : int,
			follower_id : int,
		}
	*/
	stmt := "DELETE FROM privacy_post_rel WHERE post_id = ? AND follower_id = ?;"
	_, err := db.Conn.Exec(stmt, obj["post_id"], obj["follower_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}

// Queries about group membership relations
func (db *DB) InsertGroupMemberRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			group_id : int,
			member_id : int,
		}
	*/
	stmt := "INSERT INTO group_member_rel (group_id, member_id) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["group_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
func (db *DB) DeleteGroupMemberRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			group_id : int,
			member_id : int,
		}
	*/
	stmt := "DELETE FROM group_member_rel WHERE group_id = ? AND member_id = ?;"
	_, err := db.Conn.Exec(stmt, obj["group_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}

// Queries about event membership relations
func (db *DB) InsertEventMemberRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			event_id : int,
			member_id : int,
		}
	*/
	stmt := "INSERT INTO event_member_rel (event_id, member_id) VALUES (?, ?);"
	_, err := db.Conn.Exec(stmt, obj["event_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
func (db *DB) DeleteEventMemberRel(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			event_id : int,
			member_id : int,
		}
	*/
	stmt := "DELETE FROM event_member_rel WHERE event_id = ? AND member_id = ?;"
	_, err := db.Conn.Exec(stmt, obj["event_id"], obj["member_id"])
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	return Response{1}
}
