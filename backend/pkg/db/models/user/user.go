package models_user

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (db *UserDB) InsertUser(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	stmt := "INSERT INTO users (email, password, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode) VALUES (?,?,?,?,?,?,?,?,?,?);"
	result, err := db.Conn.Exec(stmt, obj["email"], passwordHash, obj["first_name"], obj["last_name"], obj["date_birth"], obj["avatar"], obj["nickname"], obj["about"], utils.GetCurrentTime(), false)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	newUserId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response, err := db.InsertSession(int(newUserId))
	if err != nil {
		fmt.Println("Failed to create session for new user")
		return nil, err
	}

	return response, nil
}

func (db *UserDB) SelectUserById(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			id : int,
		}
	*/

	stmt := "SELECT id, email, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode FROM users WHERE id = ?;"
	result := db.Conn.QueryRow(stmt, obj["id"])

	user := models.User{}
	err := result.Scan(&user.Id, &user.Email, &user.First_name, &user.Last_name, &user.Birth_date, &user.Avatar, &user.Nickname, &user.About, &user.Created_date, &user.Private_mode)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &models.Response{Result: user}, nil
}

func (db *UserDB) SelectUserByUuid(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			uuid : string,
		}
	*/

	stmt := "SELECT id, email, first_name, last_name, date_birth, avatar, nick_name, about, date_creation, private_mode FROM users WHERE uuid = ?;"
	result := db.Conn.QueryRow(stmt, obj["uuid"])

	user := models.User{}
	err := result.Scan(&user.Id, &user.Email, &user.First_name, &user.Last_name, &user.Birth_date, &user.Avatar, &user.Nickname, &user.About, &user.Created_date, &user.Private_mode)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: user}, nil
}

func (db *UserDB) Authenticate(obj map[string]any) (*models.Response, error) {
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
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(password, []byte(obj["password"].(string)))
	if err != nil {
		return nil, err
	}

	response, err := db.InsertSession(int(id))
	if err != nil {
		fmt.Println("Failed to create session for new user")
		return nil, err
	}

	return response, nil
}
