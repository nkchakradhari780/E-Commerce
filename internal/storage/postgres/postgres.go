package postgres

import (
	"database/sql"
	"fmt"

	"github.com/nkchakradhari780/practice9/internal/config"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(cfg *config.Config) (*Postgres, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.SSLMode,
	)
	
	db, err := sql.Open("postgres", dsn); 
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres connetion pool: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return &Postgres{
		Db: db,
	}, nil
}