package models_user

import "social-network/pkg/db/models"

type UserDB struct {
	*models.DB
}

func New(db *models.DB) *UserDB {
	return &UserDB{DB: db}
}
