package repository

import (
	"github.com/flew1x/ingry.tech_test_task/internal/entity"

	"github.com/google/uuid"
)

type IBookRepository interface {
	GetAll() ([]entity.Book, error)
	GetByID(id uuid.UUID) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(id uuid.UUID) error
}
