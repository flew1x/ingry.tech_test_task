package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;not null"`
	Title  string    `json:"title" gorm:"not null"`
	Author string    `json:"author" gorm:"not null"`
	Year   uint16    `json:"year" gorm:"not null"`
}
