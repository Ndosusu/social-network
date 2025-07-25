package models_notif

import "social-network/pkg/db/models"

type NotifDB struct {
	*models.DB
}

func New(db *models.DB) *NotifDB {
	return &NotifDB{DB: db}
}
