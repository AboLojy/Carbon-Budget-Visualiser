package db

import (
	"database/sql"

	configs "github.com/AboLojy/Carbon-Budget-Visualiser/config"
	_ "github.com/lib/pq"
)

func NewPostgresDb(cfg configs.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil
}
