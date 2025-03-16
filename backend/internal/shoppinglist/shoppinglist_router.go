package shoppinglist

import (
	"encoding/json"
	// Todo: this dependency needs to be removed

	"fmt"
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
	router.HandleFunc("/shopping-list/{item}", shoppinglistRouter.createList).Methods("POST")

	return shoppinglistRouter
}

func (s *ShoppinglistRouter) findList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	fmt.Printf("Trying to find shopping list - UserId: %s\n", userId)
}

func (s *ShoppinglistRouter) createList(w http.ResponseWriter, r *http.Request) {
	var shoppinglistRequest []ShoppinglistItemDto

	if err := json.NewDecoder(r.Body).Decode(&shoppinglistRequest); err != nil {
		log.Println("error parsing request")
		http.Error(w, "error parsing request", http.StatusBadRequest)
	}

	var shoppinglist []ShoppinglistItem
	for _, shoppinglistItemDto := range shoppinglistRequest {
		shoppinglistItem := NewShoppinglistItem(
			shoppinglistItemDto.Name,
			shoppinglistItemDto.AddedAt,
			shoppinglistItemDto.AddedByUserId,
			shoppinglistItemDto.Bought)

		shoppinglist = append(shoppinglist, *shoppinglistItem)
	}

	_, err := s.service.SaveShoppingList(shoppinglist)
	if err != nil {
		log.Println("Error saving shopping list:", err.Error())
	}
	json.NewEncoder(w).Encode(shoppinglistRequest)
}
