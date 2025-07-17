package models

import (
	"fmt"
	"social-network/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) InsertUser(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			email : string,
			first_name : string,
			last_name : string,
			password : string
			date_birth : string,
			nickname : string,
			avatar : string
			about : string,
		}
	*/

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(obj["password"].(string)), 12)
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	stmt := "INSERT INTO users (email, password, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode) VALUES (?,?,?,?,?,?,?,?,?,?,?);"
	result, err := db.Conn.Exec(stmt, obj["email"], passwordHash, obj["first_name"], obj["last_name"], obj["date_birth"], obj["avatar"], obj["nickname"], obj["about"], utils.GetCurrentTime(), false)
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	response := db.InsertSession(int(newUserId))
	if response.Result == 0 {
		fmt.Println("Failed to create session for new user")
		return Response{0}
	}

	return db.SelectUserById(map[string]any{"id": newUserId})
}

func (db *DB) SelectUserById(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/

	stmt := "SELECT id, email, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode FROM users WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	user := User{}
	err := result.Scan(&user.Id, &user.Email, &user.First_name, &user.Last_name, &user.Birth_date, &user.Avatar, &user.Nickname, &user.About, &user.Created_date, &user.Private_mode)
	if err != nil {
		fmt.Println(err)
		return Response{User{}}
	}

	return Response{user}
}

func (db *DB) SelectUserByUuid(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			uuid : string,
		}
	*/

	stmt := "SELECT id, email, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode FROM users WHERE uuid = ?;"
	result := db.Conn.QueryRow(stmt, obj["uuid"])

	user := User{}
	err := result.Scan(&user.Id, &user.Email, &user.First_name, &user.Last_name, &user.Birth_date, &user.Avatar, &user.Nickname, &user.About, &user.Created_date, &user.Private_mode)
	if err != nil {
		fmt.Println(err)
		return Response{User{}}
	}

	return Response{user}
}

func (db *DB) Authenticate(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			mail : string,
			password : string,
		}
	*/
	var id int
	var password []byte
	stmt := "SELECT id, password FROM users WHERE email = ?;"
	result := db.Conn.QueryRow(stmt, obj["mail"])
	err := result.Scan(&id, &password)
	if err != nil {
		return Response{User{}}
	}

	err = bcrypt.CompareHashAndPassword(password, []byte(obj["password"].(string)))
	if err != nil {
		return Response{User{}}
	}

	response := db.InsertSession(int(id))
	if response.Result == 0 {
		fmt.Println("Failed to create session for new user")
		return Response{0}
	}

	return db.SelectUserById(map[string]any{"id": id})
}
