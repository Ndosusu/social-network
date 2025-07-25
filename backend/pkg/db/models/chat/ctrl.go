package models_chat

import "social-network/pkg/db/models"

type ChatDB struct {
	*models.DB
}

func New(db *models.DB) *ChatDB {
	return &ChatDB{DB: db}
}
