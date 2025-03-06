package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

const DB_CONNECTION_STRING = "postgres://postgres:postgres@localhost:5432/family_planner"

func Connect() *pgxpool.Pool {
	// Todo: make sure this is only called once and always returns the same dbpool after first init
	dbpool, err := pgxpool.New(context.Background(), DB_CONNECTION_STRING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool
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
			return fmt.Errorf("err: %v", err)
		} else {
			fmt.Printf("Executed SQL script: %s\n", sqlfile)
		}
	}

	return nil
}
