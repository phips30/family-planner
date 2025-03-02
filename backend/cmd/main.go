package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    user "family-planner/backend/internal/user"
)

// To be defined via ENV vars
const PORT = "8080"

func main() {
    r := mux.NewRouter()

    user.RegisterRoutes(r)

    fmt.Printf("Starting server on port %s\n", PORT)
    err := http.ListenAndServe(":" + PORT, r)
    if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}
