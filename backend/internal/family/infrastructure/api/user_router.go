package api

import (
	"encoding/json"
	"family-planner/backend/internal/family/domain/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	router  *mux.Router
	service service.UserService
}

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserRouter(router *mux.Router, service service.UserService) *UserRouter {
	userRouter := &UserRouter{
		router:  router,
		service: service,
	}

	router.HandleFunc("/user", userRouter.createUser).Methods("POST")
	router.HandleFunc("/user", userRouter.findUser).Methods("GET")

	return userRouter
}

func (u *UserRouter) findUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get("email")

	fmt.Printf("Trying to find user - Email: %s\n", email)

	user, _ := u.service.FindByEmail(email)
	if user == nil {
		fmt.Printf("User not found %s", user)
		http.Error(w, "User not found", http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(UserDto{Name: user.Name, Email: user.Email})
	}
}

func (u *UserRouter) createUser(w http.ResponseWriter, r *http.Request) {
	var userRequest *UserDto

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		fmt.Println("error")
		return
	}

	fmt.Printf("Trying to create user - Name: %s, Email: %s\n", userRequest.Name, userRequest.Email)
	newUser, err := u.service.CreateUser(userRequest.Name, userRequest.Email)

	if err != nil || newUser == nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("User created - Id: %s Name: %s, Email: %s\n", newUser.Id, newUser.Name, newUser.Email)
	fmt.Fprintf(w, "Name: %s, Email: %s\n", newUser.Name, newUser.Email)
}
