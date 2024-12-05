package v1

import (
	"errors"
	"net/http"

	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"github.com/flew1x/ingry.tech_test_task/internal/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type getBooksResponse struct {
	Books []entity.Book `json:"books"`
}

func (h *Handler) getBooks(c echo.Context) error {
	books, err := h.service.Book.GetAll()
	if err != nil {
		if errors.Is(err, service.ErrAllBooksNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return ErrInternalServerError
	}

	return c.JSON(http.StatusOK, getBooksResponse{Books: books})
}

type getBookByIDResponse struct {
	Book entity.Book `json:"book"`
}

func (h *Handler) getBookByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return ErrInvalidID
	}

	book, err := h.service.Book.GetByID(id)
	if err != nil {
		if errors.Is(err, service.ErrBookNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return ErrInternalServerError
	}

	return c.JSON(http.StatusOK, getBookByIDResponse{
		Book: book,
	})
}

type createBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   uint16 `json:"year"`
}

func (h *Handler) createBook(c echo.Context) error {
	req := createBookRequest{}

	if err := c.Bind(&req); err != nil {
		return ErrInvalidRequest
	}

	switch {
	case req.Title == "":
		return ErrTitleShouldNotBeEmpty
	case req.Author == "":
		return ErrAuthorShouldNotBeEmpty
	case req.Year == 0:
		return ErrYearShouldNotBeEmpty
	}

	book, err := h.service.Book.Create(req.Title, req.Author, req.Year)
	if err != nil {
		if errors.Is(err, service.ErrBookAlreadyExists) {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		return ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, book)
}

type updateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   uint16 `json:"year"`
}

func (h *Handler) updateBook(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return ErrInvalidID
	}

	req := updateBookRequest{}

	if err := c.Bind(&req); err != nil {
		return ErrInvalidRequest
	}

	switch {
	case req.Title == "":
		return ErrTitleShouldNotBeEmpty
	case req.Author == "":
		return ErrAuthorShouldNotBeEmpty
	case req.Year == 0:
		return ErrYearShouldNotBeEmpty
	}

	book, err := h.service.Book.Update(entity.Book{
		ID:     id,
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	})
	if err != nil {
		if errors.Is(err, service.ErrBookNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return ErrInternalServerError
	}

	return c.JSON(http.StatusOK, book)
}

func (h *Handler) deleteBook(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return ErrInvalidID
	}

	err = h.service.Book.Delete(id)
	if err != nil {
		if errors.Is(err, service.ErrBookNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
