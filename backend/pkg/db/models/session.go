package models

import (
	"fmt"
	"social-network/pkg/utils"

	"github.com/gofrs/uuid"
)

func (db *DB) InsertSession(userId int) Response {
	newUUID, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}

	stmt := "INSERT INTO sessions (uuid, user_id, date_creation) VALUES (?,?,?);"
	_, err = db.Conn.Exec(stmt, newUUID, userId, utils.GetCurrentTime())
	if err != nil {
		fmt.Println(err)
		return Response{0}
	}
	return Response{1}
}

func (db *DB) GetSessionByUuid(obj map[string]any) Response {
	/*
		expected input (as json object) :
		{
			uuid : string,
	}*/
	stmt := "SELECT id, uuid, user_id, date_creation, date_expires FROM sessions WHERE uuid = ? AND (date_expires >= ? OR date_expires IS NULL);"
	result := db.Conn.QueryRow(stmt, obj["id"])

	session := Session{}
	err := result.Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt, &session.ExpiresAt)
	if err != nil {
		fmt.Println(err)
		return Response{Session{}}
	}

	return Response{session}
}
