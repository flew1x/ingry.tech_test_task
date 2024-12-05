package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Book IBookRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Book: NewPostgresBookRepository(db),
	}
}
