package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				return c.JSON(he.Code, echo.Map{"error": he.Message})
			}

			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return nil
	}
}
