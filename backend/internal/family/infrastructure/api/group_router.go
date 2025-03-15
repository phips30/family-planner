package api

import (
	"encoding/json"
	"family-planner/backend/internal/family/domain/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type GroupRouter struct {
	router  *mux.Router
	service service.GroupService
}

type NewGroupDto struct {
	GroupName string `json:"groupName"`
	Email     string `json:"email"`
}

type GroupMemberDto struct {
	GroupId uuid.UUID `json:"groupId"`
	Email   string    `json:"email"`
}

type GroupMemberResponseDto struct {
	User     UserDto   `json:"user"`
	JoinedAt time.Time `json:"joinedAt"`
}

type GroupResponseDto struct {
	GroupId uuid.UUID                `json:"groupId"`
	Name    string                   `json:"name"`
	Members []GroupMemberResponseDto `json:"members"`
}

func NewGroupRouter(router *mux.Router, service service.GroupService) *GroupRouter {
	groupRouter := &GroupRouter{
		router:  router,
		service: service,
	}

	router.HandleFunc("/group", groupRouter.createGroup).Methods("POST")
	router.HandleFunc("/group/{email}", groupRouter.getGroupByUserEmail).Methods("GET")
	router.HandleFunc("/group/member", groupRouter.addGroupMember).Methods("POST")

	return groupRouter
}

func (g *GroupRouter) createGroup(w http.ResponseWriter, r *http.Request) {
	var newGroupRequest *NewGroupDto

	if err := json.NewDecoder(r.Body).Decode(&newGroupRequest); err != nil {
		log.Println("error")
		return
	}

	log.Printf("Trying to create new group - Name: %s, Requested by: %s\n", newGroupRequest.GroupName, newGroupRequest.Email)

	newGroup, err := g.service.CreateGroup(newGroupRequest.GroupName, newGroupRequest.Email)
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
	fmt.Printf("Entered group -Name: %s, \n", newGroupMember.Group.Name)
	w.WriteHeader(http.StatusOK)
}

func (g *GroupRouter) getGroupByUserEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	fmt.Printf("Trying to find group members for email: %s\n", email)

	groupAgg, _ := g.service.FindGroupMembersByEmail(email)
	if groupAgg == nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
	} else {
		var groupMembers []GroupMemberResponseDto
		for _, groupMember := range groupAgg.GroupMembers {
			groupMembers = append(groupMembers, GroupMemberResponseDto{
				User:     UserDto{Name: groupMember.User.Name, Email: groupMember.User.Email},
				JoinedAt: groupMember.CreatedAt,
			})
		}

		json.NewEncoder(w).Encode(GroupResponseDto{
			GroupId: groupAgg.Group.Id,
			Name:    groupAgg.Group.Name,
			Members: groupMembers})
	}
}
