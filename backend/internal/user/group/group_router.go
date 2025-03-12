package group

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GroupRouter struct {
	router  *mux.Router
	service *GroupService
}

type GroupDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGroupRouter(router *mux.Router, service GroupService) *GroupRouter {
	groupRouter := &GroupRouter{
		router:  router,
		service: service,
	}

	router.HandleFunc("/group", groupRouter.createGroup).Methods("POST")

	return groupRouter
}

func (g *GroupRouter) createGroup(w http.ResponseWriter, r *http.Request) {
	var newGroupRequest *GroupDto

	if err := json.NewDecoder(r.Body).Decode(&newGroupRequest); err != nil {
		log.Println("error")
		return
	}

	log.Printf("Trying to create new group - Name: %s, Requested by: %s\n", newGroupRequest.Email, newGroupRequest.Email)

	newGroup, err := g.service.CreateNewGroup(newGroupRequest.Name, newGroupRequest.Email)
	if err != nil || newGroup == nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Group created - Id: %s Name: %s, \n", newGroup.Id, newGroup.Name)
	w.WriteHeader(http.StatusOK)
}
