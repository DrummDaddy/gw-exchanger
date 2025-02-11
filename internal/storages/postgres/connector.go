package postgres

import (
	"database/sql"
	"fmt"
	"gw-exchanger/internal/config"

	_ "github.com/lib/pq"
)

func NewConnector(cfg config.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	return sql.Open("postgres", connStr)
}
