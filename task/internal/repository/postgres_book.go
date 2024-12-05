package repository

import (
	"errors"

	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresBookRepository struct {
	db *gorm.DB
}

func NewPostgresBookRepository(db *gorm.DB) *PostgresBookRepository {
	return &PostgresBookRepository{
		db: db,
	}
}

func (r *PostgresBookRepository) GetAll() ([]entity.Book, error) {
	var books []entity.Book

	if err := r.db.Find(&books).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAllBooksNotFound
		}

		return nil, err
	}

	if len(books) == 0 {
		return nil, ErrAllBooksNotFound
	}

	return books, nil
}

func (r *PostgresBookRepository) GetByID(id uuid.UUID) (entity.Book, error) {
	var book entity.Book

	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, ErrBookNotFound
		}

		return entity.Book{}, err
	}

	return book, nil
}

func (r *PostgresBookRepository) Create(book entity.Book) (entity.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return entity.Book{}, ErrBookAlreadyExists
		}

		return entity.Book{}, err
	}

	return book, nil
}

func (r *PostgresBookRepository) Update(book entity.Book) (entity.Book, error) {
	result := r.db.Model(&book).Updates(book)
	if err := result.Error; err != nil {
		return entity.Book{}, err
	}

	if result.RowsAffected == 0 {
		return entity.Book{}, ErrBookNotFound
	}

	return book, nil
}

func (r *PostgresBookRepository) Delete(id uuid.UUID) error {
	result := r.db.Where("id = ?", id).Delete(&entity.Book{})

	if result.RowsAffected == 0 {
		return ErrBookNotFound
	}

	return result.Error
}
