package models_event

import "social-network/pkg/db/models"

type EventDB struct {
	*models.DB
}

func New(db *models.DB) *EventDB {
	return &EventDB{DB: db}
}
