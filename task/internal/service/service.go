package service

import "github.com/flew1x/ingry.tech_test_task/internal/repository"

type Service struct {
	Book IBookService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repository.Book),
	}
}
