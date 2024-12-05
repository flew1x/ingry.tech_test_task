package v1

import (
	"github.com/flew1x/ingry.tech_test_task/internal/service"

	"github.com/flew1x/ingry.tech_test_task/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	service *service.Service
	config  *config.Config
}

func NewHandler(service *service.Service, config *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  config,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	ech := echo.New()

	ech.Use(middleware.Logger())
	ech.Use(middleware.Recover())

	ech.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	ech.GET("/health", h.health)

	api := ech.Group("/api/v1")
	{
		books := api.Group("/books")
		{
			books.GET("", h.getBooks)
			books.GET("/:id", h.getBookByID)
			books.POST("", h.createBook)
			books.PUT("/:id", h.updateBook)
			books.DELETE("/:id", h.deleteBook)
		}
	}

	return ech
}
