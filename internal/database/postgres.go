package database

import (
	"database/sql"
	"fmt"

	"go-todo/internal/config"

	_ "github.com/lib/pq"
)

func Open(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	return sql.Open("postgres", dsn)
}
