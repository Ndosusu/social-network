package models_user

import (
	"fmt"
	"social-network/pkg/db/models"
	"social-network/pkg/utils"

	"github.com/gofrs/uuid"
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
		fmt.Println(err)
		return nil, err
	}
	return db.GetSessionByUuid(map[string]any{"id": newUUID})
}

func (db *UserDB) GetSessionByUuid(obj map[string]any) (*models.Response, error) {
	/*
		expected input (as json object) :
		{
			uuid : string,
	}*/
	stmt := "SELECT id, uuid, user_id, date_creation, COALESCE(date_expiration, '') FROM sessions WHERE uuid = ? AND (date_expiration >= ? OR date_expiration IS NULL);"
	result := db.Conn.QueryRow(stmt, obj["id"], utils.GetCurrentTime())

	session := models.Session{}
	err := result.Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt, &session.ExpiresAt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.Response{Result: session}, nil
}
