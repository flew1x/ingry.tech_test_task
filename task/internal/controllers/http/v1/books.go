package v1

import (
	"errors"
	"net/http"

	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"github.com/flew1x/ingry.tech_test_task/internal/service"

	_ "github.com/flew1x/ingry.tech_test_task/docs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type getBooksResponse struct {
	Books []entity.Book `json:"books"`
}

// GetBooks godoc
// @Summary Retrieves a list of all books
// @Description Fetches the complete list of books and returns it in JSON format.
// @Tags books
// @Produce json
// @Success 200 {array} entity.Book "Successfully retrieved list of books"
// @Failure 404 {object} echo.HTTPError "Books not found"
// @Failure 500 {object} echo.HTTPError "Internal server error" {"error": "Internal server error"}
// @Router /api/v1/books [get]
func (h *Handler) GetBooks(c echo.Context) error {
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

// GetBookByID godoc
// @Summary Retrieves a book by its ID
// @Description Fetches a book based on the provided ID and returns it in JSON format.
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} getBookByIDResponse "Successfully retrieved the book"
// @Failure 400 {object} echo.HTTPError "Invalid ID"
// @Failure 404 {object} echo.HTTPError "Book not found"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /api/v1/books/{id} [get]
func (h *Handler) GetBookByID(c echo.Context) error {
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

// CreateBook godoc
// @Summary Creates a new book
// @Description Creates a new book with the provided title, author, and year.
// @Tags books
// @Accept json
// @Produce json
// @Param book body createBookRequest true "Book properties"
// @Success 201 {object} entity.Book "Successfully created the book"
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 409 {object} echo.HTTPError "Book already exists"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /api/v1/books [post]
func (h *Handler) CreateBook(c echo.Context) error {
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

// UpdateBook godoc
// @Summary Updates an existing book
// @Description Updates an existing book with the provided title, author, and year.
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body updateBookRequest true "Book properties"
// @Success 200 {object} entity.Book "Successfully updated the book"
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 404 {object} echo.HTTPError "Book not found"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /api/v1/books/{id} [patch]
func (h *Handler) UpdateBook(c echo.Context) error {
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

// @Summary Deletes a book
// @Description Deletes a book based on the provided ID.
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 "Book successfully deleted"
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 404 {object} echo.HTTPError "Book not found"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /api/v1/books/{id} [delete]
func (h *Handler) DeleteBook(c echo.Context) error {
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
