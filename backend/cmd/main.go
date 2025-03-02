package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    user "family-planner/backend/internal/user"
    "family-planner/backend/internal/db"
)

// To be defined via ENV vars
const PORT = "8080"

func main() {
    dbpool := db.Connect()
    defer dbpool.Close()

    fmt.Printf("Initializing database ...\n")
    err := db.InitDb(dbpool)
    if err != nil {
        fmt.Printf("Error initializing database: %s\n", err)
    }
    fmt.Printf("Initializing database completed.\n")


    r := mux.NewRouter()

    user.RegisterRoutes(r)

    fmt.Printf("Starting server on port %s\n", PORT)
    err = http.ListenAndServe(":" + PORT, r)
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
        os.Exit(1)
    }
}
