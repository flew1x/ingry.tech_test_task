package config

import (
	"net/url"
	"time"
)

const (
	phPath             = "postgres"
	pgPasswordFieldEnv = "DB_PASSWORD"
)

type IPostgresConfig interface {
	// GetPostgresHost returns the host for the postgres database.
	GetPostgresHost() string

	// GetPostgresUserInfo returns the user info for the postgres database.
	GetPostgresUserInfo() *url.Userinfo

	// GetPostgresDatabaseName returns the database name for the postgres database.
	GetPostgresDatabaseName() string

	// GetPostgresSSLMode returns the SSL mode for the postgres database.
	GetPostgresSSLMode() string

	// GetPostgresMaxCons returns the maximum number of connections for the postgres database.
	GetPostgresMaxCons() int

	// GetPostgresMaxIdleCons returns the maximum number of idle connections for the postgres database.
	GetPostgresMaxIdleCons() int

	// GetPostgresMaxConLifetime returns the maximum lifetime of a connection for the postgres database.
	GetPostgresMaxConLifetime() time.Duration
}

type PostgresConfig struct {
	Host            string        `koanf:"host"`
	User            string        `koanf:"user"`
	DatabaseName    string        `koanf:"dbname"`
	SSLMode         string        `koanf:"sslmode"`
	MaxCons         int           `koanf:"max_cons"`
	MaxIdleCons     int           `koanf:"max_idle_cons"`
	MaxConnLifetime time.Duration `koanf:"max_conn_lifetime"`
	Password        string        `koanf:"password"`
}

func NewPostgresConfig() *PostgresConfig {
	postgresConfig := &PostgresConfig{
		Password: mustStringFromEnv(pgPasswordFieldEnv),
	}

	mustUnmarshalStruct(phPath, &postgresConfig)

	return postgresConfig
}

func (c *PostgresConfig) GetPostgresHost() string {
	return c.Host
}

func (c *PostgresConfig) GetPostgresUserInfo() *url.Userinfo {
	return url.UserPassword(c.User, c.Password)
}

func (c *PostgresConfig) GetPostgresDatabaseName() string {
	return c.DatabaseName
}

func (c *PostgresConfig) GetPostgresSSLMode() string {
	return c.SSLMode
}

func (c *PostgresConfig) GetPostgresMaxCons() int {
	return c.MaxCons
}

func (c *PostgresConfig) GetPostgresMaxIdleCons() int {
	return c.MaxIdleCons
}

func (c *PostgresConfig) GetPostgresMaxConLifetime() time.Duration {
	return c.MaxConnLifetime
}

func (c *PostgresConfig) GetPostgresPassword() string {
	return c.Password
}
