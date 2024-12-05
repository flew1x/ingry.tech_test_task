package repository

import (
	"github.com/flew1x/ingry.tech_test_task/internal/database"
	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"github.com/google/uuid"
)

type InMemoryBookRepository struct {
	db database.MemoryDatabase
}

func NewInMemoryBookRepository(db database.MemoryDatabase) *InMemoryBookRepository {
	return &InMemoryBookRepository{
		db: db,
	}
}

func (r *InMemoryBookRepository) GetAll() ([]entity.Book, error) {
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

func (r *InMemoryBookRepository) GetByID(id uuid.UUID) (entity.Book, error) {
	book, ok := r.db.Get(id.String())
	if !ok {
		return entity.Book{}, ErrBookNotFound
	}

	return book.(entity.Book), nil
}

func (r *InMemoryBookRepository) Create(book entity.Book) (entity.Book, error) {
	if book.ID == uuid.Nil {
		return entity.Book{}, ErrBookIDIsNil
	}

	if _, ok := r.db.Get(book.ID.String()); ok {
		return entity.Book{}, ErrBookAlreadyExists
	}

	r.db.Set(book.ID.String(), book)

	return book, nil
}

func (r *InMemoryBookRepository) Update(book entity.Book) (entity.Book, error) {
	if book.ID == uuid.Nil {
		return entity.Book{}, ErrBookIDIsNil
	}

	if _, ok := r.db.Get(book.ID.String()); !ok {
		return entity.Book{}, ErrBookNotFound
	}

	r.db.Set(book.ID.String(), book)

	return book, nil
}

func (r *InMemoryBookRepository) Delete(id uuid.UUID) error {
	if _, ok := r.db.Get(id.String()); !ok {
		return ErrBookNotFound
	}

	r.db.Delete(id.String())

	return nil
}
