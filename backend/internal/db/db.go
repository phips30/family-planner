package db

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

    "github.com/jackc/pgx/v5/pgxpool"
)

const DB_CONNECTION_STRING = "postgres://postgres:postgres@localhost:5432/family_planner"

func Connect() *pgxpool.Pool {
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
        log.Fatal(err)
    }

    for _, sqlfile := range sqlfiles {
        data, err := ioutil.ReadFile(sqlfile)
        if err != nil {
            log.Fatal(err)
        }

        sql := string(data)
        _, err = pool.Exec(context.Background(), sql)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Executed SQL script: %s\n", sqlfile)

    }

    return nil
}