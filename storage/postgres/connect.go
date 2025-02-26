package postgres

import (
	"database/sql"
	"edutest/pkg/config"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	connector := fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s sslmode = disable", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PASSWORD)
	db, err := sql.Open("postgres", connector)
	if err != nil {
		return nil, err
	}
	return db, nil
}
