package repository

import "github.com/flew1x/ingry.tech_test_task/internal/database"

type Repository struct {
	Book IBookRepository
}

func NewRepository(db database.MemoryDatabase) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
	}
}
