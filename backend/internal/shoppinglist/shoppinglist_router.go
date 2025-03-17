package shoppinglist

import (
	"encoding/json"
	// Todo: this dependency needs to be removed

	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRouter struct {
	router        *mux.Router
	mongoDbClient *mongo.Database
	service       *ShoppinglistService
}

type ShoppinglistItemDto struct {
	Name          string    `json:"name"`
	AddedByUserId uuid.UUID `json:"addedByUserId"`
	AddedAt       time.Time `json:"addedAt"`
	Bought        bool      `json:"bought"`
}

func NewShoppinglistRouter(router *mux.Router, mongoDbClient *mongo.Database, shoppinglistService *ShoppinglistService) *ShoppinglistRouter {
	shoppinglistRouter := &ShoppinglistRouter{
		router:        router,
		mongoDbClient: mongoDbClient,
		service:       shoppinglistService,
	}

	router.HandleFunc("/shopping-list/{userId}", shoppinglistRouter.findList).Methods("GET")
	router.HandleFunc("/shopping-list", shoppinglistRouter.createList).Methods("POST")

	return shoppinglistRouter
}

func (s *ShoppinglistRouter) findList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["userId"])
	if err != nil {
		log.Println("No proper uuid provided: ", vars["userId"])
		http.Error(w, "No proper uuid provided", http.StatusBadRequest)
	}

	shoppingListItems, err := s.service.LoadShoppingListForUserid(userId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var shoppinglistResponse []ShoppinglistItemDto
	for _, shoppinlistItem := range shoppingListItems {
		shoppinglistResponse = append(shoppinglistResponse, *mapFromDomainObject(shoppinlistItem))
	}

	json.NewEncoder(w).Encode(shoppinglistResponse)

}

func (s *ShoppinglistRouter) createList(w http.ResponseWriter, r *http.Request) {
	var shoppinglistRequest []ShoppinglistItemDto
	if err := json.NewDecoder(r.Body).Decode(&shoppinglistRequest); err != nil {
		log.Println("error parsing request")
		http.Error(w, "error parsing request", http.StatusBadRequest)
	}

	var shoppinglist []ShoppinglistItem
	for _, shoppinglistItemDto := range shoppinglistRequest {
		shoppinglist = append(shoppinglist, *shoppinglistItemDto.mapToDomainObject())
	}

	_, err := s.service.SaveShoppingList(shoppinglist)
	if err != nil {
		log.Println("Error saving shopping list:", err.Error())
	}
	json.NewEncoder(w).Encode(shoppinglistRequest)
}

func (itemdto *ShoppinglistItemDto) mapToDomainObject() *ShoppinglistItem {
	return NewShoppinglistItem(
		itemdto.Name,
		itemdto.AddedAt,
		itemdto.AddedByUserId,
		itemdto.Bought)
}

func mapFromDomainObject(shoppinglistItem ShoppinglistItem) *ShoppinglistItemDto {
	return &ShoppinglistItemDto{
		Name:          shoppinglistItem.Name,
		AddedByUserId: shoppinglistItem.UserId,
		AddedAt:       shoppinglistItem.AddedAt,
		Bought:        shoppinglistItem.Bought,
	}
}
