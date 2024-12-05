package repository

import (
	"github.com/flew1x/ingry.tech_test_task/internal/database"
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

type BookRepository struct {
	db database.MemoryDatabase
}

func NewBookRepository(db database.MemoryDatabase) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) GetAll() ([]entity.Book, error) {
	books := make([]entity.Book, 0)

	for _, value := range r.db.GetAll() {
		book := value.(entity.Book)

		if book.ID == uuid.Nil {
			continue
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) GetByID(id uuid.UUID) (entity.Book, error) {
	book, ok := r.db.Get(id.String())
	if !ok {
		return entity.Book{}, ErrBookNotFound
	}

	return book.(entity.Book), nil
}

func (r *BookRepository) Create(book entity.Book) (entity.Book, error) {
	if book.ID == uuid.Nil {
		return entity.Book{}, ErrBookIDIsNil
	}

	if _, ok := r.db.Get(book.ID.String()); ok {
		return entity.Book{}, ErrBookAlreadyExists
	}

	r.db.Set(book.ID.String(), book)

	return book, nil
}

func (r *BookRepository) Update(book entity.Book) (entity.Book, error) {
	if book.ID == uuid.Nil {
		return entity.Book{}, ErrBookIDIsNil
	}

	if _, ok := r.db.Get(book.ID.String()); !ok {
		return entity.Book{}, ErrBookNotFound
	}

	r.db.Set(book.ID.String(), book)

	return book, nil
}

func (r *BookRepository) Delete(id uuid.UUID) error {
	if _, ok := r.db.Get(id.String()); !ok {
		return ErrBookNotFound
	}

	r.db.Delete(id.String())

	return nil
}
