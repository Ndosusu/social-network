package models_user

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (db *UserDB) InsertSession(userId int) (*models.Response, error) {
	newUUID, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	stmt := "INSERT INTO sessions (uuid, user_id, date_creation) VALUES (?,?,?);"
	_, err = db.Conn.Exec(stmt, newUUID, userId, utils.GetCurrentTime())
	if err != nil {
		return nil, err
	}
	return db.GetSessionByUuid(map[string]any{"session_uuid": newUUID})
}

func (db *UserDB) GetSessionByUuid(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			session_uuid : string,
		}
	*/
	stmt := `SELECT
				id, 
				uuid, 
				user_id, 
				date_creation, 
				COALESCE(date_expiration, '') 
			FROM sessions 
			WHERE 
				uuid = ? 
				AND (date_expiration >= ? OR date_expiration IS NULL);`
	result := db.Conn.QueryRow(stmt, obj["session_uuid"], utils.GetCurrentTime())

	session := models.Session{
		User: &models.User{Id: utils.NOT_SCANNED},
	}
	err := result.Scan(&session.Id, &session.Uuid, &session.User.Id, &session.CreatedAt, &session.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return &models.Response{Result: session}, nil
}

func (db *UserDB) CloseSession(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			session_uuid : string,
		}
	*/
	stmt := "UPDATE sessions SET date_expiration = ? WHERE uuid = ?;"
	result, err := db.Conn.Exec(stmt, utils.GetCurrentTime(), obj["session_uuid"])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &models.Response{Result: result}, nil
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
		fmt.Println("Failed to create session for the user")
		return nil, err
	}

	return response, nil
}
