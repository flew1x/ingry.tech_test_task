package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/flew1x/ingry.tech_test_task/internal/config"
	v1 "github.com/flew1x/ingry.tech_test_task/internal/controllers/http/v1"
	"github.com/flew1x/ingry.tech_test_task/internal/database"
	"github.com/flew1x/ingry.tech_test_task/internal/repository"
	"github.com/flew1x/ingry.tech_test_task/internal/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	config *config.Config
	router *echo.Echo
}

func NewApp(config *config.Config) *App {
	db := database.InitDatabase(config.PostgresConfig)

	repo := repository.NewRepository(db.DB)
	serv := service.NewService(repo)

	handler := v1.NewHandler(serv, config)
	router := handler.InitRoutes()

	return &App{
		config: config,
		router: router,
	}
}

func (a *App) Run() error {
	if err := a.router.Start(fmt.Sprintf(":%d", a.config.RestConfig.GetPort())); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	if err := a.router.Close(); err != nil {
		return err
	}

	return nil
}
