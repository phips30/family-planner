package api

import (
	"encoding/json"
	"family-planner/backend/internal/family/domain/service"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type GroupRouter struct {
	router  *mux.Router
	service service.GroupService
}

type NewGroupDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GroupMemberDto struct {
	GroupId uuid.UUID `json:"groupId"`
	Email   string    `json:"email"`
}

type GroupDto struct {
	GroupId uuid.UUID `json:"groupId"`
	Name    string    `json:"name"`
	Members []UserDto `json:"members"`
}

func NewGroupRouter(router *mux.Router, service service.GroupService) *GroupRouter {
	groupRouter := &GroupRouter{
		router:  router,
		service: service,
	}

	router.HandleFunc("/group", groupRouter.createGroup).Methods("POST")
	router.HandleFunc("/group/member", groupRouter.addGroupMember).Methods("POST")
	router.HandleFunc("/group/member/{groupId}", groupRouter.getGroupMembers).Methods("GET")

	return groupRouter
}

func (g *GroupRouter) createGroup(w http.ResponseWriter, r *http.Request) {
	var newGroupRequest *NewGroupDto

	if err := json.NewDecoder(r.Body).Decode(&newGroupRequest); err != nil {
		log.Println("error")
		return
	}

	log.Printf("Trying to create new group - Name: %s, Requested by: %s\n", newGroupRequest.Email, newGroupRequest.Email)

	newGroup, err := g.service.CreateGroup(newGroupRequest.Name, newGroupRequest.Email)
	if err != nil || newGroup == nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Group created - Id: %s Name: %s, \n", newGroup.Id, newGroup.Name)
	w.WriteHeader(http.StatusOK)
}

func (g *GroupRouter) addGroupMember(w http.ResponseWriter, r *http.Request) {
	var newGroupMemberRequest *GroupMemberDto

	if err := json.NewDecoder(r.Body).Decode(&newGroupMemberRequest); err != nil {
		log.Println("error")
		return
	}

	log.Printf("Trying to enter group - GroupId: %s, Requested by: %s\n", newGroupMemberRequest.GroupId, newGroupMemberRequest.Email)

	newGroupMember, err := g.service.AddGroupMember(newGroupMemberRequest.GroupId, newGroupMemberRequest.Email)
	if err != nil || newGroupMember == nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Entered group - Id: %s Name: %s, \n", newGroupMember.Id, newGroupMember.Name)
	w.WriteHeader(http.StatusOK)
}

func (g *GroupRouter) getGroupMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		http.Error(w, "No proper UUID provided", http.StatusBadRequest)
	}

	fmt.Printf("Trying to find group members for group id: %s\n", groupId)

	groupAgg, _ := g.service.FindGroupMembers(groupId)
	if groupAgg == nil {
		fmt.Printf("Group not found %s", groupAgg)
		http.Error(w, "User not found", http.StatusBadRequest)
	} else {
		var groupMembers []UserDto
		for _, groupMember := range groupAgg.GroupMembers {
			groupMembers = append(groupMembers, UserDto{
				Name:  groupMember.Name,
				Email: groupMember.Email,
			})
		}

		json.NewEncoder(w).Encode(GroupDto{
			GroupId: groupAgg.Group.Id,
			Name:    groupAgg.Group.Name,
			Members: groupMembers})
	}
}
