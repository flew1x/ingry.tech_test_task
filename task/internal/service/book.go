package service

import (
	"errors"

	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"github.com/flew1x/ingry.tech_test_task/internal/repository"

	"github.com/google/uuid"
)

type IBookService interface {
	GetAll() ([]entity.Book, error)
	GetByID(id uuid.UUID) (entity.Book, error)
	Create(title, author string, year uint16) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(id uuid.UUID) error
}

type BookService struct {
	repository repository.IBookRepository
}

func NewBookService(repository repository.IBookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (b *BookService) GetAll() ([]entity.Book, error) {
	books, err := b.repository.GetAll()
	if err != nil {
		if errors.Is(err, repository.ErrAllBooksNotFound) {
			return nil, ErrAllBooksNotFound
		}

		return nil, err
	}

	return books, nil
}

func (b *BookService) GetByID(id uuid.UUID) (entity.Book, error) {
	book, err := b.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			return entity.Book{}, ErrBookNotFound
		}

		return entity.Book{}, err
	}

	return book, nil
}

func (b *BookService) Create(title, author string, year uint16) (entity.Book, error) {
	book := entity.Book{
		ID:     uuid.New(),
		Title:  title,
		Author: author,
		Year:   year,
	}

	book, err := b.repository.Create(book)
	if err != nil {
		if errors.Is(err, repository.ErrBookAlreadyExists) {
			return entity.Book{}, ErrBookAlreadyExists
		}

		return entity.Book{}, err
	}

	return book, nil
}

func (b *BookService) Update(book entity.Book) (entity.Book, error) {
	book, err := b.repository.Update(book)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			return entity.Book{}, ErrBookNotFound
		}

		return entity.Book{}, err
	}

	return book, nil
}

func (b *BookService) Delete(id uuid.UUID) error {
	err := b.repository.Delete(id)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			return ErrBookNotFound
		}

		return err

	}

	return nil
}
