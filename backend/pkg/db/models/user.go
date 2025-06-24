package models

import (
	"fmt"
	"social-network/pkg/utils"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int
	Uuid         string
	Email        string
	Nickname     string
	First_name   string
	Last_name    string
	Birth_date   string
	About        string
	Avatar       string
	Created_date string
	Private_mode bool
}

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

	newUUID, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	stmt := "INSERT INTO users(uuid, email, password, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode) VALUES (?,?,?,?,?,?,?,?,?,?,?);"
	result, err := db.Conn.Exec(stmt, newUUID, obj["email"], passwordHash, obj["first_name"], obj["last_name"], obj["date_birth"], obj["avatar"], obj["nickname"], obj["about"], utils.GetCurrentTime(), false)
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
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

	stmt := "SELECT id, uuid, email, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	user := User{}
	err := result.Scan(&user.Id, &user.Uuid, &user.Email, &user.First_name, &user.Last_name, &user.Birth_date, &user.Avatar, &user.Nickname, &user.About, &user.Created_date, &user.Private_mode)
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

	return db.SelectUserById(map[string]any{"id": id})
}
