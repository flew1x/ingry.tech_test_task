package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
