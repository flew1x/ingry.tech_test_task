package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/flew1x/ingry.tech_test_task/docs"
	"github.com/flew1x/ingry.tech_test_task/internal/app"
	"github.com/flew1x/ingry.tech_test_task/internal/config"

	"golang.org/x/sync/errgroup"
)

var (
	ErrGracefulStop = errors.New("graceful stop signal received")
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Test Task
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email kruezz@mail.ru

// @BasePath /api/v1
func main() {
	cfg := config.NewConfig()
	cfg.InitConfig(os.Getenv("CONFIG_PATH"), os.Getenv("CONFIG_FILE"))

	application := app.NewApp(cfg)

	var errorGroup errgroup.Group

	errorGroup.Go(func() error {
		return application.Run()
	})

	errorGroup.Go(func() error {
		return gracefulStop(ErrGracefulStop)
	})

	switch err := errorGroup.Wait(); {
	case errors.Is(errorGroup.Wait(), ErrGracefulStop):
		application.Stop()

		log.Println("application stopped")
	default:
		log.Println("failed to start application:", err)
	}
}

func gracefulStop(gracefulStopError error) error {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	return gracefulStopError
}
