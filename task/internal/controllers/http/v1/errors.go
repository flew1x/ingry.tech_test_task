package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrInvalidID           = echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	ErrInvalidRequest      = echo.NewHTTPError(http.StatusBadRequest, "invalid request")

	ErrTitleShouldNotBeEmpty  = echo.NewHTTPError(http.StatusBadRequest, "title should not be empty")
	ErrAuthorShouldNotBeEmpty = echo.NewHTTPError(http.StatusBadRequest, "author should not be empty")
	ErrYearShouldNotBeEmpty   = echo.NewHTTPError(http.StatusBadRequest, "year should not be empty")
)
