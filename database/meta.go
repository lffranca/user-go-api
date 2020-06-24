package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// GetDBMeta GetDBMeta
func GetDBMeta() (*sql.DB, error) {
	dbHost := os.Getenv("META_DB_HOST")
	if dbHost == "" {
		return nil, errors.New("META_DB_HOST")
	}
	dbUser := os.Getenv("META_DB_USER")
	if dbUser == "" {
		return nil, errors.New("META_DB_USER")
	}
	dbPass := os.Getenv("META_DB_PASSWORD")
	if dbPass == "" {
		return nil, errors.New("META_DB_PASSWORD")
	}
	dbDatabase := os.Getenv("META_DB_DATABASE")
	if dbDatabase == "" {
		return nil, errors.New("META_DB_DATABASE")
	}
	dbPort := os.Getenv("META_DB_PORT")
	if dbPort == "" {
		return nil, errors.New("META_DB_PORT")
	}
	dbSSLMode := os.Getenv("META_DB_SSLMODE")
	if dbSSLMode == "" {
		return nil, errors.New("META_DB_SSLMODE")
	}

	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbDatabase,
		dbSSLMode)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
