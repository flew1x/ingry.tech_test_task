package v1

import (
	"github.com/flew1x/ingry.tech_test_task/internal/service"

	"github.com/flew1x/ingry.tech_test_task/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	ech.Use(ErrorHandler)

	ech.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	ech.GET("/health", h.health)
	ech.GET("/swagger/*", echoSwagger.WrapHandler)

	api := ech.Group("/api/v1")
	{
		books := api.Group("/books")
		{
			books.GET("", h.GetBooks)
			books.GET("/:id", h.GetBookByID)
			books.POST("", h.CreateBook)
			books.PUT("/:id", h.UpdateBook)
			books.DELETE("/:id", h.DeleteBook)
		}
	}

	return ech
}
