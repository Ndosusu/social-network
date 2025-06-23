package models

import (
	"fmt"

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

func (db *DB) CreateUser(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			email : string,
			first_name : string,
			last_name : string,
			password : string
			birth_date : string,
			nickname : string,
			avatar : string
			about : string,
			method : CreateUser
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

	stmt := "INSERT INTO users(uuid, nickname, first_name, last_name, age, gender, email, password) VALUES (?,?,?,?,?,?,?,?);"
	result, err := db.Conn.Exec(stmt, newUUID, obj["nickname"], obj["first_name"], obj["last_name"], obj["age"], obj["gender"], obj["email"], passwordHash)
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
			method : SelectUserById
		}
	*/

	stmt := "SELECT id, uuid, nickname, age, gender, first_name, last_name, email FROM users WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	user := User{}
	err := result.Scan(&user.Id, &user.Uuid, &user.Nickname, &user.First_name, &user.Last_name, &user.Email)
	if err != nil {
		fmt.Println(err)
		return Response{User{}}
	}

	return Response{user}
}
