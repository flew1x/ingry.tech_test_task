package service

import "errors"

var (
	ErrAllBooksNotFound  = errors.New("all books not found")
	ErrBookNotFound      = errors.New("book not found")
	ErrBookAlreadyExists = errors.New("book already exists")
)
