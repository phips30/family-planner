package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    fmt.Println("Hello, World!")
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
