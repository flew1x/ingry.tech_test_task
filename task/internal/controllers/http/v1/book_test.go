package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"github.com/flew1x/ingry.tech_test_task/internal/service"
	"github.com/flew1x/ingry.tech_test_task/internal/service/mocks"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBooks(t *testing.T) {
	e := echo.New()
	mockService := mocks.NewIBookService(t)
	handler := &Handler{service: &service.Service{Book: mockService}}

	t.Run("GET /books - success", func(t *testing.T) {
		mockBooks := []entity.Book{
			{ID: uuid.New(), Title: "Book 1", Author: "Author 1", Year: 2021},
			{ID: uuid.New(), Title: "Book 2", Author: "Author 2", Year: 2020},
		}
		mockService.On("GetAll").Return(mockBooks, nil)

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.GetBooks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var resp getBooksResponse
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, len(mockBooks), len(resp.Books))
		}
	})

	t.Run("GET /books/:id - success", func(t *testing.T) {
		id := uuid.New()
		mockBook := entity.Book{ID: id, Title: "Book 1", Author: "Author 1", Year: 2021}
		mockService.On("GetByID", id).Return(mockBook, nil)

		req := httptest.NewRequest(http.MethodGet, "/books/"+id.String(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		if assert.NoError(t, handler.GetBookByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var resp getBookByIDResponse
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, mockBook.Title, resp.Book.Title)
		}
	})

	t.Run("GET /books/:id - not found", func(t *testing.T) {
		id := uuid.New()
		mockService.On("GetByID", id).Return(entity.Book{}, service.ErrBookNotFound)

		req := httptest.NewRequest(http.MethodGet, "/books/"+id.String(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		err := handler.GetBookByID(c)
		if assert.Error(t, err) {
			httpErr, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusNotFound, httpErr.Code)
			assert.Equal(t, "book not found", httpErr.Message)
		}
	})

	t.Run("POST /books - success", func(t *testing.T) {
		reqBody := createBookRequest{
			Title:  "New Book",
			Author: "Author",
			Year:   2024,
		}
		mockBook := entity.Book{
			ID:     uuid.New(),
			Title:  reqBody.Title,
			Author: reqBody.Author,
			Year:   reqBody.Year,
		}
		mockService.On("Create", reqBody.Title, reqBody.Author, reqBody.Year).Return(mockBook, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.CreateBook(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var resp entity.Book
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, mockBook.Title, resp.Title)
		}
	})

	t.Run("POST /books - invalid request", func(t *testing.T) {
		reqBody := createBookRequest{}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.CreateBook(c)
		if assert.Error(t, err) {
			httpErr, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusBadRequest, httpErr.Code)
		}
	})

	t.Run("PUT /books/:id - success", func(t *testing.T) {
		id := uuid.New()
		reqBody := updateBookRequest{
			Title:  "Updated Book",
			Author: "Updated Author",
			Year:   2024,
		}
		mockBook := entity.Book{
			ID:     id,
			Title:  reqBody.Title,
			Author: reqBody.Author,
			Year:   reqBody.Year,
		}
		mockService.On("Update", mockBook).Return(mockBook, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/books/"+id.String(), bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		if assert.NoError(t, handler.UpdateBook(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var resp entity.Book
			err := json.Unmarshal(rec.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, mockBook.Title, resp.Title)
		}
	})

	t.Run("PUT /books/:id - not found", func(t *testing.T) {
		id := uuid.New()
		reqBody := updateBookRequest{
			Title:  "Updated Book",
			Author: "Updated Author",
			Year:   2024,
		}
		mockService.On("Update", entity.Book{
			ID:     id,
			Title:  reqBody.Title,
			Author: reqBody.Author,
			Year:   reqBody.Year,
		}).Return(entity.Book{}, service.ErrBookNotFound)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/books/"+id.String(), bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		err := handler.UpdateBook(c)
		if assert.Error(t, err) {
			httpErr, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusNotFound, httpErr.Code)
			assert.Equal(t, "book not found", httpErr.Message)
		}
	})

	t.Run("PUT /books/:id - invalid request", func(t *testing.T) {
		id := uuid.New()
		reqBody := updateBookRequest{
			Title:  "",
			Author: "Author",
			Year:   2024,
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/books/"+id.String(), bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		err := handler.UpdateBook(c)
		if assert.Error(t, err) {
			httpErr, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusBadRequest, httpErr.Code)
			assert.Equal(t, "title should not be empty", httpErr.Message)
		}
	})

	t.Run("DELETE /books/:id - success", func(t *testing.T) {
		id := uuid.New()
		mockService.On("Delete", id).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/books/"+id.String(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		if assert.NoError(t, handler.DeleteBook(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})

	t.Run("DELETE /books/:id - not found", func(t *testing.T) {
		id := uuid.New()
		mockService.On("Delete", id).Return(service.ErrBookNotFound)

		req := httptest.NewRequest(http.MethodDelete, "/books/"+id.String(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(id.String())

		err := handler.DeleteBook(c)
		if assert.Error(t, err) {
			httpErr, ok := err.(*echo.HTTPError)
			assert.True(t, ok)
			assert.Equal(t, http.StatusNotFound, httpErr.Code)
			assert.Equal(t, "book not found", httpErr.Message)
		}
	})
}
