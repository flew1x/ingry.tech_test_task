package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"github.com/flew1x/ingry.tech_test_task/internal/config"
	"github.com/flew1x/ingry.tech_test_task/internal/entity"
	"gorm.io/gorm"
)

type GormClient struct {
	DB *gorm.DB
}

func InitDatabase(config config.IPostgresConfig) *GormClient {
	db := OpenPostgres(config)

	return &GormClient{DB: db}
}

func OpenPostgres(cfg config.IPostgresConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		AddressTemplate,
		cfg.GetPostgresUserInfo(),
		cfg.GetPostgresHost(),
		cfg.GetPostgresDatabaseName(),
		cfg.GetPostgresSSLMode(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entity.Book{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
