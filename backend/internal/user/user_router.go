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

func (u *UserRouter) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user", u.createUser).Methods("POST")
}

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("error")
		return
	}

	fmt.Printf("Name: %s, DeviceId: %s\n", user.Name, user.DeviceId)
	var us, err = u.service.CreateUser(user.Name, user.DeviceId)

	fmt.Printf(err.Error())
	fmt.Printf("Name: %s, DeviceId: %s\n", us.Name, us.DeviceId)
	fmt.Fprintf(w, "Name: %s, DeviceId: %s\n", us.Name, us.DeviceId)

}
