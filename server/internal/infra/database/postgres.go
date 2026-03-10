package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/m-bromo/my-game-list/config"
)

func NewPostgresConnection(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.PostgresDB.Host,
		cfg.PostgresDB.Port,
		cfg.PostgresDB.Name,
		cfg.PostgresDB.User,
		cfg.PostgresDB.Password,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
