package handler

import (
	"github.com/petaki/inertia-go"
	"gorm.io/gorm"
)

type Handler struct {
	Inertia *inertia.Inertia
	DB      *gorm.DB
}

func New(inertia *inertia.Inertia, db *gorm.DB) *Handler {
	return &Handler{
		Inertia: inertia,
		DB:      db,
	}
}
