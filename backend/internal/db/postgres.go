package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

const DB_CONNECTION_STRING = "postgres://postgres:postgres@localhost:5432/family_planner"

func ConnectPostgres() (*pgxpool.Pool, error) {
	// Todo: make sure this is only called once and always returns the same dbpool after first init
	dbpool, err := pgxpool.New(context.Background(), DB_CONNECTION_STRING)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}
	return dbpool, nil
}

func InitDb(pool *pgxpool.Pool) error {
	sqlfiles, err := filepath.Glob("../db/sql/*.sql")
	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	for _, sqlfile := range sqlfiles {
		data, err := os.ReadFile(sqlfile)
		if err != nil {
			return fmt.Errorf("err: %v", err)
		}

		sql := string(data)
		_, err = pool.Exec(context.Background(), sql)
		if err != nil {
			log.Printf("err: %v", err)
		} else {
			log.Printf("Executed SQL script: %s\n", sqlfile)
		}
	}

	return nil
}
