package main

import (
    "fmt"
    "net/http"
    "os"
    "encoding/json"

    "github.com/gorilla/mux"
    user "family-planner/backend/internal/user"
)

func main() {
    r := mux.NewRouter()

    fmt.Println("Hello, World!")
    r.HandleFunc("/user", CreateUser).Methods("POST")

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    err := http.ListenAndServe(":8080", r)
    if err != nil {
        fmt.Printf("error starting server: %s\n", err)
        os.Exit(1)
    }
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var u user.User

    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        fmt.Println("error")
        return
    }
    fmt.Printf("Name: %s, DeviceId: %s\n", u.Name, u.DeviceId)
}
