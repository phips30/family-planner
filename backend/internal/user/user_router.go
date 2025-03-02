package user

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/user", createUser).Methods("POST")
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var u User

    if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
        fmt.Println("error")
        return
    }
    fmt.Printf("Name: %s, DeviceId: %s\n", u.Name, u.DeviceId)
    fmt.Fprintf(w, "Name: %s, DeviceId: %s\n", u.Name, u.DeviceId)

}