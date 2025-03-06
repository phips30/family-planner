package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	router  *mux.Router
	service UserService
}

func NewUserRouter(router *mux.Router, service UserService) *UserRouter {
	userRouter := &UserRouter{
		router:  router,
		service: service,
	}

	router.HandleFunc("/user", userRouter.createUser).Methods("POST")

	return userRouter
}

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	var user *User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error")
		return
	}

	fmt.Printf("Name: %s, DeviceId: %s\n", user.Name, user.DeviceId)
	newUser, err := u.service.CreateUser(user.Name, user.DeviceId)

	if err != nil || newUser == nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		fmt.Printf("Name: %s, DeviceId: %s\n", newUser.Name, newUser.DeviceId)
		fmt.Fprintf(w, "Name: %s, DeviceId: %s\n", newUser.Name, newUser.DeviceId)
	}

}
