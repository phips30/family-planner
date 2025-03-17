package api

import (
	"encoding/json"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"family-planner/backend/internal/shoppinglist/domain/service"
	"family-planner/backend/internal/shoppinglist/infrastructure/api/dto"

	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRouter struct {
	router        *mux.Router
	mongoDbClient *mongo.Database
	service       *service.ShoppinglistService
}

func NewShoppinglistRouter(router *mux.Router, mongoDbClient *mongo.Database, shoppinglistService *service.ShoppinglistService) *ShoppinglistRouter {
	shoppinglistRouter := &ShoppinglistRouter{
		router:        router,
		mongoDbClient: mongoDbClient,
		service:       shoppinglistService,
	}

	router.HandleFunc("/shopping-list/user/{userId}", shoppinglistRouter.findListForUser).Methods("GET")
	router.HandleFunc("/shopping-list/group/{groupId}", shoppinglistRouter.findListForGroup).Methods("GET")
	router.HandleFunc("/shopping-list", shoppinglistRouter.createList).Methods("POST")

	return shoppinglistRouter
}

func (s *ShoppinglistRouter) findListForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["userId"])
	if err != nil {
		log.Println("No proper uuid provided: ", vars["userId"])
		http.Error(w, "No proper uuid provided", http.StatusBadRequest)
	}

	shoppinglistItems, err := s.service.LoadShoppingListForUserId(userId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var shoppinglistResponse []dto.ShoppinglistUserItemDto
	for _, shoppinlistItem := range shoppinglistItems {
		var shoppinglistUserItemDto dto.ShoppinglistUserItemDto
		shoppinglistResponse = append(shoppinglistResponse, *shoppinglistUserItemDto.MapFromDomainObject(shoppinlistItem))
	}

	json.NewEncoder(w).Encode(shoppinglistResponse)
}

func (s *ShoppinglistRouter) findListForGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		log.Println("No proper uuid provided: ", vars["groupId"])
		http.Error(w, "No proper uuid provided", http.StatusBadRequest)
	}

	shoppinglistItems, err := s.service.LoadShoppingListForGroupId(groupId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var shoppinglistResponse []dto.ShoppinglistGroupItemDto
	for _, shoppinlistItem := range shoppinglistItems {
		var shoppinglistGroupItemDto dto.ShoppinglistGroupItemDto
		shoppinglistResponse = append(shoppinglistResponse, *shoppinglistGroupItemDto.MapFromDomainObject(shoppinlistItem))
	}

	json.NewEncoder(w).Encode(shoppinglistResponse)
}

func (s *ShoppinglistRouter) createList(w http.ResponseWriter, r *http.Request) {
	var shoppinglistRequest []dto.ShoppinglistGroupItemDto
	if err := json.NewDecoder(r.Body).Decode(&shoppinglistRequest); err != nil {
		log.Println("error parsing request")
		http.Error(w, "error parsing request", http.StatusBadRequest)
	}

	var shoppinglist []entity.ShoppinglistItem
	for _, shoppinglistItemDto := range shoppinglistRequest {
		shoppinglist = append(shoppinglist, *shoppinglistItemDto.MapToDomainObject())
	}

	_, err := s.service.SaveShoppingList(shoppinglist)
	if err != nil {
		log.Println("Error saving shopping list:", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}
