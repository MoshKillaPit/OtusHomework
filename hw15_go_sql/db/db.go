package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

type DB struct {
	Conn *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	// Открытие подключения
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Проверка соединения
	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Connected to the database successfully")
	return &DB{Conn: conn}, nil
}

func (db *DB) Close() error {
	if db.Conn != nil {
		return db.Conn.Close()
	}
	return nil
}
